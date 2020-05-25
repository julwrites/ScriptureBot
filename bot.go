// Brief: Bot logic
// Primary responsibility: Top level logic layer for bot

package main

import (
	"fmt"
	"log"

	bmul "github.com/julwrites/BotMultiplexer"
)

func HelpMessage(env *bmul.SessionData) string {
	return fmt.Sprintf("Hello %s! Give me a Bible reference and I'll give you the passage!\nHere are some other things I can do:\n/version - Choose your preferred Bible version", env.User.Firstname)
	//\n/tms - Get a card from the Navigators' Topical Memory System\n/dailydevo - Get reading material right now\n/subscribe - Subscribe to / Unsubscribe from daily reading material\n/search - Search for a passage, lexicon entry, word or phrase\n"
}

func RunCommands(env bmul.SessionData) bmul.SessionData {
	if len(env.User.Action) > 0 {
		log.Printf("Detected user has active action %s", env.User.Action)
		env.Msg.Command = env.User.Action
	}

	switch env.Msg.Command {
	case CMD_VERSION:
		env = SetVersion(env)
		break
	case CMD_TMS:
		fallthrough
	case CMD_DEVO:
		fallthrough
	case CMD_SUBSCRIBE:
		fallthrough
	case CMD_LEXICON:
		fallthrough
	default:
		env = GetBiblePassage(env)
	}

	return env
}

func HandleBotLogic(env bmul.SessionData) bmul.SessionData {
	env = RunCommands(env)

	log.Printf("Commands run, resulting message: %s", env.Res.Message)
	if len(env.Res.Message) == 0 {
		log.Printf("This message was not handled by bot")
		env.Res.Message = HelpMessage(&env)
	}

	return env
}
