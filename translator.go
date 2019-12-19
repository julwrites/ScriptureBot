package scripturebot

import (
	"log"
	"net/http"
)

// Translator component to handle translation of a HTTP payload into a
// consistent format, and to translate that format back into a HTTP payload
// for posting

func TranslateToProps(req *http.Request, env *SessionData) bool {
	if req.URL.Path == ("/" + env.Secrets.TELEGRAM_ID) {
		log.Printf("Telegram message")
		return true
	}

	return false
}

func TranslateToHttp(props *SessionData) bool {
	return false
}
