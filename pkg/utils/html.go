// Brief: HTML handling
// Primary responsibility: API for HTML parsing

package utils

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
		log.Printf("Error in GET call: %v", getErr)
		return nil
	}

	doc, parseErr := html.Parse(res.Body)
	if parseErr != nil {
		log.Printf("Error parsing html: %v", parseErr)
		return nil
	}

	return doc
}

func GetTextNode(node *html.Node) *html.Node {
	if node != nil {
		if node.Type == html.TextNode {
			return node
		}
		return GetTextNode(node.FirstChild)
	}
	return nil
}

// Find & Filter functions

type NodePredicate func(*html.Node) bool

func FindNode(node *html.Node, pred NodePredicate) *html.Node {
	if pred(node) {
		log.Printf("Found the node %v", node)
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

func FilterTree(node *html.Node, pred NodePredicate) []*html.Node {
	var outNodes []*html.Node
	if pred(node) {
		outNodes = append(outNodes, node)
	}
	for child := node.FirstChild; child != nil; child = child.NextSibling {
		outNodes = append(outNodes, FilterTree(child, pred)...)
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

func FilterChildren(node *html.Node, pred NodePredicate) []*html.Node {
	var outNodes []*html.Node
	for child := node.FirstChild; child != nil; child = child.NextSibling {
		if pred(child) {
			outNodes = append(outNodes, child)
		}
	}
	return outNodes
}

// Transform functions

type NodeTransform func(*html.Node) string

func MapTreeToString(node *html.Node, tran NodeTransform) []string {
	var output []string
	output = append(output, tran(node))
	for child := node.FirstChild; child != nil; child = child.NextSibling {
		output = append(output, MapTreeToString(child, tran)...)
	}
	return output
}

func MapNodeListToString(nodes []*html.Node, tran NodeTransform) []string {
	var output []string
	for _, node := range nodes {
		output = append(output, tran(node))
	}
	return output
}

// Convenience functions

func FindByClass(node *html.Node, tag string) (*html.Node, error) {
	foundNode := FindNode(node, func(node *html.Node) bool {
		if node != nil {
			for _, attr := range node.Attr {
				if attr.Key == "class" && attr.Val == tag {
					return true
				}
			}
		}
		return false
	})

	var err error
	if foundNode == nil {
		err = errors.New(fmt.Sprintf("Missing %s in the node tree", tag))
	}
	return foundNode, err
}

func FilterByNodeType(node *html.Node, nodeType html.NodeType) []*html.Node {
	return FilterTree(node, func(node *html.Node) bool { return nodeType == node.Type })
}
