package scripturebot

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var url string = "http://www.biblegateway.com/passage/?search=%s&version=%sinterface=print"

func GetReference(ref string, env *SessionData) string {
	res, err := http.Get(fmt.Sprintf(url, ref, env.User.Config.Version))
	if err != nil {
		log.Fatalf("Error getting reference: %v", err)
	}

	body, err := ioutil.ReadAll(res.Body)

	log.Printf("Got reference response: %s", string(body))

	return ""
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
