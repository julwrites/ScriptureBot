package app

import (
	"strings"
	"testing"

	"github.com/julwrites/BotPlatform/pkg/def"
	"github.com/julwrites/ScriptureBot/pkg/utils"
)

func TestGetBibleSearch(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		ResetAPIConfigCache()

		var env def.SessionData
		env.Msg.Message = "God"
		conf := utils.UserConfig{Version: "NIV"}
		env.User.Config = utils.SerializeUserConfig(conf)

		env = GetBibleSearch(env)

		if !strings.Contains(env.Res.Message, "Found") {
			t.Errorf("Expected result count, got: %s", env.Res.Message)
		}
	})
}
