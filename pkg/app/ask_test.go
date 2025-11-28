package app

import (
	"testing"

	"github.com/julwrites/BotPlatform/pkg/def"
	"github.com/julwrites/ScriptureBot/pkg/utils"
)

func TestGetBibleAsk(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		defer UnsetEnv("BIBLE_API_URL")()
		defer UnsetEnv("BIBLE_API_KEY")()
		ResetAPIConfigCache()

		var env def.SessionData
		env.Msg.Message = "Question"
		conf := utils.UserConfig{Version: "NIV"}
		env.User.Config = utils.SerializeUserConfig(conf)

		env = GetBibleAsk(env)

		if len(env.Res.Message) == 0 {
			t.Errorf("Expected answer text, got empty")
		}
	})
}
