package app

import (
	"strings"
	"testing"

	"github.com/julwrites/BotPlatform/pkg/def"
	"github.com/julwrites/ScriptureBot/pkg/utils"
)

func TestSecurity_AIQueryRestriction(t *testing.T) {
	// 1. Setup Environment
	// Set a mock admin ID
	defer SetEnv("TELEGRAM_ADMIN_ID", "99999")()

	// Mock SubmitQuery to capture the request locally
	originalSubmitQuery := SubmitQuery
	defer func() { SubmitQuery = originalSubmitQuery }()

	// We'll capture the request payload here
	var capturedReq QueryRequest
	SubmitQuery = func(req QueryRequest, result interface{}) error {
		capturedReq = req
		// Mock success response based on type
		switch r := result.(type) {
		case *OQueryResponse:
			r.Text = "AI Response"
		case *VerseResponse:
			r.Verse = "Passage Content"
		}
		return nil
	}

	tests := []struct {
		name          string
		userID        string
		message       string
		expectPassage bool
		expectAI      bool
		desc          string
	}{
		{
			name:          "Admin: Direct Question with Context",
			userID:        "99999",
			message:       "Explain John 3:16",
			expectPassage: false,
			expectAI:      true,
			desc:          "Admin should trigger AI query",
		},
		{
			name:          "Non-Admin: Direct Question with Context",
			userID:        "12345",
			message:       "Explain John 3:16",
			expectPassage: true,
			expectAI:      false, // Should fall back to passage
			desc:          "Non-Admin should NOT trigger AI query, but get passage",
		},
		{
			name:          "Admin: Natural Language Reference",
			userID:        "99999",
			message:       "I love John 3:16",
			expectPassage: false,
			expectAI:      true,
			desc:          "Admin chatting about verse should trigger AI",
		},
		{
			name:          "Non-Admin: Natural Language Reference",
			userID:        "12345",
			message:       "I love John 3:16",
			expectPassage: true,
			expectAI:      false,
			desc:          "Non-Admin chatting about verse should get passage",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Reset captured request
			capturedReq = QueryRequest{}

			env := def.SessionData{}
			env.User.Id = tt.userID
			env.Msg.Message = tt.message
			env = utils.SetUserConfig(env, `{"version":"NIV"}`)

			// Execute logic
			if strings.HasPrefix(tt.message, "Explain") {
				// ProcessNaturalLanguage handles this too if references are found
				ProcessNaturalLanguage(env)
			} else {
				ProcessNaturalLanguage(env)
			}

			// Verification
			if tt.expectAI {
				if len(capturedReq.Query.Prompt) == 0 {
					t.Errorf("Expected AI Query (Prompt) but got none")
				}
				if len(capturedReq.Query.Verses) > 0 {
					t.Errorf("Expected AI Query but got Passage Query (Verses: %v)", capturedReq.Query.Verses)
				}
			}

			if tt.expectPassage {
				if len(capturedReq.Query.Verses) == 0 {
					t.Errorf("Expected Passage Query (Verses) but got none")
				}
				if len(capturedReq.Query.Prompt) > 0 {
					t.Errorf("Expected Passage Query but got AI Query (Prompt: %s)", capturedReq.Query.Prompt)
				}
			}
		})
	}
}
