package app

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/julwrites/ScriptureBot/pkg/utils"
)

// getSecretFunc is a variable to allow mocking in tests
var getSecretFunc = utils.GetSecret

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

func getAPIConfig(projectID string) (string, string) {
	configMutex.Lock()
	defer configMutex.Unlock()

	if configInitialized {
		return cachedAPIURL, cachedAPIKey
	}

	url := os.Getenv("BIBLE_API_URL")
	key := os.Getenv("BIBLE_API_KEY")

	// If env vars are missing, try to fetch from Secret Manager
	if url == "" || key == "" {
		envProjectID := os.Getenv("GCLOUD_PROJECT_ID")
		if envProjectID != "" {
			projectID = envProjectID
		}

		if projectID != "" {
			if url == "" {
				var err error
				url, err = getSecretFunc(projectID, "BIBLE_API_URL")
				if err != nil {
					log.Printf("Failed to fetch BIBLE_API_URL from Secret Manager: %v", err)
				}
			}
			if key == "" {
				var err error
				key, err = getSecretFunc(projectID, "BIBLE_API_KEY")
				if err != nil {
					log.Printf("Failed to fetch BIBLE_API_KEY from Secret Manager: %v", err)
				}
			}
		} else {
			log.Println("GCLOUD_PROJECT_ID is not set and no project ID passed, skipping Secret Manager lookup")
		}
	}

	cachedAPIURL = url
	cachedAPIKey = key
	configInitialized = true

	return url, key
}

// SubmitQuery sends the QueryRequest to the Bible API and unmarshals the response into result.
// result should be a pointer to the expected response struct.
func SubmitQuery(req QueryRequest, result interface{}, projectID string) error {
	apiURL, apiKey := getAPIConfig(projectID)
	if apiURL == "" {
		return fmt.Errorf("BIBLE_API_URL environment variable is not set")
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
