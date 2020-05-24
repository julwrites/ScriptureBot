// Brief: Bot logic
// Primary responsibility: Top level logic layer for bot

package main

import (
	"fmt"
	"log"

	bmul "github.com/julwrites/BotMultiplexer"
)

func HelpMessage(env *bmul.SessionData) string {
	return fmt.Sprintf("Hello %s! Give me a Bible reference and I'll give you the passage!", env.User.Firstname)
	//\nHere are some other things I can do:\n/tms - Get a card from the Navigators' Topical Memory System\n/version - Choose your preferred Bible version\n/dailydevo - Get reading material right now\n/subscribe - Subscribe to / Unsubscribe from daily reading material\n/search - Search for a passage, lexicon entry, word or phrase\n"
}

func RunCommands(env *bmul.SessionData) {
	if len(env.User.Action) > 0 {
		env.Msg.Command = env.User.Action
	}

	switch env.Msg.Command {
	case CMD_VERSION:
		SetVersion(env)
		break
	case CMD_TMS:
		break
	case CMD_DEVO:
		break
	case CMD_SUBSCRIBE:
		break
	case CMD_LEXICON:
		break
	default:
		GetBiblePassage(env)
	}
}

func HandleBotLogic(env *bmul.SessionData) {
	RunCommands(env)

	log.Printf("Commands run, resulting message: %s", env.Res.Message)
	if len(env.Res.Message) == 0 {
		log.Printf("This message was not handled by bot")
		env.Res.Message = HelpMessage(env)
	}
}
