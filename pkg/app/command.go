package app

import (
	bmul "github.com/julwrites/BotMultiplexer"
)

func ProcessCommand(env bmul.SessionData) bmul.SessionData {
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
