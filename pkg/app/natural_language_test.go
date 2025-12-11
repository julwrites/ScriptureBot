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
	ResetAPIConfigCache()

	tests := []struct {
		name          string
		message       string
		expectedCheck func(string) bool
		desc          string
	}{
		// Passage Scenarios
		{
			name:    "Passage: Reference",
			message: "John 3:16",
			expectedCheck: func(msg string) bool { return strings.Contains(msg, "John 3") || strings.Contains(msg, "loved the world") },
			desc: "Should retrieve John 3:16 passage",
		},
		{
			name:    "Passage: Short Book",
			message: "Jude",
			expectedCheck: func(msg string) bool { return strings.Contains(msg, "Jude") || strings.Contains(msg, "servant of Jesus Christ") },
			desc: "Should retrieve Jude passage",
		},
		{
			name:    "Passage: Book only",
			message: "Genesis",
			expectedCheck: func(msg string) bool { return strings.Contains(msg, "Genesis 1") },
			desc: "Should retrieve Genesis 1",
		},

		// Search Scenarios
		{
			name:    "Search: One word",
			message: "Grace",
			expectedCheck: func(msg string) bool { return strings.Contains(msg, "Found") || strings.Contains(msg, "No results") },
			desc: "Should perform search for Grace",
		},
		{
			name:    "Search: Short phrase",
			message: "Jesus wept",
			expectedCheck: func(msg string) bool { return strings.Contains(msg, "Found") },
			desc: "Should perform search for Jesus wept",
		},
		{
			name:    "Search: 3 words",
			message: "Love of God",
			expectedCheck: func(msg string) bool { return strings.Contains(msg, "Found") },
			desc: "Should perform search for Love of God",
		},

		// Ask Scenarios
		{
			name:    "Ask: Question",
			message: "What does the bible say about love?",
			expectedCheck: func(msg string) bool { return len(msg) == 0 },
			desc: "Should not ask the AI (Question)",
		},
		{
			name:    "Ask: With Reference",
			message: "Explain John 3:16",
			expectedCheck: func(msg string) bool { return !strings.Contains(msg, "Found") },
			desc: "Should ask the AI (With Reference)",
		},
		{
			name:    "Ask: Compare",
			message: "Compare Genesis 1 and John 1",
			expectedCheck: func(msg string) bool { return true },
			desc: "Should ask the AI (Compare)",
		},
		{
			name:    "Ask: Short Question",
			message: "Who is Jesus?",
			expectedCheck: func(msg string) bool { return len(msg) == 0 && !strings.Contains(msg, "Found") },
			desc: "Should not ask the AI (Short Question)",
		},
		{
			name:    "Ask: Embedded Reference",
			message: "What does it say in Mark 5?",
			expectedCheck: func(msg string) bool { return true },
			desc: "Should ask the AI (Embedded Reference)",
		},
		{
			name:    "Ask: Book name in text",
			message: "I like Genesis",
			expectedCheck: func(msg string) bool { return !strings.Contains(msg, "Found") },
			desc: "Should ask the AI (Found reference Genesis)",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			env := def.SessionData{}
			env.Msg.Message = tt.message
			env = utils.SetUserConfig(env, `{"version":"NIV"}`)

			res := ProcessNaturalLanguage(env)

			if !tt.expectedCheck(res.Res.Message) {
				t.Errorf("Response did not match expectation for %s. Got: %s", tt.desc, res.Res.Message)
			}
		})
	}
}
