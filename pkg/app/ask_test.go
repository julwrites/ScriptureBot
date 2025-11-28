package app

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/julwrites/BotPlatform/pkg/def"
	"github.com/julwrites/ScriptureBot/pkg/utils"
)

func TestGetBibleAsk(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
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

	t.Run("Error", func(t *testing.T) {
		handler := newMockApiHandler()
		ts := httptest.NewServer(handler)
		defer ts.Close()

		handler.statusCode = http.StatusInternalServerError

		defer setEnv("BIBLE_API_URL", ts.URL)()
		ResetAPIConfigCache()

		var env def.SessionData
		env.Msg.Message = "error"
		conf := utils.UserConfig{Version: "NIV"}
		env.User.Config = utils.SerializeUserConfig(conf)

		env = GetBibleAsk(env)

		if !strings.Contains(env.Res.Message, "Sorry") {
			t.Errorf("Expected error message, got: %s", env.Res.Message)
		}
	})
}
