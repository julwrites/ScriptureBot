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
	log.Printf("Loading session data...")
	env, ok := bmul.TranslateToProps(req, bmul.TYPE_TELEGRAM)
	if !ok {
		log.Printf("This message was not translatable to bot language")
		return
	}

	env.Secrets = *secrets
	log.Printf("\tLoaded secrets...")

	log.Printf("Loading user...")

	env = RegisterUser(env)

	env = HandleBotLogic(env)

	if !bmul.PostFromProps(env) {
		log.Printf("This message was not translatable from bot language")
		return
	}

	PushUser(env) // Any change to the user throughout the commands should be put to database
}
