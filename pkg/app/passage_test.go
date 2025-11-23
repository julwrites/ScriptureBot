package app

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/julwrites/BotPlatform/pkg/def"
	"github.com/julwrites/ScriptureBot/pkg/utils"
)

func TestGetBiblePassageHtml(t *testing.T) {
	doc := GetPassageHtml("gen 8", "NIV")

	if doc == nil {
		t.Errorf("Could not retrieve bible passage")
	}
}

func TestGetReference(t *testing.T) {
	doc := GetPassageHtml("gen 1", "NIV")

	if doc == nil {
		t.Fatalf("Could not retrieve Bible passage for testing")
	}

	ref := GetReference(doc)

	if ref != "Genesis 1" {
		t.Errorf("Expected reference 'Genesis 1', but got '%s'", ref)
	}
}

func TestGetPassage(t *testing.T) {
	doc := GetPassageHtml("john 8", "NIV")

	passage := GetPassage("John 8", doc, "NIV")

	if len(passage) == 0 {
		t.Errorf("Failed TestGetPassage")
	}
}

func TestGetBiblePassage(t *testing.T) {
	// Mock server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var req QueryRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}

		if len(req.Query.Verses) > 0 && req.Query.Verses[0] == "error" {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if len(req.Query.Verses) > 0 && req.Query.Verses[0] == "empty" {
			json.NewEncoder(w).Encode(VerseResponse{})
			return
		}

		resp := VerseResponse{
			Verse: "In the beginning God created the heavens and the earth.",
		}
		json.NewEncoder(w).Encode(resp)
	}))
	defer ts.Close()

	os.Setenv("BIBLE_API_URL", ts.URL)
	defer os.Unsetenv("BIBLE_API_URL")

	t.Run("Success", func(t *testing.T) {
		var env def.SessionData
		env.Msg.Message = "gen 1"
		var conf utils.UserConfig
		conf.Version = "NIV"
		env.User.Config = utils.SerializeUserConfig(conf)
		env = GetBiblePassage(env)

		if env.Res.Message != "In the beginning God created the heavens and the earth." {
			t.Errorf("Expected passage text, got '%s'", env.Res.Message)
		}
	})

	t.Run("Error", func(t *testing.T) {
		var env def.SessionData
		env.Msg.Message = "error"
		env = GetBiblePassage(env)

		if env.Res.Message != "Sorry, I couldn't retrieve that passage. Please check the reference or try again later." {
			t.Errorf("Expected error message, got '%s'", env.Res.Message)
		}
	})

	t.Run("Empty", func(t *testing.T) {
		var env def.SessionData
		env.Msg.Message = "empty"
		env = GetBiblePassage(env)

		if env.Res.Message != "No verses found." {
			t.Errorf("Expected empty message, got '%s'", env.Res.Message)
		}
	})
}

func TestCheckBibleReference(t *testing.T) {
	if CheckBibleReference("Genesis 1:1") == false {
		t.Errorf("Failed CheckBibleReference positive test")
	}

	if CheckBibleReference("Some terrible other word") == true {
		t.Errorf("Failed CheckBibleReference negative test")
	}
}
