package app

import "github.com/julwrites/BotPlatform/pkg/def"

func ProcessCommand(env def.SessionData) def.SessionData {
	switch env.Msg.Command {
	case CMD_VERSION:
		env = SetVersion(env)
		break
	case CMD_TMS:
		env = GetTMSVerse(env)
		break
	case CMD_DEVO:
		break
	case CMD_SUBSCRIBE:
		break
	case CMD_LEXICON:
		env = GetBibleWord(env)
		break
	default:
		env = GetBiblePassage(env)
	}

	return env
}
