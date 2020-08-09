// Brief: API for url querying
// Primary responsibility: API layer between querying for URLs and other functionality

package utils

import (
	"fmt"
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

func QueryBiblePassage(ref string, ver string) *html.Node {
	query := fmt.Sprintf("https://classic.biblegateway.com/passage/?search=%s&version=%s&interface=print", ref, ver)

	return QueryHtml(query)
}

func QueryBibleLexicon(word string, ver string) *html.Node {
	query := fmt.Sprintf("https://www.blueletterbible.org/search/search.cfm?Criteria=%s&t=%s#s=s_lexiconc", word, ver)

	return QueryHtml(query)
}

func QueryMCheyne() *html.Node {
	query := fmt.Sprintf("http://www.edginet.org/mcheyne/rss_feed.php?type=rss_2.0&tz=0&cal=classic&bible=esv")

	return QueryHtml(query)
}
