package app

import (
	"github.com/julwrites/BotPlatform/pkg/def"
)

func CloseAction(env def.SessionData) def.SessionData {
	env.Res.Affordances.Remove = true
	env.User.Action = ""

	return env
}
