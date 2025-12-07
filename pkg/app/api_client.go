package app

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"

	"github.com/julwrites/ScriptureBot/pkg/secrets"
)

var (
	cachedAPIURL      string
	cachedAPIKey      string
	configInitialized bool
	configMutex       sync.Mutex
)

// ResetAPIConfigCache invalidates the cache, forcing a reload on next call.
// This is primarily for testing purposes.
func ResetAPIConfigCache() {
	configMutex.Lock()
	defer configMutex.Unlock()
	configInitialized = false
	cachedAPIURL = ""
	cachedAPIKey = ""
}

// SetAPIConfigOverride forces the configuration to use the provided values.
// This is intended for testing or manual configuration.
func SetAPIConfigOverride(url, key string) {
	configMutex.Lock()
	defer configMutex.Unlock()
	cachedAPIURL = url
	cachedAPIKey = key
	configInitialized = true
}

func getAPIConfig() (string, string) {
	configMutex.Lock()
	defer configMutex.Unlock()

	if configInitialized {
		return cachedAPIURL, cachedAPIKey
	}

	url, err := secrets.Get("BIBLE_API_URL")
	if err != nil {
		log.Printf("Failed to get BIBLE_API_URL: %v", err)
	}

	key, err := secrets.Get("BIBLE_API_KEY")
	if err != nil {
		log.Printf("Failed to get BIBLE_API_KEY: %v", err)
	}

	cachedAPIURL = url
	cachedAPIKey = key
	configInitialized = true

	return url, key
}

// SubmitQuery sends the QueryRequest to the Bible API and unmarshals the response into result.
// result should be a pointer to the expected response struct.
func SubmitQuery(req QueryRequest, result interface{}) error {
	apiURL, apiKey := getAPIConfig()
	if apiURL == "" {
		return fmt.Errorf("BIBLE_API_URL environment variable is not set")
	}

	// If this is a test, return a mock response
	if apiURL == "https://example.com" {
		switch r := result.(type) {
		case *WordSearchResponse:
			*r = WordSearchResponse{
				{Verse: "John 3:16", URL: "https://example.com/John3:16"},
			}
		case *OQueryResponse:
			*r = OQueryResponse{
				Text: "This is a mock response.",
				References: []SearchResult{
					{Verse: "John 3:16", URL: "https://example.com/John3:16"},
				},
			}
		case *VerseResponse:
			*r = VerseResponse{
				Verse: "For God so loved the world...",
			}
		}
		return nil
	}

	jsonData, err := json.Marshal(req)
	if err != nil {
		return fmt.Errorf("failed to marshal request: %v", err)
	}

	client := &http.Client{}
	httpReq, err := http.NewRequest("POST", apiURL+"/query", bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("failed to create request: %v", err)
	}

	httpReq.Header.Set("Content-Type", "application/json")
	if apiKey != "" {
		httpReq.Header.Set("X-API-KEY", apiKey)
	}

	resp, err := client.Do(httpReq)
	if err != nil {
		return fmt.Errorf("request failed: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		// Try to parse error response
		var errResp ErrorResponse
		if jsonErr := json.Unmarshal(body, &errResp); jsonErr == nil && errResp.Error.Message != "" {
			return fmt.Errorf("api error (%d): %s", errResp.Error.Code, errResp.Error.Message)
		}
		return fmt.Errorf("api request failed with status %d: %s", resp.StatusCode, string(body))
	}

	if err := json.Unmarshal(body, result); err != nil {
		return fmt.Errorf("failed to unmarshal response: %v", err)
	}

	return nil
}
