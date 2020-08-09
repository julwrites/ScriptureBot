// Brief: Bible Passage handling
// Primary responsibility: Parsing of Bible Passages from HTML

package app

import (
	"fmt"
	"log"
	"strings"

	"golang.org/x/net/html"

	"github.com/julwrites/BotPlatform/pkg/def"
	"github.com/julwrites/ScriptureBot/pkg/utils"
)

func GetPassageHtml(ref string, ver string) *html.Node {
	query := fmt.Sprintf("https://classic.biblegateway.com/passage/?search=%s&version=%s&interface=print", ref, ver)

	return utils.QueryHtml(query)
}

func GetReference(doc *html.Node) string {
	refNode, err := utils.FindByClass(doc, "bcv")
	if err != nil {
		log.Printf("Error parsing for reference: %v", err)
		return ""
	}

	return refNode.FirstChild.Data
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

func GetPassage(doc *html.Node, version string) string {
	passageNode, startErr := utils.FindByClass(doc, fmt.Sprintf("version-%s result-text-style-normal text-html", version))
	if startErr != nil {
		log.Printf("Error parsing for passage: %v", startErr)
		return ""
	}

	filtNodes := utils.FilterChildren(passageNode, func(child *html.Node) bool {
		switch tag := child.Data; tag {
		case "h1":
			fallthrough
		case "h2":
			fallthrough
		case "h3":
			fallthrough
		case "h4":
			fallthrough
		case "p":
			return true
		case "div":
			for _, attr := range child.Attr {
				if attr.Key == "class" && strings.Contains(attr.Val, "poetry") {
					return true
				}
			}
		}
		return false
	})

	textBlocks := utils.MapNodeListToString(filtNodes, ParseNodesForPassage)

	var passage strings.Builder

	for _, block := range textBlocks {
		passage.WriteString(block)
		passage.WriteString("\n")
	}

	return passage.String()
}

func GetBiblePassage(env def.SessionData) def.SessionData {
	if len(env.Msg.Message) > 0 {

		doc := GetPassageHtml(env.Msg.Message, utils.DeserializeUserConfig(env.User.Config).Version)

		ref := GetReference(doc)
		log.Printf("Reference retrieved: %s", ref)

		if len(ref) > 0 {
			log.Printf("Getting passage")
			env.Res.Message = GetPassage(doc, utils.DeserializeUserConfig(env.User.Config).Version)
			log.Printf("Passage retrieved %s", env.Res.Message)
		}
	}

	return env
}

func CheckBibleReference(ref string) bool {
	log.Printf("Checking reference %s", ref)

	doc := GetPassageHtml(ref, "NIV")

	ref = GetReference(doc)

	return len(ref) > 0
}
