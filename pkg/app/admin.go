package app

import (
	"github.com/julwrites/BotPlatform/pkg/def"
)

func DumpUserList(env def.SessionData) def.SessionData {
	if env.User.Id == env.Secrets.ADMIN_ID {
		var message = ""
		// Retrieve the whole database
		// Format the message
		env.Res.Message = message
	}

	return env
}
