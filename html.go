package scripturebot

import (
	"errors"
	"fmt"
	"log"
	"net/http"
)

func Get(url string) *html.Node {
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

func Filter(node *html.Node, tag string) (*html.Node, error) {
	var node *html.Node
	var crawler func(*html.Node)

	crawler = func(node *html.Node) {
		if node.Type == html.ElementNode && node.Attr["class"] == tag {
			body = node
			return
		}
		for child := node.FirstChild; child != nil; child = child.NextSibling {
			crawler(child)
		}
	}

	crawler(doc)

	if node != nil {
		return node, nil
	}

	return nil, errors.New(fmt.Sprintf("Missing %s in the node tree", tag))
}
