package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSetWebhook(t *testing.T) {
	// Mock Telegram API
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		expectedPath := "/botFAKE_TOKEN/setWebhook"
		if r.URL.Path != expectedPath {
			t.Errorf("Expected path %s, got %s", expectedPath, r.URL.Path)
		}

		query := r.URL.Query()
		expectedURL := "https://example.com/FAKE_TOKEN"
		if query.Get("url") != expectedURL {
			t.Errorf("Expected url parameter %s, got %s", expectedURL, query.Get("url"))
		}

		resp := TelegramResponse{
			Ok:     true,
			Result: true,
		}
		json.NewEncoder(w).Encode(resp)
	}))
	defer ts.Close()

	// Override the base URL for testing
	// The server URL will be http://127.0.0.1:xxxxx
	// We need to strip the trailing slash if present, but httptest URL usually doesn't have it.
	// However, our code appends /bot%s...

	// Wait, telegramAPIBase is "https://api.telegram.org/bot".
	// So it expects the base to end with "bot" if we want to match the structure.
	// My mock server is just the root.

	// Let's adjust the override.
	originalBase := telegramAPIBase
	defer func() { telegramAPIBase = originalBase }()

	telegramAPIBase = ts.URL + "/bot" // Injecting /bot so the code works as expected

	err := setWebhook("FAKE_TOKEN", "https://example.com/FAKE_TOKEN")
	if err != nil {
		t.Fatalf("setWebhook failed: %v", err)
	}
}

func TestSetWebhook_Error(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := TelegramResponse{
			Ok:          false,
			Description: "Unauthorized",
		}
		json.NewEncoder(w).Encode(resp)
	}))
	defer ts.Close()

	originalBase := telegramAPIBase
	defer func() { telegramAPIBase = originalBase }()
	telegramAPIBase = ts.URL + "/bot"

	err := setWebhook("FAKE_TOKEN", "https://example.com/FAKE_TOKEN")
	if err == nil {
		t.Fatal("Expected error, got nil")
	}
	if err.Error() != "telegram API error: Unauthorized" {
		t.Errorf("Expected 'telegram API error: Unauthorized', got '%v'", err)
	}
}
