package app

import (
	"fmt"
	"strings"

	"github.com/julwrites/BotPlatform/pkg/def"
	"github.com/julwrites/ScriptureBot/pkg/secrets"
	"github.com/julwrites/ScriptureBot/pkg/utils"
)

func DumpUserList(env def.SessionData) def.SessionData {
	adminID, _ := secrets.Get("TELEGRAM_ADMIN_ID")
	if env.User.Id == adminID {
		var message = ""

		projectID, _ := secrets.Get("GCLOUD_PROJECT_ID")
		users := utils.GetAllUsers(projectID)
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
	adminID, _ := secrets.Get("TELEGRAM_ADMIN_ID")
	if env.User.Id == adminID {
		projectID, _ := secrets.Get("GCLOUD_PROJECT_ID")
		users := utils.GetAllUsers(projectID)
		for _, user := range users {
			config := utils.DeserializeUserConfig(user.Config)
			if len(config.Subscriptions) > 0 {
				subs := strings.Split(config.Subscriptions, ",")
				migratedSubs := []string{}

				for _, sub := range subs {
					_, err := AcronymizeDevo(sub)
					if err == nil {
						migratedSubs = append(migratedSubs, sub)
					}
				}

				if len(migratedSubs) == 0 {
					config.Subscriptions = ""
				} else {
					config.Subscriptions = strings.Join(migratedSubs, ",")
				}

				user.Config = utils.SerializeUserConfig(config)

				utils.PushUser(user, projectID)
			}
		}
	}

	return env
}
