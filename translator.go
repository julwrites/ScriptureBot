package scripturebot

import (
	"io/ioutil"
	"log"
	"net/http"
)

// Translator component to handle translation of a HTTP payload into a
// consistent format, and to translate that format back into a HTTP payload
// for posting

func TranslateToProps(req *http.Request, env *SessionData) bool {
	reqBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Fatalf("Error occurred reading http request: %s", err)
		return false
	}
	log.Printf("Request body: %s", reqBody)

	switch env.Type {
	case TYPE_TELEGRAM:
		log.Printf("Parsing Telegram message")

		return true
	}

	return false
}

func TranslateToHttp(props *SessionData) bool {
	return false
}
