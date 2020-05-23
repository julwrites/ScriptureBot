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

type NodePredicate func(*html.Node) bool

func FindNode(node *html.Node, pred NodePredicate) *html.Node {
	if pred(node) {
		return node
	}
	for child := node.FirstChild; child != nil; child = child.NextSibling {
		findNode := FindNode(child, pred)
		if findNode != nil {
			return findNode
		}
	}
	return nil
}

func FindAllNodes(node *html.Node, pred NodePredicate) []*html.Node {
	var outNodes []*html.Node
	if pred(node) {
		outNodes = append(outNodes, node)
	}
	for child := node.FirstChild; child != nil; child = child.NextSibling {
		foundNodes := FindAllNodes(child, pred)
		for _, match := range foundNodes {
			outNodes = append(outNodes, match)
		}
	}
	return outNodes
}

func FilterNode(node *html.Node, pred NodePredicate) []*html.Node {
	var outNodes []*html.Node
	if pred(node) {
		outNodes = append(outNodes, node)
	}
	for child := node.FirstChild; child != nil; child = child.NextSibling {
		matchedNodes := FilterNode(child, pred)
		for _, match := range matchedNodes {
			outNodes = append(outNodes, match)
		}
	}
	return outNodes
}

func FilterNodeList(nodes []*html.Node, pred NodePredicate) []*html.Node {
	var outNodes []*html.Node
	for _, node := range nodes {
		if pred(node) {
			outNodes = append(outNodes, node)
		}
	}
	return outNodes
}

// Parses a node and returns the first element with a particular string
func FindByClass(node *html.Node, tag string) (*html.Node, error) {
	foundNode := FindNode(node, func(node *html.Node) bool {
		for _, attr := range node.Attr {
			if attr.Key == "class" && attr.Val == tag {
				return true
			}
		}
		return false
	})
	var err error
	if foundNode == nil {
		err = errors.New(fmt.Sprintf("Missing %s in the node tree", tag))
	}
	return nil, err
}

func FilterByNodeType(node *html.Node, nodeType html.NodeType) []*html.Node {
	return FilterNode(node, func(node *html.Node) bool { return nodeType == node.Type })
}
