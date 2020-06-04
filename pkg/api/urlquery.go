// Brief: API for url querying
// Primary responsibility: API layer between querying for URLs and other functionality

package api

import (
	"fmt"
	"log"

	"golang.org/x/net/html"

	bmul "github.com/julwrites/BotMultiplexer"
)

func QueryBiblePassage(ref string, env *bmul.SessionData) *html.Node {
	query := fmt.Sprintf("https://www.biblegateway.com/passage/?search=%s&version=%s", ref, DeserializeUserConfig(env.User.Config).Version)

	log.Printf("Query String: %s", query)

	doc := GetHtml(query)

	if doc == nil {
		log.Printf("Error getting html")
		return nil
	}

	return doc
}

func QueryBibleLexicon(word string, env *bmul.SessionData) *html.Node {
	query := fmt.Sprintf("https://www.blueletterbible.org/search/search.cfm?Criteria=%s&t=%s#s=s_lexiconc", word, DeserializeUserConfig(env.User.Config).Version)

	log.Printf("Query String: %s", query)

	doc := GetHtml(query)

	if doc == nil {
		log.Printf("Error getting html")
		return nil
	}

	return doc
}
