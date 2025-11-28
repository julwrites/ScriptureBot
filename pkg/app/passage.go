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

// Direct Scraping, not using Bible AI API for intelligence, but fast for checking some simple things like references
func CheckBibleReference(ref string) bool {
	log.Printf("Checking reference %s", ref)

	doc := GetPassageHTML(ref, "NIV")
	ref = GetReference(doc)
	return len(ref) > 0
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


func ParseNodesForPassage(node *html.Node) string {
	var text string
	var parts []string

	for child := node.FirstChild; child != nil; child = child.NextSibling {
		parts = append(parts, text)

		switch tag := child.Data; tag {
		case "span":
			childText := strings.Trim(ParseNodesForPassage(child), " ")
			if len(childText) > 0 {
				parts = append(parts, childText)
			} else {
				parts = append(parts, child.Data)
			}
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
				break
			}
			childText := strings.Trim(ParseNodesForPassage(child), " ")
			if len(childText) > 0 {
				parts = append(parts, fmt.Sprintf("^%s^", childText))
			}
			break
		case "p":
		case "i":
			parts = append(parts, ParseNodesForPassage(child))
			break
		case "br":
			parts = append(parts, "\n")
			break
		default:
			parts = append(parts, child.Data)
		}
	}

	text = strings.Join(parts, "")

	if node.Data == "h1" || node.Data == "h2" || node.Data == "h3" || node.Data == "h4" {
		text = fmt.Sprintf("*%s*", text)
	}
	return text
}

func GetPassage(ref string, doc *html.Node, version string) string {
	filtNodes := utils.FilterTree(doc, func(child *html.Node) bool {
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

func ParsePassageFromHtml(ref string, rawHtml string, version string) string {
	doc, err := html.Parse(strings.NewReader(rawHtml))

	if err != nil {
		log.Printf("Error parsing html: %v", err)
		return rawHtml
	}

	return strings.TrimSpace(GetPassage(ref, doc, version))
}

func GetBiblePassageFallback(env def.SessionData) def.SessionData {
	config := utils.DeserializeUserConfig(env.User.Config)

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

	return env
}

func GetBiblePassage(env def.SessionData) def.SessionData {
	if len(env.Msg.Message) > 0 {
		config := utils.DeserializeUserConfig(env.User.Config)

		// To be replaced with a simpler algorithm
		doc := GetPassageHTML(env.Msg.Message, config.Version)
		ref := GetReference(doc)

		// If indeed a reference, attempt to query
		if len(ref) > 0 {

			// Attempt to retrieve from API
			req := QueryRequest{
				Query: QueryObject{
					Verses: []string{ref},
				},
				Context: QueryContext{
					User: UserContext{
						Version: config.Version,
					},
				},
			}

			var resp VerseResponse
			err := SubmitQuery(req, &resp, env.Secrets.PROJECT_ID)

			// Fallback to direct passage retrieval logic
			if err != nil {
				log.Printf("Error retrieving passage from API: %v. Falling back to deprecated method.", err)

				return GetBiblePassageFallback(env)
			} 

			if len(resp.Verse) > 0 {
				env.Res.Message = ParsePassageFromHtml(ref, resp.Verse, config.Version)
			}
		}

		env.Res.Message = "No verses found."
	}

	return env
}
