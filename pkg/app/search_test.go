package app

import (
	"strings"
	"testing"

	"github.com/julwrites/BotPlatform/pkg/def"
	"github.com/julwrites/ScriptureBot/pkg/utils"
)

func TestGetBibleSearch(t *testing.T) {
	// Restore original SubmitQuery after test
	originalSubmitQuery := SubmitQuery
	defer func() { SubmitQuery = originalSubmitQuery }()

	t.Run("Success: Verify Request", func(t *testing.T) {
		ResetAPIConfigCache()

		var capturedReq QueryRequest
		SubmitQuery = MockSubmitQuery(t, func(req QueryRequest) {
			capturedReq = req
		})

		var env def.SessionData
		env.Msg.Message = "God is love"
		conf := utils.UserConfig{Version: "NIV"}
		env.User.Config = utils.SerializeUserConfig(conf)

		// Set dummy API config to pass internal checks
		SetAPIConfigOverride("https://mock", "key")

		GetBibleSearch(env)

		// Verify that Words is populated and others are not
		expectedWords := []string{"God", "is", "love"}
		if len(capturedReq.Query.Words) != 3 {
			t.Errorf("Expected Query.Words to have 3 items, got %v", capturedReq.Query.Words)
		} else {
			for i, word := range capturedReq.Query.Words {
				if word != expectedWords[i] {
					t.Errorf("Expected word '%s' at index %d, got '%s'", expectedWords[i], i, word)
				}
			}
		}

		if len(capturedReq.Query.Verses) > 0 {
			t.Errorf("Expected Query.Verses to be empty, got %v", capturedReq.Query.Verses)
		}
		if capturedReq.Query.Prompt != "" {
			t.Errorf("Expected Query.Prompt to be empty, got '%s'", capturedReq.Query.Prompt)
		}
	})

	t.Run("Success: Response", func(t *testing.T) {
		defer SetEnv("BIBLE_API_URL", "https://example.com")()
		defer SetEnv("BIBLE_API_KEY", "api_key")()
		ResetAPIConfigCache()
		SubmitQuery = originalSubmitQuery // Use default mock logic

		var env def.SessionData
		env.Msg.Message = "God"
		conf := utils.UserConfig{Version: "NIV"}
		env.User.Config = utils.SerializeUserConfig(conf)

		env = GetBibleSearch(env)

		if !strings.Contains(env.Res.Message, "Found") && !strings.Contains(env.Res.Message, "No results") {
			t.Errorf("Expected result count, got: %s", env.Res.Message)
		}
	})
}
