package app

import "github.com/julwrites/BotMultiplexer/pkg/def"

func ProcessCommand(env def.SessionData) def.SessionData {
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
