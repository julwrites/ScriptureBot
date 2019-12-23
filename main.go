package scripturebot

import (
	"log"
	"net/http"

	botsecrets "github.com/julwrites/BotSecrets"
)

func TelegramHandler(res http.ResponseWriter, req *http.Request, secrets *botsecrets.SecretsData) {
	env := SessionData{}
	log.Printf("Loading session data...")

	env.Type = TYPE_TELEGRAM

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

	if !PostFromProps(&env) {
		log.Printf("This message was not translatable from bot language")
		return
	}
}
