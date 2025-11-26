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
	handler := newMockApiHandler()
	ts := httptest.NewServer(handler)
	defer ts.Close()

	t.Run("Success", func(t *testing.T) {
		defer setEnv("BIBLE_API_URL", ts.URL)()
		ResetAPIConfigCache()

		var env def.SessionData
		env.Msg.Message = "Question"
		conf := utils.UserConfig{Version: "NIV"}
		env.User.Config = utils.SerializeUserConfig(conf)

		env = GetBibleAsk(env)

		if !strings.Contains(env.Res.Message, "Answer text") {
			t.Errorf("Expected answer text, got: %s", env.Res.Message)
		}
		if !strings.Contains(env.Res.Message, "Ref 1:1") {
			t.Errorf("Expected reference, got: %s", env.Res.Message)
		}
	})

	t.Run("Error", func(t *testing.T) {
		handler.statusCode = http.StatusInternalServerError
		defer func() { handler.statusCode = http.StatusOK }()

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
