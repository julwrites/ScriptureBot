// Brief: Handler for Telegram-specific messages
// Primary responsibility: Receive and handle Telegram messages from Request to Response

package main

import (
	"log"
	"net/http"

	botmultiplexer "github.com/julwrites/BotMultiplexer"
	botsecrets "github.com/julwrites/BotSecrets"
)

func TelegramHandler(res http.ResponseWriter, req *http.Request, secrets *botsecrets.SecretsData) {
	env := botmultiplexer.SessionData{}
	log.Printf("Loading session data...")

	env.Type = botmultiplexer.TYPE_TELEGRAM

	env.Secrets = *secrets
	log.Printf("\tLoaded secrets...")

	if !botmultiplexer.TranslateToProps(req, &env) {
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

	if !botmultiplexer.PostFromProps(&env) {
		log.Printf("This message was not translatable from bot language")
		return
	}
}
