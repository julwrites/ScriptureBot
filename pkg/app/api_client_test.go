package app

import (
	"testing"
)

func TestSubmitQuery(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		// Force cleanup of environment to ensure we test Secret Manager fallback
		// This handles cases where the runner might have lingering env vars
		defer UnsetEnv("BIBLE_API_URL")()
		defer UnsetEnv("BIBLE_API_KEY")()

		ResetAPIConfigCache()

		// Use a simple Verse query to verify connectivity.
		// Avoid using Prompt ("hello") as it triggers the LLM which might be unstable (500 errors).
		req := QueryRequest{
			Query:   QueryObject{Verses: []string{"John 3:16"}},
			Context: QueryContext{User: UserContext{Version: "NIV"}},
		}
		var resp VerseResponse
		err := SubmitQuery(req, &resp, "")
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
		err := SubmitQuery(req, &resp, "")
		if err == nil {
			t.Error("Expected error when BIBLE_API_URL is unset")
		}
	})
}
