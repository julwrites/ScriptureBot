// Brief: Bible Passage handling
// Primary responsibility: Parsing of Bible Passages from HTML

package app

import (
	"fmt"
	"log"
	"net/url"
	"regexp"
	"strings"
	stdhtml "html"

	"golang.org/x/net/html"

	"github.com/julwrites/BotPlatform/pkg/def"
	"github.com/julwrites/BotPlatform/pkg/platform"
	"github.com/julwrites/ScriptureBot/pkg/utils"
)

// CheckBibleReference validates if the string is a valid Bible reference using local logic.
func CheckBibleReference(ref string) bool {
	log.Printf("Checking reference %s", ref)
	_, ok := ParseBibleReference(ref)
	return ok
}

var GetPassageHTML = func(ref, ver string) *html.Node {
	ref = url.QueryEscape(ref)
	ver = url.QueryEscape(ver)
	query := fmt.Sprintf("https://classic.biblegateway.com/passage/?search=%s&version=%s&interface=print", ref, ver)

	return utils.QueryHtml(query)
}

func GetReference(doc *html.Node) string {
	refNode, err := utils.FindByClass(doc, "bcv")
	if err != nil {
		log.Printf("Error parsing for reference: %v", err)
		return ""
	}

	return utils.GetTextNode(refNode).Data
}


func isNextSiblingBr(node *html.Node) bool {
	for next := node.NextSibling; next != nil; next = next.NextSibling {
		if next.Type == html.TextNode {
			if len(strings.TrimSpace(next.Data)) == 0 {
				continue
			}
			return false
		}
		if next.Type == html.ElementNode && next.Data == "br" {
			return true
		}
		return false
	}
	return false
}

func hasNextSignificantSibling(node *html.Node) bool {
	for next := node.NextSibling; next != nil; next = next.NextSibling {
		if next.Type == html.TextNode {
			if len(strings.TrimSpace(next.Data)) == 0 {
				continue
			}
			return true
		}
		return true // Any element
	}
	return false
}

func ParseNodesForPassage(node *html.Node) string {
	var parts []string

	for child := node.FirstChild; child != nil; child = child.NextSibling {
		// Filter out footnotes sections/cross-refs if they appear as divs
		if child.Type == html.ElementNode {
			for _, attr := range child.Attr {
				if attr.Key == "class" {
					if strings.Contains(attr.Val, "footnotes") || strings.Contains(attr.Val, "cross-refs") {
						continue
					}
				}
			}
		}

		switch tag := child.Data; tag {
		case "span":
			// Keep existing logic for span (likely poetry lines in legacy/scraped HTML)
			childText := ParseNodesForPassage(child)
			parts = append(parts, childText)
			if len(strings.TrimSpace(childText)) > 0 && !isNextSiblingBr(child) {
				parts = append(parts, "\n")
			}
		case "sup":
			// Handle superscripts (verse numbers/footnotes)
			isFootnote := func(node *html.Node) bool {
				for _, attr := range node.Attr {
					if attr.Key == "class" && attr.Val == "footnote" {
						return true
					}
				}
				return false
			}
			if isFootnote(child) {
				break
			}
			childText := ParseNodesForPassage(child)
			// Use TelegramSuperscript for unicode conversion
			if len(childText) > 0 {
				parts = append(parts, platform.TelegramSuperscript(childText))
			}
			break
		case "p":
			parts = append(parts, ParseNodesForPassage(child))
			parts = append(parts, "\n\n")
		case "b", "strong":
			parts = append(parts, fmt.Sprintf("<b>%s</b>", ParseNodesForPassage(child)))
		case "i", "em":
			parts = append(parts, fmt.Sprintf("<i>%s</i>", ParseNodesForPassage(child)))
		case "h1", "h2", "h3", "h4", "h5", "h6":
			// Ignore "Footnotes" or "Cross references" headers
			headerText := ParseNodesForPassage(child)
			if headerText == "Footnotes" || headerText == "Cross references" {
				continue
			}
			parts = append(parts, fmt.Sprintf("\n\n<b>%s</b>\n", strings.TrimSpace(headerText)))
		case "ul", "ol":
			parts = append(parts, ParseNodesForPassage(child))
		case "li":
			parts = append(parts, fmt.Sprintf("• %s\n", ParseNodesForPassage(child)))
		case "br":
			parts = append(parts, "\n")
		case "div":
			parts = append(parts, ParseNodesForPassage(child))
		default:
			if child.Type == html.TextNode {
				parts = append(parts, stdhtml.EscapeString(child.Data))
			} else if child.Type == html.ElementNode {
				// Recurse for unknown elements to preserve content
				parts = append(parts, ParseNodesForPassage(child))
			}
		}
	}

	return strings.Join(parts, "")
}

