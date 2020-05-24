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

	HandleBotLogic(&env)

	if !bmul.PostFromProps(&env) {
		log.Printf("This message was not translatable from bot language")
		return
	}
}
