// Brief: Bible Passage handling
// Primary responsibility: Parsing of Bible Passages from HTML

package app

import (
	"fmt"
	"log"
	"net/url"
	"strings"

	"golang.org/x/net/html"

	"github.com/julwrites/BotPlatform/pkg/def"
	"github.com/julwrites/BotPlatform/pkg/platform"
	"github.com/julwrites/ScriptureBot/pkg/utils"
)

// GetPassageHTMLFunc is a variable to allow mocking in tests.
// Deprecated: Using new API service
var GetPassageHTMLFunc = func(ref, ver string) *html.Node {
	ref = url.QueryEscape(ref)
	ver = url.QueryEscape(ver)
	query := fmt.Sprintf("https://classic.biblegateway.com/passage/?search=%s&version=%s&interface=print", ref, ver)

	return utils.QueryHtml(query)
}

// Deprecated: Using new API service
func GetReference(doc *html.Node) string {
	refNode, err := utils.FindByClass(doc, "bcv")
	if err != nil {
		log.Printf("Error parsing for reference: %v", err)
		return ""
	}

	return utils.GetTextNode(refNode).Data
}

// Helper functions for parsing
func isFormattingTag(tag string) bool {
	return tag == "sup" || tag == "i" || tag == "b"
}

func isHeaderTag(tag string) bool {
	return tag == "h1" || tag == "h2" || tag == "h3" || tag == "h4"
}

func wrapText(text, tag string) string {
	if strings.TrimSpace(text) == "" {
		return text
	}

	if tag == "sup" {
		return platform.TelegramSuperscript(strings.Trim(text, " "))
	}
	if tag == "i" {
		return platform.TelegramItalics(text)
	}
	if tag == "b" || isHeaderTag(tag) {
		return platform.TelegramBold(text)
	}
	return text
}

func parseNode(node *html.Node) string {
	if node.Type == html.TextNode {
		return node.Data
	}

	if node.Type != html.ElementNode {
		var content strings.Builder
		for c := node.FirstChild; c != nil; c = c.NextSibling {
			content.WriteString(parseNode(c))
		}
		return content.String()
	}

	tag := node.Data

	// Handle non-formatting tags first
	if tag == "br" {
		return "\n"
	}

	// Treat headers and paragraphs as block elements
	if tag == "p" || isHeaderTag(tag) {
		var content strings.Builder
		content.WriteString("\n")

		// Buffer to hold content of the block
		var blockContent strings.Builder

		for c := node.FirstChild; c != nil; c = c.NextSibling {
			blockContent.WriteString(parseNode(c))
		}

		if isHeaderTag(tag) {
			content.WriteString(platform.TelegramBold(blockContent.String()))
		} else {
			content.WriteString(blockContent.String())
		}

		content.WriteString("\n")
		return content.String()
	}

	if !isFormattingTag(tag) {
		var content strings.Builder
		for c := node.FirstChild; c != nil; c = c.NextSibling {
			content.WriteString(parseNode(c))
		}
		return content.String()
	}

	// Handle formatting tags (b, i, sup, h1-h4)
	if tag == "sup" {
		for _, attr := range node.Attr {
			if attr.Key == "class" && attr.Val == "footnote" {
				return "" // Ignore footnote nodes
			}
		}
	}

	var content strings.Builder
	var textBuffer strings.Builder

	flushTextBuffer := func() {
		if textBuffer.Len() > 0 {
			content.WriteString(wrapText(textBuffer.String(), tag))
			textBuffer.Reset()
		}
	}

	for c := node.FirstChild; c != nil; c = c.NextSibling {
		// Note: isHeaderTag is removed here because it's handled above as a block element
		if c.Type == html.ElementNode && isFormattingTag(c.Data) {
			flushTextBuffer()
			content.WriteString(parseNode(c))
		} else {
			textBuffer.WriteString(parseNode(c))
		}
	}
	flushTextBuffer()

	return content.String()
}

func ParsePassageFromHtml(rawHtml string) string {
	doc, err := html.Parse(strings.NewReader(rawHtml))
	if err != nil {
		log.Printf("Error parsing html: %v", err)
		return rawHtml
	}
	return strings.TrimSpace(parseNode(doc))
}

// Deprecated: Using new API service
func GetPassage(ref string, doc *html.Node, version string) string {
	passageNode, startErr := utils.FindByClass(doc, "passage-text")
	if startErr != nil {
		log.Printf("Error parsing for passage: %v", startErr)
		return ""
	}

	filtNodes := utils.FilterTree(passageNode, func(child *html.Node) bool {
		switch tag := child.Data; tag {
		case "h1":
			fallthrough
		case "h2":
			fallthrough
		case "h3":
			fallthrough
		case "h4":
			if child.FirstChild.Data == "Footnotes" || child.FirstChild.Data == "Cross references" {
				return false
			}
			fallthrough
		case "p":
			return true
		}
		return false
	})

	textBlocks := utils.MapNodeListToString(filtNodes, parseNode)

	var passage strings.Builder

	refString := fmt.Sprintf("_%s_ (%s)", ref, version)
	passage.WriteString(refString)

	for _, block := range textBlocks {
		passage.WriteString("\n")
		passage.WriteString(block)
	}

	return passage.String()
}

func GetBiblePassage(env def.SessionData) def.SessionData {
	if len(env.Msg.Message) > 0 {
		config := utils.DeserializeUserConfig(env.User.Config)

		req := QueryRequest{
			Query: QueryObject{
				Verses: []string{env.Msg.Message},
			},
			Context: QueryContext{
				User: UserContext{
					Version: config.Version,
				},
			},
		}

		var resp VerseResponse
		err := SubmitQuery(req, &resp, env.Secrets.PROJECT_ID)
		if err != nil {
			log.Printf("Error retrieving passage from API: %v. Falling back to deprecated method.", err)
			// Fallback to deprecated passage retrieval logic
			doc := GetPassageHTMLFunc(env.Msg.Message, config.Version)
			if doc == nil {
				env.Res.Message = "Sorry, I couldn't retrieve that passage. Please check the reference or try again later."
				return env
			}
			env.Res.Message = GetPassage(GetReference(doc), doc, config.Version)
			return env
		}

		if len(resp.Verse) > 0 {
			env.Res.Message = ParsePassageFromHtml(resp.Verse)
		} else {
			env.Res.Message = "No verses found."
		}
	}

	return env
}

// Deprecated: Using new API service logic inside GetBiblePassage
// Deprecated: Using new API service
func CheckBibleReference(ref string) bool {
	log.Printf("Checking reference %s", ref)

	// We could update this to check if the API returns a result,
	// but currently this function seems unused in the immediate flow or used for verification.
	// For now, keeping the old implementation as it's deprecated but still functional if the site is up.
	// If we want to fully migrate, we should check against the API.
	// However, the task says "Replace Passage retrieval functionality", not necessarily every utility.
	// Let's update it to be safe, or leave it deprecated.
	// Given it makes a network call, better to leave it or update it.
	// The prompt says "Please do not remove the original code, but mark it as 'to be deprecated'".
	// So I will leave the logic as is for the deprecated parts.

	doc := GetPassageHTMLFunc(ref, "NIV")
	ref = GetReference(doc)
	return len(ref) > 0
}