// Collapse multiple newlines (potentially with spaces in between) to max 2 newlines
// \n\s*\n\s*\n+ -> \n\n
var newlineRegex = regexp.MustCompile(`\n\s*\n[\s\n]*`)

func CleanPassageText(text string) string {
	text = newlineRegex.ReplaceAllString(text, "\n\n")
	return strings.TrimSpace(text)
}

func GetPassage(ref string, doc *html.Node, version string) string {
	// Replaced FilterTree with direct parsing of the root node
	// This allows handling arbitrary structure (divs, lists) returned by the API

	text := ParseNodesForPassage(doc)
	text = CleanPassageText(text)

	var passage strings.Builder

	if len(ref) > 0 {
		// Use HTML formatting for reference
		refString := fmt.Sprintf("<i>%s</i> (%s)", ref, version)
		passage.WriteString(refString)
	}

	passage.WriteString("\n")
	passage.WriteString(text)

	return passage.String()
}

func ParsePassageFromHtml(ref string, rawHtml string, version string) string {
	doc, err := html.Parse(strings.NewReader(rawHtml))

	if err != nil {
		log.Printf("Error parsing html: %v", err)
		return rawHtml
	}

	// html.Parse returns a doc with html->body structure.
	// GetPassage -> ParseNodesForPassage will traverse it.
	// We might want to find 'body' to avoid processing 'head'?
	// ParseNodesForPassage iterates children. doc->html->body.
	// We can let it recurse.
	return strings.TrimSpace(GetPassage(ref, doc, version))
}

func GetBiblePassageFallback(env def.SessionData) def.SessionData {
	config := utils.DeserializeUserConfig(utils.GetUserConfig(env))

	doc := GetPassageHTML(env.Msg.Message, config.Version)
	ref := GetReference(doc)

	if doc == nil {
		env.Res.Message = "Sorry, I couldn't retrieve that passage. Please check the reference or try again later."
		return env
	}

	// Scrape the passage text
	passageNode, startErr := utils.FindByClass(doc, "passage-text")
	if startErr != nil {
		log.Printf("Error parsing for passage: %v", startErr)
		return env
	}

	// Attempt to get the passage
	env.Res.Message = GetPassage(ref, passageNode, config.Version)
	env.Res.ParseMode = def.TELEGRAM_PARSE_MODE_HTML

	return env
}

func GetBiblePassage(env def.SessionData) def.SessionData {
	if len(env.Msg.Message) > 0 {
		// Identify and normalize bible reference
		ref, ok := ParseBibleReference(env.Msg.Message)

		if ok {
			env.Msg.Message = ref
		}

		config := utils.DeserializeUserConfig(utils.GetUserConfig(env))

		// If indeed a reference, attempt to query
		if len(ref) > 0 {
			log.Printf("%s", ref);

			// Attempt to retrieve from API
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
			err := SubmitQuery(req, &resp)

			// Fallback to direct passage retrieval logic
			if err != nil {
				log.Printf("Error retrieving passage from API: %v. Falling back to deprecated method.", err)

				return GetBiblePassageFallback(env)
			} 

			if len(resp.Verse) > 0 {
				env.Res.Message = ParsePassageFromHtml(env.Msg.Message, resp.Verse, config.Version)
				env.Res.ParseMode = def.TELEGRAM_PARSE_MODE_HTML
				return env
			}
		}
	}

	return env
}
