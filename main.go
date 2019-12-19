package ScriptureBot

import (
	"log"
	"net/http"
)

// Bot methods
func HandleBotLogic(props *SessionData) bool {
	return false
}

func telegramHandler(res http.ResponseWriter, req *http.Request, secrets *SecretsData) {
	env := SessionData{}

	log.Printf("Loading session data...")

	env.Secrets = *secrets

	log.Printf("\tLoaded secrets...")

	if !TranslateToProps(req, &env) {
		log.Printf("This message was not translatable to bot language")
		return
	}

	log.Printf("\tLoaded message...")

	if !HandleBotLogic(&env) {
		log.Printf("This message was not handled by bot")
		return
	}

	if !TranslateToHttp(&env) {
		log.Printf("This message was not translatable from bot language")
		return
	}
}
