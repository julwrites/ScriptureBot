package scripturebot

import (
	"log"
	"net/url"
)

var formatQuery string = "http://www.biblegateway.com/passage/?search=%s&version=%sinterface=print"

func GetReference(ref string, env *SessionData) string {
	query := formatQuery.Sprintf(url, ref, env.User.Config.Version)
	query = url.QueryEscape(query)

	doc := Get(query)

	if doc == nil {
		log.Fatalf("Error getting reference")
		return ""
	}

	ref := Filter(doc, "passage-display-bcv")

	return ref.Data
}

func GetPassage(ref string, env *SessionData) string {
	return ""
}

func GetBiblePassage(env *SessionData) bool {
	if len(env.Msg.Message) > 0 {

		ref := GetReference(env.Msg.Message, env)

		if len(ref) > 0 {
			env.Res.Message = GetPassage(ref, env)

			return true
		}
	}

	return false
}
