package app

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"golang.org/x/net/html"

	"github.com/julwrites/BotPlatform/pkg/def"
	"github.com/julwrites/ScriptureBot/pkg/utils"
)

func setEnv(key, value string) func() {
	ResetAPIConfigCache()
	return utils.SetEnv(key, value)
}

func TestGetBiblePassageHtml(t *testing.T) {
	doc := GetPassageHTMLFunc("gen 8", "NIV")

	if doc == nil {
		t.Errorf("Could not retrieve bible passage")
	}
}

func TestGetReference(t *testing.T) {
	doc := GetPassageHTMLFunc("gen 1", "NIV")

	if doc == nil {
		t.Fatalf("Could not retrieve Bible passage for testing")
	}

	ref := GetReference(doc)

	if ref != "Genesis 1" {
		t.Errorf("Expected reference 'Genesis 1', but got '%s'", ref)
	}
}

func TestGetPassage(t *testing.T) {
	doc := GetPassageHTMLFunc("john 8", "NIV")

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

		verse := ""
		if len(req.Query.Verses) > 0 {
			verse = req.Query.Verses[0]
		}

		switch verse {
		case "gen 1":
			resp := VerseResponse{
				Verse: "<p>In the beginning God created the heavens and the earth.</p>",
			}
			json.NewEncoder(w).Encode(resp)
		case "empty":
			json.NewEncoder(w).Encode(VerseResponse{})
		default: // Any other case will trigger an error, forcing fallback
			w.WriteHeader(http.StatusInternalServerError)
		}
	}))
	defer ts.Close()

	defer setEnv("BIBLE_API_URL", ts.URL)()
	defer setEnv("BIBLE_API_KEY", "test_key")()

	t.Run("Success", func(t *testing.T) {
		var env def.SessionData
		env.Msg.Message = "gen 1"
		var conf utils.UserConfig
		conf.Version = "NIV"
		env.User.Config = utils.SerializeUserConfig(conf)
		env = GetBiblePassage(env)

		if env.Res.Message != `In the beginning God created the heavens and the earth\.` {
			t.Errorf("Expected passage text, got '%s'", env.Res.Message)
		}
	})

	t.Run("Fallback on API error", func(t *testing.T) {
		var env def.SessionData
		env.Msg.Message = "John 1:1" // Use a valid reference to test the fallback
		var conf utils.UserConfig
		conf.Version = "NIV"
		env.User.Config = utils.SerializeUserConfig(conf)

		// Mock the GetPassageHTMLFunc to avoid network calls.
		originalGetPassageHTML := GetPassageHTMLFunc
		defer func() { GetPassageHTMLFunc = originalGetPassageHTML }()
		GetPassageHTMLFunc = func(ref, ver string) *html.Node {
			// Return a mock HTML node.
			// The structure should be minimal but sufficient for GetPassage to parse.
			mockHTML := `<html><body><div class="passage-text"><div class="bcv">John 1:1</div><p>Mocked passage text.</p></div></body></html>`
			doc, _ := html.Parse(strings.NewReader(mockHTML))
			return doc
		}

		env = GetBiblePassage(env)

		// The fallback should now use the mocked function.
		if !strings.Contains(env.Res.Message, "Mocked passage text\\.") {
			t.Errorf("Expected fallback message to contain 'Mocked passage text.', got: '%s'", env.Res.Message)
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

func TestParsePassageFromHtml(t *testing.T) {
	t.Run("Valid HTML with superscript", func(t *testing.T) {
		html := `<p><span><sup>12 </sup>But to all who did receive him, who believed in his name, he gave the right to become children of God,</span></p>`
		expected := `^12^But to all who did receive him, who believed in his name, he gave the right to become children of God,`
		if got := ParsePassageFromHtml(html); got != expected {
			t.Errorf("ParsePassageFromHtml() = %v, want %v", got, expected)
		}
	})

	t.Run("HTML with italics", func(t *testing.T) {
		html := `<p><i>This is italic.</i></p>`
		expected := `_This is italic\._`
		if got := ParsePassageFromHtml(html); got != expected {
			t.Errorf("ParsePassageFromHtml() = %v, want %v", got, expected)
		}
	})

	t.Run("HTML with bold", func(t *testing.T) {
		html := `<p><b>This is bold.</b></p>`
		expected := `*This is bold\.*`
		if got := ParsePassageFromHtml(html); got != expected {
			t.Errorf("ParsePassageFromHtml() = %v, want %v", got, expected)
		}
	})

	t.Run("HTML with line breaks", func(t *testing.T) {
		html := `<p>Line 1.<br>Line 2.</p>`
		expected := "Line 1\\.\nLine 2\\."
		if got := ParsePassageFromHtml(html); got != expected {
			t.Errorf("ParsePassageFromHtml() = %v, want %v", got, expected)
		}
	})

	t.Run("Invalid HTML", func(t *testing.T) {
		html := `<p>This is malformed HTML`
		expected := `This is malformed HTML`
		if got := ParsePassageFromHtml(html); got != expected {
			t.Errorf("ParsePassageFromHtml() = %v, want %v", got, expected)
		}
	})

	t.Run("Nested HTML tags", func(t *testing.T) {
		html := `<p><b>This is bold, <i>and this is italic.</i></b></p>`
		expected := `*This is bold, *_and this is italic\._`
		if got := ParsePassageFromHtml(html); got != expected {
			t.Errorf("ParsePassageFromHtml() = %v, want %v", got, expected)
		}
	})

	t.Run("MarkdownV2 escaping", func(t *testing.T) {
		html := `<p>This has special characters: *_. [hello](world)!</p>`
		expected := `This has special characters: \*\_\. \[hello\]\(world\)\!`
		if got := ParsePassageFromHtml(html); got != expected {
			t.Errorf("ParsePassageFromHtml() = %v, want %v", got, expected)
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
