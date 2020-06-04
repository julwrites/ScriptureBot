package app

import (
	bmul "github.com/julwrites/BotMultiplexer"
	"github.com/julwrites/ScriptureBot/pkg/api"
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
		env = api.GetBiblePassage(env)
	}

	return env
}
