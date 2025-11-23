package bot

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/julwrites/BotPlatform/pkg/def"
	"github.com/julwrites/ScriptureBot/pkg/app"
	"github.com/julwrites/ScriptureBot/pkg/utils"
)

func TestRunCommands(t *testing.T) {
	// Mock server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var req app.QueryRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}

		resp := app.VerseResponse{
			Verse: "Not so the wicked!",
		}
		json.NewEncoder(w).Encode(resp)
	}))
	defer ts.Close()

	// Override API config
	app.SetAPIConfigOverride(ts.URL, "dummy")

	var env def.SessionData
	var conf utils.UserConfig
	conf.Version = "NIV"
	env.User.Config = utils.SerializeUserConfig(conf)
	env.Msg.Message = "psalm 1"

	env = RunCommands(env)

	if !strings.Contains(env.Res.Message, "Not so the wicked!") {
		t.Errorf("Failed TestRunCommands Passage command. Got: %s", env.Res.Message)
	}
}

func TestUserCheck(t *testing.T) {
	var user def.UserData

	user.Firstname = "User"
	user.Lastname = ""
	user.Username = "username"
	user.Id = "id"
	user.Type = "Individual"
	user.Action = ""
	user.Config = ""

	cache := user

	if user != cache {
		t.Errorf("Failed TestUserCheck identical user")
	}

	user.Lastname = "name"

	if user == cache {
		t.Errorf("Failed TestUserCheck changed user")
	}
}
