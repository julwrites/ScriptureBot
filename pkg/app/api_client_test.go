package app

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSubmitQuery(t *testing.T) {
	handler := newMockApiHandler()
	ts := httptest.NewServer(handler)
	defer ts.Close()

	t.Run("Success", func(t *testing.T) {
		defer setEnv("BIBLE_API_URL", ts.URL)()
		ResetAPIConfigCache()

		req := QueryRequest{Query: QueryObject{Prompt: "hello"}}
		var resp OQueryResponse
		err := SubmitQuery(req, &resp, "")
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if resp.Text != "Answer text" {
			t.Errorf("Expected 'Answer text', got '%s'", resp.Text)
		}
	})

	t.Run("API Error", func(t *testing.T) {
		handler.statusCode = http.StatusInternalServerError
		handler.rawResponse = `{"error": {"code": 500, "message": "simulated error"}}`
		defer func() { // Reset handler
			handler.statusCode = http.StatusOK
			handler.rawResponse = ""
		}()

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
		handler.rawResponse = `{invalid json`
		defer func() { handler.rawResponse = "" }()

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
