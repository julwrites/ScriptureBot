// Brief: Scripture API
// Primary responsibility: Query calls for Scripture functionality

package main

import (
	"fmt"
	"log"

	bmul "github.com/julwrites/BotMultiplexer"
	"golang.org/x/net/html"
)

var passageQuery string = "http://www.biblegateway.com/passage/?search=%s&version=%s"

func Query(ref string, env *bmul.SessionData) *html.Node {
	query := fmt.Sprintf(passageQuery, ref, GetUserConfig(&env.User).Version)

	log.Printf("Query String: %s", query)

	doc := GetHtml(query)

	if doc == nil {
		log.Fatalf("Error getting html")
		return nil
	}

	return doc
}

func GetReference(doc *html.Node, env *bmul.SessionData) string {
	foundRef, err := FindByClass(doc, "bcv")
	if err != nil {
		log.Fatalf("Error parsing for reference: %v", err)
	}

	return foundRef.FirstChild.Data
}

func GetPassage(doc *html.Node, env *bmul.SessionData) string {
	passage, err := FindByClass(doc, "passage-text")
	if err != nil {
		log.Fatalf("Error parsing for passage: %v", err)
	}

	return fmt.Sprintf("I currently can't parse a passage but here's what I got so far: %s", passage.Data)
}

func GetBiblePassage(env *bmul.SessionData) bool {
	if len(env.Msg.Message) > 0 {

		doc := Query(env.Msg.Message, env)
		ref := GetReference(doc, env)

		if len(ref) > 0 {
			env.Res.Message = GetPassage(doc, env)

			return true
		}
	}

	return false
}
