package app

import (
	"testing"
)

func TestSubmitQuery(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		ResetAPIConfigCache()

		req := QueryRequest{Query: QueryObject{Prompt: "hello"}}
		var resp OQueryResponse
		err := SubmitQuery(req, &resp, "")
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		// In integration test mode, we expect some content but can't assert exact text
		// because the live API response might vary.
		if len(resp.Text) == 0 && len(resp.References) == 0 {
			t.Errorf("Expected some content (text or references), got empty response")
		}
	})

	t.Run("No URL", func(t *testing.T) {
		defer setEnv("BIBLE_API_URL", "")()
		ResetAPIConfigCache()

		req := QueryRequest{}
		var resp VerseResponse
		err := SubmitQuery(req, &resp, "")
		if err == nil {
			t.Error("Expected error when BIBLE_API_URL is unset")
		}
	})
}
