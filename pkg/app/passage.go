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

func GetPassageHtml(ref, ver string) *html.Node {
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

		doc := GetPassageHtml(env.Msg.Message, utils.DeserializeUserConfig(env.User.Config).Version)
		if doc == nil {
			return env
		}

		ref := GetReference(doc)
		log.Printf("Reference retrieved: %s", ref)

		if len(ref) > 0 {
			log.Printf("Getting passage")
			env.Res.Message = GetPassage(ref, doc, utils.DeserializeUserConfig(env.User.Config).Version)
			log.Printf("Passage retrieved length: %d", len(env.Res.Message))
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
