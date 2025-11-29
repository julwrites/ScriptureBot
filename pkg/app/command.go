package app

import (
	"github.com/julwrites/BotPlatform/pkg/def"
	"github.com/julwrites/BotPlatform/pkg/platform"
)

func ProcessCommand(env def.SessionData, bot platform.Platform) def.SessionData {
	switch env.Msg.Command {
	case ADM_CMD_DUMP:
		env = DumpUserList(env)
		break
	case ADM_MIGRATE:
		env = Migrate(env)
		break
	case CMD_VERSION:
		env = SetVersion(env)
		break
	case CMD_TMS:
		env = GetTMSVerse(env)
		break
	case CMD_DEVO:
		env = GetDevo(env, bot)
		break
	case CMD_SUBSCRIBE:
		env = UpdateSubscription(env)
		break
	case CMD_SEARCH:
		env = GetBibleSearch(env)
		break
	case CMD_ASK:
		env = GetBibleAsk(env)
		break
	case CMD_CLOSE:
		env = CloseAction(env)
	default:
		env = GetBiblePassage(env)
	}

	return env
}
