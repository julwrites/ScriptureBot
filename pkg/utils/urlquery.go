// Brief: API for url querying
// Primary responsibility: API layer between querying for URLs and other functionality

package utils

import (
	"fmt"
	"log"

	"golang.org/x/net/html"
)

func QueryBiblePassage(ref string, ver string) *html.Node {
	query := fmt.Sprintf("https://www.biblegateway.com/passage/?search=%s&version=%s", ref, ver)

	log.Printf("Query String: %s", query)

	doc := GetHtml(query)

	if doc == nil {
		log.Printf("Error getting html")
		return nil
	}

	return doc
}

func QueryBibleLexicon(word string, ver string) *html.Node {
	query := fmt.Sprintf("https://www.blueletterbible.org/search/search.cfm?Criteria=%s&t=%s#s=s_lexiconc", word, ver)

	log.Printf("Query String: %s", query)

	doc := GetHtml(query)

	if doc == nil {
		log.Printf("Error getting html")
		return nil
	}

	return doc
}
