package app

import (
	"errors"
	"strings"
	"testing"

	"golang.org/x/net/html"

	"github.com/julwrites/BotPlatform/pkg/def"
	"github.com/julwrites/ScriptureBot/pkg/utils"
)

func mockGetPassageHTML(htmlStr string) *html.Node {
	doc, _ := html.Parse(strings.NewReader(htmlStr))
	return doc
}

func TestGetReference(t *testing.T) {
	// Mock GetPassageHTML
	originalGetPassageHTML := GetPassageHTML
	defer func() { GetPassageHTML = originalGetPassageHTML }()

	GetPassageHTML = func(ref, ver string) *html.Node {
		return mockGetPassageHTML(`
			<div class="bcv">Genesis 1</div>
		`)
	}

	doc := GetPassageHTML("gen 1", "NIV")
	ref := GetReference(doc)

	if ref != "Genesis 1" {
		t.Errorf("Expected reference 'Genesis 1', but got '%s'", ref)
	}
}

func TestGetPassage(t *testing.T) {
	// Mock GetPassageHTML
	originalGetPassageHTML := GetPassageHTML
	defer func() { GetPassageHTML = originalGetPassageHTML }()

	GetPassageHTML = func(ref, ver string) *html.Node {
		return mockGetPassageHTML(`
			<div class="passage-text">
				<p>In the beginning was the Word.</p>
			</div>
		`)
	}

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
		env = utils.SetUserConfig(env, utils.SerializeUserConfig(conf))

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
		if capturedReq.User.Version != "NIV" {
			t.Errorf("Expected User.Version to be 'NIV', got '%s'", capturedReq.User.Version)
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
		env = utils.SetUserConfig(env, utils.SerializeUserConfig(conf))
		env = GetBiblePassage(env)

		if len(env.Res.Message) < 10 {
			t.Errorf("Expected passage text, got '%s'", env.Res.Message)
		}
		// Verify ParseMode is set
		if env.Res.ParseMode != "HTML" {
			t.Errorf("Expected ParseMode 'HTML', got '%s'", env.Res.ParseMode)
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

	t.Run("Fallback: Scrape", func(t *testing.T) {
		defer UnsetEnv("BIBLE_API_URL")()
		defer UnsetEnv("BIBLE_API_KEY")()
		ResetAPIConfigCache()

		// Mock GetPassageHTML for fallback
		originalGetPassageHTML := GetPassageHTML
		defer func() { GetPassageHTML = originalGetPassageHTML }()

		GetPassageHTML = func(ref, ver string) *html.Node {
			return mockGetPassageHTML(`
				<div class="bcv">Genesis 1</div>
				<div class="passage-text">
					<p>In the beginning God created the heavens and the earth.</p>
				</div>
			`)
		}

		var env def.SessionData
		env.Msg.Message = "gen 1"
		var conf utils.UserConfig
		conf.Version = "NIV"
		env = utils.SetUserConfig(env, utils.SerializeUserConfig(conf))

		// Override SubmitQuery to force failure
		originalSubmitQuerySub := SubmitQuery
		defer func() { SubmitQuery = originalSubmitQuerySub }()
		SubmitQuery = func(req QueryRequest, result interface{}) error {
			return errors.New("forced api error")
		}

		env = GetBiblePassage(env)

		if !strings.Contains(env.Res.Message, "In the beginning") {
			t.Errorf("Expected fallback passage content, got '%s'", env.Res.Message)
		}
		// Fallback should also use HTML mode
		if env.Res.ParseMode != "HTML" {
			t.Errorf("Expected ParseMode 'HTML' in fallback, got '%s'", env.Res.ParseMode)
		}
	})
}

func TestParsePassageFromHtml(t *testing.T) {
	t.Run("Valid HTML with superscript", func(t *testing.T) {
		html := `<p><span><sup>12 </sup>But to all who did receive him, who believed in his name, he gave the right to become children of God,</span></p>`
		// Updated expectation: unicode superscripts and HTML formatting
		expected := `¹²But to all who did receive him, who believed in his name, he gave the right to become children of God,`
		if got := ParsePassageFromHtml("", html, ""); got != expected {
			t.Errorf("ParsePassageFromHtml() = %s, want %s", got, expected)
		}
	})

	t.Run("HTML with italics", func(t *testing.T) {
		html := `<p><i>This is italic.</i></p>`
		expected := `<i>This is italic.</i>`
		if got := ParsePassageFromHtml("", html, ""); got != expected {
			t.Errorf("ParsePassageFromHtml() = %v, want %v", got, expected)
		}
	})

	t.Run("HTML with bold", func(t *testing.T) {
		html := `<p><b>This is bold.</b></p>`
		expected := `<b>This is bold.</b>`
		if got := ParsePassageFromHtml("", html, ""); got != expected {
			t.Errorf("ParsePassageFromHtml() = %v, want %v", got, expected)
		}
	})

	t.Run("HTML with spans", func(t *testing.T) {
		html := `<p><span>Line 1.</span><br><span>    </span><span>Line 2.</span></p>`
		expected := "Line 1.\n    Line 2."
		if got := ParsePassageFromHtml("", html, ""); got != expected {
			t.Errorf("ParsePassageFromHtml() = %v, want %v", got, expected)
		}
	})

	t.Run("Poetry double newline check", func(t *testing.T) {
		html := `<p><span>Line 1</span><br><span>Line 2</span></p>`
		expected := "Line 1\nLine 2"
		if got := ParsePassageFromHtml("", html, ""); got != expected {
			t.Errorf("ParsePassageFromHtml() = %q, want %q", got, expected)
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
		expected := `<b>This is bold, <i>and this is italic.</i></b>`
		if got := ParsePassageFromHtml("", html, ""); got != expected {
			t.Errorf("ParsePassageFromHtml() = %v, want %v", got, expected)
		}
	})

	t.Run("Lists", func(t *testing.T) {
		html := `<ul><li>Item 1</li><li>Item 2</li></ul>`
		// Note: The ParseNodesForPassage appends newline after each Item.
		// strings.TrimSpace removes the last newline.
		// Item 1\nItem 2\n -> Item 1\nItem 2
		expected := "• Item 1\n• Item 2"
		if got := ParsePassageFromHtml("", html, ""); got != expected {
			t.Errorf("ParsePassageFromHtml() = %q, want %q", got, expected)
		}
	})

	t.Run("Headers", func(t *testing.T) {
		html := `<h1>Header</h1>`
		// Code: \n\n<b>Header</b>\n
		// TrimSpace -> <b>Header</b>
		expected := "<b>Header</b>"
		if got := ParsePassageFromHtml("", html, ""); got != expected {
			t.Errorf("ParsePassageFromHtml() = %q, want %q", got, expected)
		}
	})

	t.Run("Divs and escaping", func(t *testing.T) {
		html := `<div>Text &lt;with&gt; symbols</div>`
		expected := "Text &lt;with&gt; symbols"
		if got := ParsePassageFromHtml("", html, ""); got != expected {
			t.Errorf("ParsePassageFromHtml() = %q, want %q", got, expected)
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
