package app

import (
	"strings"
	"testing"

	"github.com/julwrites/BotPlatform/pkg/def"
	"github.com/julwrites/ScriptureBot/pkg/utils"
)

func TestGetBibleSearch(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		defer SetEnv("BIBLE_API_URL", "https://example.com")()
		defer SetEnv("BIBLE_API_KEY", "api_key")()
		ResetAPIConfigCache()

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
