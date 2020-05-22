// Brief: Scripture API
// Primary responsibility: Query calls for Scripture functionality

package main

import (
	"fmt"
	"log"
	"net/url"

	bmul "github.com/julwrites/BotMultiplexer"
)

var passageQuery string = "www.biblegateway.com/passage/?search=%s&version=%s"

func GetReference(ref string, env *bmul.SessionData) string {
	query := fmt.Sprintf(passageQuery, ref, GetUserConfig(&env.User).Version)

	log.Printf("Query String: %s", query)

	doc := GetHtml(query)

	if doc == nil {
		log.Fatalf("Error getting reference")
		return ""
	}

	foundRef, err := Find(doc, "passage-display-bcv")
	if err != nil {
		log.Fatalf("Error in Finding of Reference: %v", err)
	}

	return foundRef.Data
}

func GetPassage(ref string, env *bmul.SessionData) string {
	query := fmt.Sprintf(passageQuery, ref, GetUserConfig(&env.User).Version)
	query = url.QueryEscape(query)

	doc := GetHtml(query)

	if doc == nil {
		log.Fatalf("Error getting reference")
		return ""
	}

	foundRef, err := Find(doc, "passage-display-bcv")
	if err != nil {
		log.Fatalf("Error in Finding of Reference: %v", err)
	}

	return foundRef.Data
}

func GetBiblePassage(env *bmul.SessionData) bool {
	if len(env.Msg.Message) > 0 {

		ref := GetReference(env.Msg.Message, env)

		if len(ref) > 0 {
			env.Res.Message = GetPassage(ref, env)

			return true
		}
	}

	return false
}
