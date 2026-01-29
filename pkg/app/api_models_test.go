package app

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestQueryRequest_Marshal(t *testing.T) {
	req := QueryRequest{
		Query: QueryObject{
			Verses: []string{"John 3:16"},
			Words:  []string{"Love"},
			Prompt: "Tell me about love",
			Context: QueryContext{
				History: []string{"Previous query"},
				Schema:  "custom",
				Verses:  []string{"Gen 1:1"},
				Words:   []string{"Create"},
			},
		},
		User: UserOptions{
			Version: "NIV",
		},
	}

	data, err := json.Marshal(req)
	if err != nil {
		t.Fatalf("Failed to marshal QueryRequest: %v", err)
	}

	expectedFields := []string{
		`"verses":["John 3:16"]`,
		`"words":["Love"]`,
		`"prompt":"Tell me about love"`,
		`"history":["Previous query"]`,
		`"schema":"custom"`,
		`"version":"NIV"`,
	}

	jsonStr := string(data)
	for _, field := range expectedFields {
		if !strings.Contains(jsonStr, field) {
			t.Errorf("JSON output missing expected field: %s. Got: %s", field, jsonStr)
		}
	}
}

func TestVerseResponse_Unmarshal(t *testing.T) {
	jsonStr := `{"verse": "For God so loved the world..."}`
	var resp VerseResponse
	err := json.Unmarshal([]byte(jsonStr), &resp)
	if err != nil {
		t.Fatalf("Failed to unmarshal VerseResponse: %v", err)
	}

	if resp.Verse != "For God so loved the world..." {
		t.Errorf("Expected verse text, got: %s", resp.Verse)
	}
}

func TestWordSearchResponse_Unmarshal(t *testing.T) {
	jsonStr := `[
		{"verse": "John 3:16", "url": "http://bible.com/john3:16"},
		{"verse": "1 John 4:8", "url": "http://bible.com/1john4:8"}
	]`
	var resp WordSearchResponse
	err := json.Unmarshal([]byte(jsonStr), &resp)
	if err != nil {
		t.Fatalf("Failed to unmarshal WordSearchResponse: %v", err)
	}

	if len(resp) != 2 {
		t.Errorf("Expected 2 results, got %d", len(resp))
	}
	if resp[0].Verse != "John 3:16" {
		t.Errorf("Expected first verse to be John 3:16, got %s", resp[0].Verse)
	}
}

func TestPromptResponse_Unmarshal(t *testing.T) {
	jsonStr := `{
		"data": {
			"text": "God is love.",
			"references": [{"verse": "1 John 4:8", "url": "http://bible.com"}]
		},
		"meta": {
			"ai_provider": "openai"
		}
	}`
	var resp PromptResponse
	err := json.Unmarshal([]byte(jsonStr), &resp)
	if err != nil {
		t.Fatalf("Failed to unmarshal PromptResponse: %v", err)
	}

	if resp.Data.Text != "God is love." {
		t.Errorf("Expected text 'God is love.', got '%s'", resp.Data.Text)
	}
	if len(resp.Data.References) != 1 {
		t.Errorf("Expected 1 reference, got %d", len(resp.Data.References))
	}
	if resp.Meta.AIProvider != "openai" {
		t.Errorf("Expected AI provider 'openai', got '%s'", resp.Meta.AIProvider)
	}
}

func TestOQueryResponse_Unmarshal(t *testing.T) {
	// This tests direct unmarshal of OQueryResponse, which is still used internally or if schema matches
	jsonStr := `{
		"text": "God is love.",
		"references": [{"verse": "1 John 4:8", "url": "http://bible.com"}]
	}`
	var resp OQueryResponse
	err := json.Unmarshal([]byte(jsonStr), &resp)
	if err != nil {
		t.Fatalf("Failed to unmarshal OQueryResponse: %v", err)
	}

	if resp.Text != "God is love." {
		t.Errorf("Expected text 'God is love.', got '%s'", resp.Text)
	}
	if len(resp.References) != 1 {
		t.Errorf("Expected 1 reference, got %d", len(resp.References))
	}
}

func TestErrorResponse_Unmarshal(t *testing.T) {
	jsonStr := `{"error": {"code": 500, "message": "Internal Server Error"}}`
	var resp ErrorResponse
	err := json.Unmarshal([]byte(jsonStr), &resp)
	if err != nil {
		t.Fatalf("Failed to unmarshal ErrorResponse: %v", err)
	}

	if resp.Error.Code != 500 {
		t.Errorf("Expected error code 500, got %d", resp.Error.Code)
	}
	if resp.Error.Message != "Internal Server Error" {
		t.Errorf("Expected error message, got '%s'", resp.Error.Message)
	}
}
