package app

import (
	"os"
	"testing"
)

func TestSubmitQuery(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		// Check if we should run integration test against real API
		// If BIBLE_API_URL is set and not example.com, we assume integration test mode
		realURL, hasURL := os.LookupEnv("BIBLE_API_URL")
		if hasURL && realURL != "" && realURL != "https://example.com" {
			t.Logf("Running integration test against real API: %s", realURL)
			// Ensure we have a key
			if _, hasKey := os.LookupEnv("BIBLE_API_KEY"); !hasKey {
				t.Log("Warning: BIBLE_API_URL set but BIBLE_API_KEY missing. Test might fail.")
			}
		} else {
			// Mock mode
			defer SetEnv("BIBLE_API_URL", "https://example.com")()
			defer SetEnv("BIBLE_API_KEY", "api_key")()
		}

		ResetAPIConfigCache()

		// Use a simple Verse query to verify connectivity.
		// Avoid using Prompt ("hello") as it triggers the LLM which might be unstable (500 errors).
		req := QueryRequest{
			Query:   QueryObject{Verses: []string{"John 3:16"}},
			User:    UserOptions{Version: "NIV"},
		}
		var resp VerseResponse
		err := SubmitQuery(req, &resp)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		// In integration test mode, we expect some content
		if len(resp.Verse) == 0 {
			t.Errorf("Expected verse content, got empty response")
		}
	})

	t.Run("No URL", func(t *testing.T) {
		defer SetEnv("BIBLE_API_URL", "")()
		ResetAPIConfigCache()

		req := QueryRequest{}
		var resp VerseResponse
		err := SubmitQuery(req, &resp)
		if err == nil {
			t.Error("Expected error when BIBLE_API_URL is unset")
		}
	})
}

func TestGetVersions(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		defer SetEnv("BIBLE_API_URL", "https://example.com")()
		defer SetEnv("BIBLE_API_KEY", "api_key")()
		ResetAPIConfigCache()

		resp, err := GetVersions(1, 10, "", "", "")
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if resp.Total != 2 {
			t.Errorf("Expected total 2, got %d", resp.Total)
		}
		if len(resp.Data) != 2 {
			t.Errorf("Expected 2 versions, got %d", len(resp.Data))
		}
	})

	t.Run("No URL", func(t *testing.T) {
		defer SetEnv("BIBLE_API_URL", "")()
		ResetAPIConfigCache()

		_, err := GetVersions(1, 10, "", "", "")
		if err == nil {
			t.Error("Expected error when BIBLE_API_URL is unset")
		}
	})
}
