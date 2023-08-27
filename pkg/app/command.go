package app

import "github.com/julwrites/BotPlatform/pkg/def"

func ProcessCommand(env def.SessionData) def.SessionData {
	switch env.Msg.Command {
	case ADM_CMD_DUMP:
		env = DumpUserList(env)
		break
	case CMD_VERSION:
		env = SetVersion(env)
		break
	case CMD_TMS:
		env = GetTMSVerse(env)
		break
	case CMD_DEVO:
		env = GetDevo(env)
		break
	case CMD_SUBSCRIBE:
		env = UpdateSubscription(env)
		break
	case CMD_LEXICON:
		env = GetBibleWord(env)
		break
	case CMD_CLOSE:
		env = CloseAction(env)
	default:
		env = GetBiblePassage(env)
	}

	return env
}
