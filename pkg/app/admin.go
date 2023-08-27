package app

import (
	"fmt"

	"github.com/julwrites/BotPlatform/pkg/def"
	"github.com/julwrites/ScriptureBot/pkg/utils"
)

func DumpUserList(env def.SessionData) def.SessionData {
	if env.User.Id == env.Secrets.ADMIN_ID {
		var message = ""

		users := utils.GetAllUsers(env.Secrets.PROJECT_ID)
		message += fmt.Sprintf("%d Users: \n", len(users))
		for _, user := range users {
			message += user.Firstname + " " + user.Lastname + " - @" + user.Username + "\n"
		}

		// Retrieve the whole database
		// Format the message
		env.Res.Message = message
	}

	return env
}
