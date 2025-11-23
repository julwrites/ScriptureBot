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

func TestGetBibleSearch(t *testing.T) {
	// Mock server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var req QueryRequest
		json.NewDecoder(r.Body).Decode(&req)

		// Check if words contains "error"
		for _, word := range req.Query.Words {
			if word == "error" {
				http.Error(w, "Error", http.StatusInternalServerError)
				return
			}
			if word == "empty" {
				json.NewEncoder(w).Encode(WordSearchResponse{})
				return
			}
		}

		resp := WordSearchResponse{
			{Verse: "Found 1:1", URL: "http://found1"},
		}
		json.NewEncoder(w).Encode(resp)
	}))
	defer ts.Close()

	os.Setenv("BIBLE_API_URL", ts.URL)
	defer os.Unsetenv("BIBLE_API_URL")

	t.Run("Success", func(t *testing.T) {
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
		var env def.SessionData
		env.Msg.Message = "empty"

		env = GetBibleSearch(env)

		if !strings.Contains(env.Res.Message, "No results found") {
			t.Errorf("Expected no results message, got: %s", env.Res.Message)
		}
	})

	t.Run("Error", func(t *testing.T) {
		var env def.SessionData
		env.Msg.Message = "error"

		env = GetBibleSearch(env)

		if !strings.Contains(env.Res.Message, "Sorry") {
			t.Errorf("Expected error message, got: %s", env.Res.Message)
		}
	})
}
