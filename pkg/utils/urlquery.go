// Brief: API for url querying
// Primary responsibility: API layer between querying for URLs and other functionality

package utils

import (
	"log"

	"golang.org/x/net/html"
)

func QueryHtml(query string) *html.Node {
	log.Printf("Query String: %s", query)

	doc := GetHtml(query)

	if doc == nil {
		log.Printf("Error getting html")
		return nil
	}

	return doc
}
