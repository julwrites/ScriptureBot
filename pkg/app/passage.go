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
	"github.com/julwrites/ScriptureBot/pkg/utils"
)

// Deprecated: Using new API service
func GetPassageHtml(ref, ver string) *html.Node {
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

func ParseNodesForPassage(node *html.Node) string {
	var parts []string

	for child := node.FirstChild; child != nil; child = child.NextSibling {
		if child.Type == html.TextNode {
			parts = append(parts, child.Data)
		} else if child.Type == html.ElementNode {
			var subParts string
			switch child.Data {
			case "sup":
				isFootnote := func(node *html.Node) bool {
					for _, attr := range node.Attr {
						if attr.Key == "class" && attr.Val == "footnote" {
							return true
						}
					}
					return false
				}
				if isFootnote(child) {
					continue
				}
				childText := ParseNodesForPassage(child)
				if len(childText) > 0 {
					subParts = fmt.Sprintf("<b>%s</b>", childText)
				}
			case "i":
				childText := ParseNodesForPassage(child)
				subParts = fmt.Sprintf("<i>%s</i>", childText)
			case "p", "span", "body", "html":
				subParts = ParseNodesForPassage(child)
			case "br":
				subParts = "\n"
			default:
				subParts = ParseNodesForPassage(child)
			}
			parts = append(parts, subParts)
		}
	}

	text := strings.Join(parts, "")

	if node.Data == "h1" || node.Data == "h2" || node.Data == "h3" || node.Data == "h4" {
		text = fmt.Sprintf("<b>%s</b>", text)
	}
	return text
}

func ParsePassageFromHtml(rawHtml string) string {
	doc, err := html.Parse(strings.NewReader(rawHtml))
	if err != nil {
		log.Printf("Error parsing html: %v", err)
		return rawHtml
	}

	return ParseNodesForPassage(doc)
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

	textBlocks := utils.MapNodeListToString(filtNodes, ParseNodesForPassage)

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
			log.Printf("Error retrieving passage: %v", err)
			// Fallback or error message?
			// For now, let's just log it and potentially return a friendly error message to the user if critical
			// But sticking to existing behavior, maybe just return env unmodified or empty message?
			// The original code returned env if doc == nil.
			// Let's inform the user.
			env.Res.Message = "Sorry, I couldn't retrieve that passage. Please check the reference or try again later."
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

	doc := GetPassageHtml(ref, "NIV")
	ref = GetReference(doc)
	return len(ref) > 0
}
