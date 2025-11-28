package app

import (
	"net/http"
	"net/http/httptest"
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

	t.Run("API Error", func(t *testing.T) {
		handler := newMockApiHandler()
		ts := httptest.NewServer(handler)
		defer ts.Close()

		handler.statusCode = http.StatusInternalServerError
		handler.rawResponse = `{"error": {"code": 500, "message": "simulated error"}}`

		defer setEnv("BIBLE_API_URL", ts.URL)()
		ResetAPIConfigCache()

		req := QueryRequest{Query: QueryObject{Prompt: "error"}}
		var resp VerseResponse
		err := SubmitQuery(req, &resp, "")
		if err == nil {
			t.Error("Expected error, got nil")
		}
		if err.Error() != "api error (500): simulated error" {
			t.Errorf("Expected specific API error, got: %v", err)
		}
	})

	t.Run("Bad JSON", func(t *testing.T) {
		handler := newMockApiHandler()
		ts := httptest.NewServer(handler)
		defer ts.Close()

		handler.rawResponse = `{invalid json`

		defer setEnv("BIBLE_API_URL", ts.URL)()
		ResetAPIConfigCache()

		req := QueryRequest{Query: QueryObject{Prompt: "badjson"}}
		var resp VerseResponse
		err := SubmitQuery(req, &resp, "")
		if err == nil {
			t.Error("Expected error for bad JSON, got nil")
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
