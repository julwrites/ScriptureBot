// Brief: Handler for Telegram-specific messages
// Primary responsibility: Receive and handle Telegram messages from Request to Response

package main

import (
	"log"
	"net/http"

	bmul "github.com/julwrites/BotMultiplexer"
	botsecrets "github.com/julwrites/BotSecrets"
)

func TelegramHandler(res http.ResponseWriter, req *http.Request, secrets *botsecrets.SecretsData) {
	env := bmul.SessionData{}
	log.Printf("Loading session data...")

	env.Type = bmul.TYPE_TELEGRAM

	env.Secrets = *secrets
	log.Printf("\tLoaded secrets...")

	if !bmul.TranslateToProps(req, &env) {
		log.Printf("This message was not translatable to bot language")
		return
	}

	log.Printf("Loading user...")

	CompareAndUpdateUser(&env)

	log.Printf("\tLoaded message...")

	if !HandleBotLogic(&env) {
		log.Printf("This message was not handled by bot")
		return
	}

	if !bmul.PostFromProps(&env) {
		log.Printf("This message was not translatable from bot language")
		return
	}
}
