package app

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSubmitQuery(t *testing.T) {
	// Mock server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check headers
		if r.Header.Get("Content-Type") != "application/json" {
			t.Errorf("Expected Content-Type application/json, got %s", r.Header.Get("Content-Type"))
		}

		// Decode request to verify it
		var req QueryRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}

		// Simple response based on input
		if req.Query.Prompt == "error" {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"error": {"code": 500, "message": "simulated error"}}`))
			return
		}

		if req.Query.Prompt == "badjson" {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{invalid json`))
			return
		}

		// Success response
		resp := VerseResponse{Verse: "Success Verse"}
		json.NewEncoder(w).Encode(resp)
	}))
	defer ts.Close()

	// Set env vars
	defer setEnv("BIBLE_API_URL", ts.URL)()

	// Test Case 1: Success
	t.Run("Success", func(t *testing.T) {
		req := QueryRequest{Query: QueryObject{Prompt: "hello"}}
		var resp VerseResponse
		err := SubmitQuery(req, &resp)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if resp.Verse != "Success Verse" {
			t.Errorf("Expected 'Success Verse', got '%s'", resp.Verse)
		}
	})

	// Test Case 2: API Error
	t.Run("API Error", func(t *testing.T) {
		req := QueryRequest{Query: QueryObject{Prompt: "error"}}
		var resp VerseResponse
		err := SubmitQuery(req, &resp)
		if err == nil {
			t.Error("Expected error, got nil")
		}
		// Expect error message to contain "simulated error"
		if err != nil && err.Error() != "api error (500): simulated error" {
			t.Errorf("Expected specific API error, got: %v", err)
		}
	})

	// Test Case 3: Bad JSON Response
	t.Run("Bad JSON", func(t *testing.T) {
		req := QueryRequest{Query: QueryObject{Prompt: "badjson"}}
		var resp VerseResponse
		err := SubmitQuery(req, &resp)
		if err == nil {
			t.Error("Expected error for bad JSON, got nil")
		}
	})

	// Test Case 4: No URL set
	t.Run("No URL", func(t *testing.T) {
		// Temporarily unset/clear the env var
		restore := setEnv("BIBLE_API_URL", "")
		defer restore()
		// Also unset PROJECT_ID to avoid Secret Manager lookup
		defer setEnv("GCLOUD_PROJECT_ID", "")()

		req := QueryRequest{}
		var resp VerseResponse
		err := SubmitQuery(req, &resp)
		if err == nil {
			t.Error("Expected error when BIBLE_API_URL is unset")
		}
	})
}

func TestGetAPIConfig_SecretManagerFallback(t *testing.T) {
	// Ensure Env Vars are empty
	defer setEnv("BIBLE_API_URL", "")()
	defer setEnv("BIBLE_API_KEY", "")()
	defer setEnv("GCLOUD_PROJECT_ID", "test-project")()

	// Mock the secret function
	oldGetSecret := getSecretFunc
	defer func() { getSecretFunc = oldGetSecret }()

	getSecretFunc = func(project, name string) (string, error) {
		if project != "test-project" {
			return "", fmt.Errorf("unexpected project: %s", project)
		}
		if name == "BIBLE_API_URL" {
			return "http://secret-url.com", nil
		}
		if name == "BIBLE_API_KEY" {
			return "secret-key", nil
		}
		return "", fmt.Errorf("unexpected secret: %s", name)
	}

	url, key := getAPIConfig()

	if url != "http://secret-url.com" {
		t.Errorf("Expected URL 'http://secret-url.com', got '%s'", url)
	}
	if key != "secret-key" {
		t.Errorf("Expected Key 'secret-key', got '%s'", key)
	}
}
