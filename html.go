// Brief: HTML handling
// Primary responsibility: Parsing of HTML

package scripturebot

import (
	"errors"
	"fmt"

	"golang.org/x/net/html"
)

// Parses a node and returns the first element with a particular string
func Find(node *html.Node, tag string) (*html.Node, error) {
	if node.Type == html.ElementNode {
		for _, attr := range node.Attr {
			if attr.Key == "class" && attr.Val == tag {
				return node, nil
			}
		}
	}
	for child := node.FirstChild; child != nil; child = child.NextSibling {
		findNode, err := Find(child, tag)
		if err == nil {
			return findNode, nil
		}
	}

	if node != nil {
		return node, nil
	}

	return nil, errors.New(fmt.Sprintf("Missing %s in the node tree", tag))
}
