package app

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/julwrites/BotPlatform/pkg/def"
	"github.com/julwrites/ScriptureBot/pkg/utils"
)

func TestGetBibleAsk(t *testing.T) {
	// Mock server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var req QueryRequest
		json.NewDecoder(r.Body).Decode(&req)

		if req.Query.Prompt == "error" {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		resp := OQueryResponse{
			Text: "Answer text",
			References: []SearchResult{
				{Verse: "Ref 1:1", URL: "http://ref1"},
			},
		}
		json.NewEncoder(w).Encode(resp)
	}))
	defer ts.Close()

	os.Setenv("BIBLE_API_URL", ts.URL)
	defer os.Unsetenv("BIBLE_API_URL")

	t.Run("Success", func(t *testing.T) {
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
