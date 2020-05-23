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

func GetPassage(doc *html.Node, env *bmul.SessionData) string {
	passageNode, startErr := FindByClass(doc, "passage-text")
	if startErr != nil {
		log.Printf("Error parsing for passage: %v", startErr)
		return ""
	}

	var candNodes []*html.Node
	for child := passageNode.FirstChild; child != nil; child = child.NextSibling {
		candNodes = append(candNodes, child)
	}
	log.Printf("Candidate notes number %d", len(candNodes))
	filtNodes := FilterNodeList(candNodes, func(node *html.Node) bool {
		for _, attr := range node.Attr {
			if attr.Key == "class" && attr.Val == "footnotes" {
				return false
			}
		}
		return true
	})

	var textNodes []*html.Node
	for _, node := range filtNodes {
		for _, textNode := range FilterNode(node, func(node *html.Node) bool {
			return node.Type == html.TextNode
		}) {
			textNodes = append(textNodes, textNode)
		}
	}

	var passage strings.Builder

	for _, node := range textNodes {
		passage.WriteString(node.Data)
	}

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
// 	env.Msg.Message = "rev 1"
// 	GetBiblePassage(&env)
// }
