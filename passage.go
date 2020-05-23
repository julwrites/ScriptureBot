// Brief: Scripture API
// Primary responsibility: Query calls for Scripture functionality

package main

import (
	"fmt"
	"log"
	"strings"

	bmul "github.com/julwrites/BotMultiplexer"
	"golang.org/x/net/html"
)

func GetReference(doc *html.Node, env *bmul.SessionData) string {
	refNode, err := FindByClass(doc, "bcv")
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
			parts = append(parts, "*")
			parts = append(parts, ParseNodesForPassage(child))
			parts = append(parts, "*")
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
			parts = append(parts, "^")
			parts = append(parts, ParseNodesForPassage(child))
			parts = append(parts, "^")
			break
		default:
			parts = append(parts, child.Data)
		}
	}
	text = strings.Join(parts, "")
	return text
}

func GetPassage(doc *html.Node, env *bmul.SessionData) string {
	passageNode, startErr := FindByClass(doc, fmt.Sprintf("version-%s result-text-style-normal text-html", GetUserConfig(&env.User).Version))
	if startErr != nil {
		log.Printf("Error parsing for passage: %v", startErr)
		return ""
	}

	filtNodes := FilterChildren(passageNode, func(child *html.Node) bool {
		switch tag := child.Data; tag {
		case "h1":
			fallthrough
		case "h2":
			fallthrough
		case "h3":
			fallthrough
		case "p":
			return true
		}
		return false
	})

	log.Printf("Candidate nodes number %d", len(filtNodes))

	var textBlocks []string
	for _, node := range filtNodes {
		textBlocks = append(textBlocks, MapNodeList(node, ParseNodesForPassage)...)
		textBlocks = append(textBlocks, "\n")
	}

	var passage strings.Builder

	for _, block := range textBlocks {
		passage.WriteString(block)
	}

	log.Printf("%s", passage.String())

	return fmt.Sprintf("I currently can't parse a passage but here's what I got so far: %s", passage.String())
}

func GetBiblePassage(env *bmul.SessionData) {
	if len(env.Msg.Message) > 0 {

		doc := QueryBibleGateway(env.Msg.Message, env)

		ref := GetReference(doc, env)
		log.Printf("Reference retrieved: %s", ref)

		if len(ref) > 0 {
			log.Printf("Getting passage")
			env.Res.Message = GetPassage(doc, env)
		}
	}
}

// func main() {
// 	var env bmul.SessionData
// 	var config UserConfig
// 	config.Version = "NIV"
// 	UpdateUserConfig(&env.User, config)
// 	env.Msg.Message = "gal 1"
// 	GetBiblePassage(&env)
// }
