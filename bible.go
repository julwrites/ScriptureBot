// Brief: Scripture API
// Primary responsibility: Query calls for Scripture functionality

package scripturebot

import (
	"fmt"
	"log"
	"net/url"

	botmultiplexer "github.com/julwrites/BotMultiplexer"
)

var formatQuery string = "http://www.biblegateway.com/passage/?search=%s&version=%sinterface=print"

func GetReference(ref string, env *botmultiplexer.SessionData) string {
	query := fmt.Sprintf(formatQuery, ref, env.User.Config.Version)
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

func GetPassage(ref string, env *botmultiplexer.SessionData) string {
	return ""
}

func GetBiblePassage(env *botmultiplexer.SessionData) bool {
	if len(env.Msg.Message) > 0 {

		ref := GetReference(env.Msg.Message, env)

		if len(ref) > 0 {
			env.Res.Message = GetPassage(ref, env)

			return true
		}
	}

	return false
}
