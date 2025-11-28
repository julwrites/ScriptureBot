package app

import (
	"encoding/json"
	"net/http"
	"os"
)

// setEnv is a helper function to temporarily set an environment variable and return a function to restore it.
// It also temporarily unsets GCLOUD_PROJECT_ID to ensure secrets are loaded from environment variables
// (like the mock server URL) rather than Google Secret Manager.
func setEnv(key, value string) func() {
	originalValue, isSet := os.LookupEnv(key)
	os.Setenv(key, value)

	// Unset GCLOUD_PROJECT_ID to prevent Secret Manager usage during tests,
	// unless we are explicitly setting GCLOUD_PROJECT_ID itself.
	var projectID string
	var projectIDSet bool
	if key != "GCLOUD_PROJECT_ID" {
		projectID, projectIDSet = os.LookupEnv("GCLOUD_PROJECT_ID")
		if projectIDSet {
			os.Unsetenv("GCLOUD_PROJECT_ID")
		}
	}

	return func() {
		if isSet {
			os.Setenv(key, originalValue)
		} else {
			os.Unsetenv(key)
		}

		// Restore GCLOUD_PROJECT_ID if we unset it as a side effect
		if key != "GCLOUD_PROJECT_ID" && projectIDSet {
			os.Setenv("GCLOUD_PROJECT_ID", projectID)
		}
	}
}

// mockApiHandler is a flexible handler for the mock server.
type mockApiHandler struct {
	verseResponse      VerseResponse
	wordSearchResponse WordSearchResponse
	oQueryResponse     OQueryResponse
	statusCode         int
	rawResponse        string
}

// newMockApiHandler creates a new mockApiHandler with default success responses.
func newMockApiHandler() *mockApiHandler {
	return &mockApiHandler{
		statusCode: http.StatusOK,
		verseResponse: VerseResponse{
			Verse: "<p>In the beginning God created the heavens and the earth.</p>",
		},
		wordSearchResponse: WordSearchResponse{
			{Verse: "Found 1:1", URL: "http://found1"},
		},
		oQueryResponse: OQueryResponse{
			Text: "Answer text",
			References: []SearchResult{
				{Verse: "Ref 1:1", URL: "http://ref1"},
			},
		},
	}
}

// ServeHTTP handles the incoming requests and sends the configured response.
func (h *mockApiHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(h.statusCode)

	if h.rawResponse != "" {
		w.Write([]byte(h.rawResponse))
		return
	}

	var req QueryRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	if len(req.Query.Words) > 0 {
		json.NewEncoder(w).Encode(h.wordSearchResponse)
		return
	}

	if req.Query.Prompt != "" {
		json.NewEncoder(w).Encode(h.oQueryResponse)
		return
	}

	json.NewEncoder(w).Encode(h.verseResponse)
}
