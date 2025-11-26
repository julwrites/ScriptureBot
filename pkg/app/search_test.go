package app

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/julwrites/BotPlatform/pkg/def"
	"github.com/julwrites/ScriptureBot/pkg/utils"
)

func TestGetBibleSearch(t *testing.T) {
	handler := newMockApiHandler()
	ts := httptest.NewServer(handler)
	defer ts.Close()

	t.Run("Success", func(t *testing.T) {
		defer setEnv("BIBLE_API_URL", ts.URL)()
		ResetAPIConfigCache()

		var env def.SessionData
		env.Msg.Message = "Found"
		conf := utils.UserConfig{Version: "NIV"}
		env.User.Config = utils.SerializeUserConfig(conf)

		env = GetBibleSearch(env)

		if !strings.Contains(env.Res.Message, "Found 1 results") {
			t.Errorf("Expected result count, got: %s", env.Res.Message)
		}
		if !strings.Contains(env.Res.Message, "Found 1:1") {
			t.Errorf("Expected verse ref, got: %s", env.Res.Message)
		}
	})

	t.Run("Empty", func(t *testing.T) {
		handler.wordSearchResponse = WordSearchResponse{}
		defer func() {
			handler.wordSearchResponse = WordSearchResponse{
				{Verse: "Found 1:1", URL: "http://found1"},
			}
		}()

		defer setEnv("BIBLE_API_URL", ts.URL)()
		ResetAPIConfigCache()

		var env def.SessionData
		env.Msg.Message = "empty"

		env = GetBibleSearch(env)

		if !strings.Contains(env.Res.Message, "No results found") {
			t.Errorf("Expected no results message, got: %s", env.Res.Message)
		}
	})

	t.Run("Error", func(t *testing.T) {
		handler.statusCode = http.StatusInternalServerError
		defer func() { handler.statusCode = http.StatusOK }()

		defer setEnv("BIBLE_API_URL", ts.URL)()
		ResetAPIConfigCache()

		var env def.SessionData
		env.Msg.Message = "error"

		env = GetBibleSearch(env)

		if !strings.Contains(env.Res.Message, "Sorry") {
			t.Errorf("Expected error message, got: %s", env.Res.Message)
		}
	})
}
