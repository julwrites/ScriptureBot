package app

import (
	"testing"

	"github.com/julwrites/BotPlatform/pkg/def"
	"github.com/julwrites/ScriptureBot/pkg/utils"
)

func TestGetBibleAsk(t *testing.T) {
	// Restore original SubmitQuery after test
	originalSubmitQuery := SubmitQuery
	defer func() { SubmitQuery = originalSubmitQuery }()

	t.Run("Success: Verify Request", func(t *testing.T) {
		defer SetEnv("TELEGRAM_ADMIN_ID", "12345")()
		ResetAPIConfigCache()

		var capturedReq QueryRequest
		SubmitQuery = MockSubmitQuery(t, func(req QueryRequest) {
			capturedReq = req
		})

		var env def.SessionData
		env.Msg.Message = "Who is God?"
		env.User.Id = "12345"
		conf := utils.UserConfig{Version: "NIV"}
		env.User.Config = utils.SerializeUserConfig(conf)

		// Set dummy API config
		SetAPIConfigOverride("https://mock", "key")

		GetBibleAsk(env)

		if capturedReq.Query.Prompt != "Who is God?" {
			t.Errorf("Expected Query.Prompt to be 'Who is God?', got '%s'", capturedReq.Query.Prompt)
		}
		if len(capturedReq.Query.Verses) > 0 {
			t.Errorf("Expected Query.Verses to be empty, got %v", capturedReq.Query.Verses)
		}
		if len(capturedReq.Query.Words) > 0 {
			t.Errorf("Expected Query.Words to be empty, got %v", capturedReq.Query.Words)
		}
	})

	t.Run("Success: Verify Request with Context", func(t *testing.T) {
		ResetAPIConfigCache()

		var capturedReq QueryRequest
		SubmitQuery = MockSubmitQuery(t, func(req QueryRequest) {
			capturedReq = req
		})

		var env def.SessionData
		env.Msg.Message = "Explain this"
		conf := utils.UserConfig{Version: "NIV"}
		env.User.Config = utils.SerializeUserConfig(conf)
		contextVerses := []string{"John 3:16", "Genesis 1:1"}

		// Set dummy API config
		SetAPIConfigOverride("https://mock", "key")

		GetBibleAskWithContext(env, contextVerses)

		if capturedReq.Query.Prompt != "Explain this" {
			t.Errorf("Expected Query.Prompt to be 'Explain this', got '%s'", capturedReq.Query.Prompt)
		}
		if len(capturedReq.Context.Verses) != 2 {
			t.Errorf("Expected Context.Verses to have 2 items, got %v", capturedReq.Context.Verses)
		}
		if capturedReq.Context.Verses[0] != "John 3:16" {
			t.Errorf("Expected Context.Verses[0] to be 'John 3:16', got '%s'", capturedReq.Context.Verses[0])
		}
	})

	t.Run("Non-admin user", func(t *testing.T) {
		// Set admin ID to something else
		defer SetEnv("TELEGRAM_ADMIN_ID", "admin_id")()
		// Set mock API config so search works
		ResetAPIConfigCache()
		SetAPIConfigOverride("https://example.com", "api_key")
		SubmitQuery = originalSubmitQuery

		var env def.SessionData
		env.User.Id = "user_id"
		env.Msg.Message = "Question"
		conf := utils.UserConfig{Version: "NIV"}
		env.User.Config = utils.SerializeUserConfig(conf)

		env = GetBibleAsk(env)

		// Expect fallback to search
		expected := "Found 1 results for 'Question':\n- John 3:16\n"
		if env.Res.Message != expected {
			t.Errorf("Expected search result message, got: %s", env.Res.Message)
		}
	})

	t.Run("Admin user", func(t *testing.T) {
		defer SetEnv("TELEGRAM_ADMIN_ID", "admin_id")()
		ResetAPIConfigCache()
		SetAPIConfigOverride("https://example.com", "api_key")
		SubmitQuery = originalSubmitQuery

		var env def.SessionData
		env.User.Id = "admin_id"
		env.Msg.Message = "Question"
		conf := utils.UserConfig{Version: "NIV"}
		env.User.Config = utils.SerializeUserConfig(conf)

		env = GetBibleAsk(env)

		expected := "This is a mock response.\n\n*References:*\n- John 3:16"
		if env.Res.Message != expected {
			t.Errorf("Expected admin response, got: %s", env.Res.Message)
		}
	})
}
