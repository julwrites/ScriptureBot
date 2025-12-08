package app

import (
	"testing"

	"github.com/julwrites/BotPlatform/pkg/def"
	"github.com/julwrites/ScriptureBot/pkg/utils"
)

func TestGetBibleAsk(t *testing.T) {
	t.Run("Non-admin user", func(t *testing.T) {
		// Set admin ID to something else
		defer SetEnv("TELEGRAM_ADMIN_ID", "admin_id")()
		// Set mock API config so search works
		ResetAPIConfigCache()
		SetAPIConfigOverride("https://example.com", "api_key")

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
