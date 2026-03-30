package app

import (
	"strings"
	"testing"

	"github.com/julwrites/BotPlatform/pkg/def"
	"github.com/julwrites/ScriptureBot/pkg/utils"
)

func TestProcessNaturalLanguage(t *testing.T) {
	// Set dummy API keys to prevent real API calls
	defer SetEnv("BIBLE_API_URL", "https://example.com")()
	defer SetEnv("BIBLE_API_KEY", "api_key")()
	defer SetEnv("TELEGRAM_ADMIN_ID", "12345")()
	ResetAPIConfigCache()

	tests := []struct {
		name          string
		message       string
		isAdmin       bool
		expectedCheck func(string) bool
		desc          string
	}{
		// Passage Scenarios
		{
			name:    "Passage: Reference",
			message: "John 3:16",
			isAdmin: true,
			expectedCheck: func(msg string) bool {
				return strings.Contains(msg, "John 3") || strings.Contains(msg, "loved the world")
			},
			desc: "Should retrieve John 3:16 passage",
		},
		{
			name:    "Passage: Short Book",
			message: "Jude",
			isAdmin: true,
			expectedCheck: func(msg string) bool {
				return strings.Contains(msg, "Jude") || strings.Contains(msg, "servant of Jesus Christ")
			},
			desc: "Should retrieve Jude passage",
		},
		{
			name:          "Passage: Book only",
			message:       "Genesis",
			isAdmin:       true,
			expectedCheck: func(msg string) bool { return strings.Contains(msg, "Genesis 1") },
			desc:          "Should retrieve Genesis 1",
		},

		// Search Scenarios
		{
			name:          "Search: One word",
			message:       "Grace",
			isAdmin:       true,
			expectedCheck: func(msg string) bool { return strings.Contains(msg, "Found") || strings.Contains(msg, "No results") },
			desc:          "Should perform search for Grace",
		},
		{
			name:          "Search: Short phrase",
			message:       "Jesus wept",
			isAdmin:       true,
			expectedCheck: func(msg string) bool { return strings.Contains(msg, "Found") },
			desc:          "Should perform search for Jesus wept",
		},
		{
			name:          "Search: 3 words",
			message:       "Love of God",
			isAdmin:       true,
			expectedCheck: func(msg string) bool { return strings.Contains(msg, "Found") },
			desc:          "Should perform search for Love of God",
		},

		// Ask Scenarios (Admin)
		{
			name:          "Ask: Question (Admin)",
			message:       "What does the bible say about love?",
			isAdmin:       true,
			expectedCheck: func(msg string) bool { return strings.Contains(msg, "mock response") },
			desc:          "Should ask the AI (Question)",
		},
		{
			name:          "Ask: With Reference (Admin)",
			message:       "Explain John 3:16",
			isAdmin:       true,
			expectedCheck: func(msg string) bool { return strings.Contains(msg, "mock response") },
			desc:          "Should ask the AI (With Reference)",
		},
		{
			name:          "Ask: Compare (Admin)",
			message:       "Compare Genesis 1 and John 1",
			isAdmin:       true,
			expectedCheck: func(msg string) bool { return strings.Contains(msg, "mock response") },
			desc:          "Should ask the AI (Compare)",
		},
		{
			name:          "Ask: Short Question (Admin)",
			message:       "Who is Jesus?",
			isAdmin:       true,
			expectedCheck: func(msg string) bool { return strings.Contains(msg, "mock response") },
			desc:          "Should ask the AI (Short Question)",
		},
		{
			name:          "Ask: Embedded Reference (Admin)",
			message:       "What does it say in Mark 5?",
			isAdmin:       true,
			expectedCheck: func(msg string) bool { return strings.Contains(msg, "mock response") },
			desc:          "Should ask the AI (Embedded Reference)",
		},
		{
			name:          "Ask: Book name in text (Admin)",
			message:       "I like Genesis",
			isAdmin:       true,
			expectedCheck: func(msg string) bool { return strings.Contains(msg, "mock response") },
			desc:          "Should ask the AI (Found reference Genesis)",
		},

		// Ask Scenarios (Non-Admin)
		{
			name:          "Ask: Question (Non-Admin)",
			message:       "What does the bible say about love?",
			isAdmin:       false,
			expectedCheck: func(msg string) bool { return strings.Contains(msg, "I'm sorry, I didn't understand that") },
			desc:          "Should return fallback message for non-admin question",
		},
		{
			name:    "Ask: With Reference (Non-Admin)",
			message: "Explain John 3:16",
			isAdmin: false,
			// Since "John 3:16" is extracted, non-admin falls back to GetBiblePassage for the first reference.
			expectedCheck: func(msg string) bool { return strings.Contains(msg, "John 3") || strings.Contains(msg, "loved the world") },
			desc:          "Should retrieve passage for non-admin with reference",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			env := def.SessionData{}
			if tt.isAdmin {
				env.User.Id = "12345"
			} else {
				env.User.Id = "67890"
			}
			env.Msg.Message = tt.message
			env = utils.SetUserConfig(env, `{"version":"NIV"}`)

			res := ProcessNaturalLanguage(env)

			if !tt.expectedCheck(res.Res.Message) {
				t.Errorf("Response did not match expectation for %s. Got: %s", tt.desc, res.Res.Message)
			}
		})
	}
}
