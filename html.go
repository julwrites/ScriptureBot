// Brief: HTML handling
// Primary responsibility: Parsing of HTML

package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"golang.org/x/net/html"
)

func GetHtml(url string) *html.Node {
	res, getErr := http.Get(url)
	if getErr != nil {
		log.Fatalf("Error in GET call: %v", getErr)
		return nil
	}

	doc, parseErr := html.Parse(res.Body)
	if parseErr != nil {
		log.Fatalf("Error parsing html: %v", parseErr)
	}

	return doc
}

// Parses a node and returns the first element with a particular string
func FindByClass(node *html.Node, tag string) (*html.Node, error) {
	for _, attr := range node.Attr {
		if attr.Key == "class" && attr.Val == tag {
			return node, nil
		}
	}
	for child := node.FirstChild; child != nil; child = child.NextSibling {
		findNode, err := FindByClass(child, tag)
		if err == nil {
			return findNode, nil
		}
	}

	return nil, errors.New(fmt.Sprintf("Missing %s in the node tree", tag))
}

func FindByNodeType(node *html.Node, nodeType html.NodeType) []*html.Node {
	var nodes []*html.Node
	if node.Type == nodeType {
		nodes = append(nodes, node)
	}
	for child := node.FirstChild; child != nil; child = child.NextSibling {
		matchedNodes := FindByNodeType(child, nodeType)
		for _, match := range matchedNodes {
			nodes = append(nodes, match)
		}
	}

	return nodes
}
