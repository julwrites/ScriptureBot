package app

import (
	"strings"
	"testing"

	"github.com/julwrites/BotPlatform/pkg/def"
	"github.com/julwrites/ScriptureBot/pkg/utils"
)

func TestGetBiblePassageHtml(t *testing.T) {
	doc := GetPassageHTML("gen 8", "NIV")

	if doc == nil {
		t.Errorf("Could not retrieve bible passage")
	}
}

func TestGetReference(t *testing.T) {
	doc := GetPassageHTML("gen 1", "NIV")

	if doc == nil {
		t.Fatalf("Could not retrieve Bible passage for testing")
	}

	ref := GetReference(doc)

	if ref != "Genesis 1" {
		t.Errorf("Expected reference 'Genesis 1', but got '%s'", ref)
	}
}

func TestGetPassage(t *testing.T) {
	doc := GetPassageHTML("john 8", "NIV")

	passage := GetPassage("John 8", doc, "NIV")

	if len(passage) == 0 {
		t.Errorf("Failed TestGetPassage")
	}
}

func TestGetBiblePassage(t *testing.T) {
	// Restore original SubmitQuery after test
	originalSubmitQuery := SubmitQuery
	defer func() { SubmitQuery = originalSubmitQuery }()

	t.Run("Success: Verify Request", func(t *testing.T) {
		defer UnsetEnv("BIBLE_API_URL")()
		defer UnsetEnv("BIBLE_API_KEY")()
		ResetAPIConfigCache()

		var capturedReq QueryRequest
		SubmitQuery = MockSubmitQuery(t, func(req QueryRequest) {
			capturedReq = req
		})

		var env def.SessionData
		env.Msg.Message = "gen 1"
		var conf utils.UserConfig
		conf.Version = "NIV"
		env.User.Config = utils.SerializeUserConfig(conf)

		// Set dummy API config to pass internal checks
		SetAPIConfigOverride("https://mock", "key")

		GetBiblePassage(env)

		// Verify that Verses is populated and others are not
		if len(capturedReq.Query.Verses) != 1 || capturedReq.Query.Verses[0] != "Genesis 1" {
			t.Errorf("Expected Query.Verses to contain 'Genesis 1', got %v", capturedReq.Query.Verses)
		}
		if len(capturedReq.Query.Words) > 0 {
			t.Errorf("Expected Query.Words to be empty, got %v", capturedReq.Query.Words)
		}
		if capturedReq.Query.Prompt != "" {
			t.Errorf("Expected Query.Prompt to be empty, got '%s'", capturedReq.Query.Prompt)
		}
	})

	t.Run("Success: Response", func(t *testing.T) {
		defer UnsetEnv("BIBLE_API_URL")()
		defer UnsetEnv("BIBLE_API_KEY")()
		ResetAPIConfigCache()
		SetAPIConfigOverride("https://example.com", "key")
		SubmitQuery = originalSubmitQuery // Use default mock logic for response testing

		var env def.SessionData
		env.Msg.Message = "gen 1"
		var conf utils.UserConfig
		conf.Version = "NIV"
		env.User.Config = utils.SerializeUserConfig(conf)
		env = GetBiblePassage(env)

		if len(env.Res.Message) < 10 {
			t.Errorf("Expected passage text, got '%s'", env.Res.Message)
		}
	})

	t.Run("Empty", func(t *testing.T) {
		defer UnsetEnv("BIBLE_API_URL")()
		defer UnsetEnv("BIBLE_API_KEY")()
		ResetAPIConfigCache()

		var env def.SessionData
		env.Msg.Message = "nonexistentbook 99:99"
		env = GetBiblePassage(env)

		// Expecting some form of error message or empty fallback
		// If parsing fails, it might return empty string
		if len(env.Res.Message) > 0 && !strings.Contains(env.Res.Message, "No verses found") && !strings.Contains(env.Res.Message, "Sorry") {
			t.Errorf("Expected failure message, got '%s'", env.Res.Message)
		}
	})
}

func TestParsePassageFromHtml(t *testing.T) {
	t.Run("Valid HTML with superscript", func(t *testing.T) {
		html := `<p><span><sup>12 </sup>But to all who did receive him, who believed in his name, he gave the right to become children of God,</span></p>`
		expected := `^12 ^But to all who did receive him, who believed in his name, he gave the right to become children of God,`
		if got := ParsePassageFromHtml("", html, ""); got != expected {
			t.Errorf("ParsePassageFromHtml() = %v, want %v", got, expected)
		}
	})

	t.Run("HTML with italics", func(t *testing.T) {
		html := `<p><i>This is italic.</i></p>`
		expected := `_This is italic._`
		if got := ParsePassageFromHtml("", html, ""); got != expected {
			t.Errorf("ParsePassageFromHtml() = %v, want %v", got, expected)
		}
	})

	t.Run("HTML with bold", func(t *testing.T) {
		html := `<p><b>This is bold.</b></p>`
		expected := `*This is bold.*`
		if got := ParsePassageFromHtml("", html, ""); got != expected {
			t.Errorf("ParsePassageFromHtml() = %v, want %v", got, expected)
		}
	})

	t.Run("HTML with spans", func(t *testing.T) {
		html := `<p><span>Line 1.</span><br><span>    </span><span>Line 2.</span></p>`
		expected := "Line 1.\n\n    Line 2."
		if got := ParsePassageFromHtml("", html, ""); got != expected {
			t.Errorf("ParsePassageFromHtml() = %v, want %v", got, expected)
		}
	})

	t.Run("HTML with line breaks", func(t *testing.T) {
		html := `<p>Line 1.<br>Line 2.</p>`
		expected := "Line 1.\nLine 2."
		if got := ParsePassageFromHtml("", html, ""); got != expected {
			t.Errorf("ParsePassageFromHtml() = %v, want %v", got, expected)
		}
	})

	t.Run("Invalid HTML", func(t *testing.T) {
		html := `<p>This is malformed HTML`
		expected := `This is malformed HTML`
		if got := ParsePassageFromHtml("", html, ""); got != expected {
			t.Errorf("ParsePassageFromHtml() = %v, want %v", got, expected)
		}
	})

	t.Run("Nested HTML tags", func(t *testing.T) {
		html := `<p><b>This is bold, <i>and this is italic.</i></b></p>`
		expected := `*This is bold, _and this is italic._*`
		if got := ParsePassageFromHtml("", html, ""); got != expected {
			t.Errorf("ParsePassageFromHtml() = %v, want %v", got, expected)
		}
	})

	t.Run("MarkdownV2 escaping", func(t *testing.T) {
		// Note: We no longer escape explicitly in ParsePassageFromHtml as we rely on the platform
		// to handle it later (via PostTelegram).
		// However, returning raw characters like * might cause issues if not handled by platform.
		// For now, we expect them to be returned raw.
		html := `<p>This has special characters: *_. [hello](world)!</p>`
		expected := `This has special characters: *_. [hello](world)!`
		if got := ParsePassageFromHtml("", html, ""); got != expected {
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
