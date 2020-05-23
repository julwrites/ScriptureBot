// Brief: Bot logic
// Primary responsibility: Top level logic layer for bot

package main

import (
	"log"

	bmul "github.com/julwrites/BotMultiplexer"
)

func HelpMessage(env *bmul.SessionData) {
	env.Res.Message = "Hi, this message comes from the bot"
}

func RunCommands(env *bmul.SessionData) {
	switch env.Msg.Command {
	default:
		GetBiblePassage(env)
	}
}

func HandleBotLogic(env *bmul.SessionData) {
	RunCommands(env)

	if len(env.Res.Message) == 0 {
		log.Printf("This message was not handled by bot")
		HelpMessage(env)
	}
}
