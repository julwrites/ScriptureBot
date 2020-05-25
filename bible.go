package main

import (
	"fmt"
	"log"

	bmul "github.com/julwrites/BotMultiplexer"
	"golang.org/x/net/html"
)

func QueryBibleGateway(ref string, env *bmul.SessionData) *html.Node {
	query := fmt.Sprintf("https://www.biblegateway.com/passage/?search=%s&version=%s", ref, DeserializeUserConfig(env.User.Config).Version)

	log.Printf("Query String: %s", query)

	doc := GetHtml(query)

	if doc == nil {
		log.Printf("Error getting html")
		return nil
	}

	return doc
}

func QueryBlueLetterBible(word string, env *bmul.SessionData) *html.Node {
	query := fmt.Sprintf("https://www.blueletterbible.org/search/search.cfm?Criteria=%s&t=%s#s=s_lexiconc", word, DeserializeUserConfig(env.User.Config).Version)

	log.Printf("Query String: %s", query)

	doc := GetHtml(query)

	if doc == nil {
		log.Printf("Error getting html")
		return nil
	}

	return doc
}
