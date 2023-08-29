package app

import (
	"fmt"
	"strings"

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

func Migrate(env def.SessionData) def.SessionData {
	if env.User.Id == env.Secrets.ADMIN_ID {
		users := utils.GetAllUsers(env.Secrets.PROJECT_ID)
		for _, user := range users {
			config := utils.DeserializeUserConfig(user.Config)
			if len(config.Subscriptions) > 0 {
				subs := strings.Split(config.Subscriptions, ",")
				migratedSubs := []string{}

				for _, sub := range subs {
					if sub == "NTBRP" {
						migratedSubs = append(migratedSubs, "N5XBRP")
					} else {
						migratedSubs = append(migratedSubs, sub)
					}
				}

				config.Subscriptions = strings.Join(migratedSubs, ",")

				user.Config = utils.SerializeUserConfig(config)

				utils.PushUser(user, env.Secrets.PROJECT_ID)
			}
		}
	}

	return env
}
