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

		var env def.SessionData
		env.User.Id = "user_id"
		env.Msg.Message = "Question"
		conf := utils.UserConfig{Version: "NIV"}
		env.User.Config = utils.SerializeUserConfig(conf)

		env = GetBibleAsk(env)

		if env.Res.Message != "Sorry, this feature is only available to the administrator." {
			t.Errorf("Expected permission denied message, got: %s", env.Res.Message)
		}
	})

	t.Run("Admin user", func(t *testing.T) {
		defer SetEnv("TELEGRAM_ADMIN_ID", "admin_id")()
		defer SetEnv("BIBLE_API_URL", "https://example.com")()
		defer SetEnv("BIBLE_API_KEY", "api_key")()
		ResetAPIConfigCache()

		var env def.SessionData
		env.User.Id = "admin_id"
		env.Msg.Message = "Question"
		conf := utils.UserConfig{Version: "NIV"}
		env.User.Config = utils.SerializeUserConfig(conf)

		// This will still fail because it makes a real API call
		// but it will pass the admin check
		env = GetBibleAsk(env)

		if env.Res.Message == "Sorry, this feature is only available to the administrator." {
			t.Errorf("Expected to pass admin check, but it failed")
		}
	})
}
