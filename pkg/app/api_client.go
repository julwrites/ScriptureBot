package app

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func getAPIConfig() (string, string) {
	url := os.Getenv("BIBLE_API_URL")
	key := os.Getenv("BIBLE_API_KEY")
	// Fallback/Default for development if needed, but per plan we rely on env vars.
	// Users should ensure these are set.
	return url, key
}

// SubmitQuery sends the QueryRequest to the Bible API and unmarshals the response into result.
// result should be a pointer to the expected response struct.
func SubmitQuery(req QueryRequest, result interface{}) error {
	apiURL, apiKey := getAPIConfig()
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
