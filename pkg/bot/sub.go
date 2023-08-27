package bot

import (
	"log"
	"strings"

	"github.com/julwrites/BotPlatform/pkg/def"
	"github.com/julwrites/BotPlatform/pkg/platform"
	"github.com/julwrites/BotPlatform/pkg/secrets"
	"github.com/julwrites/ScriptureBot/pkg/app"
	"github.com/julwrites/ScriptureBot/pkg/utils"
)

func HandleSubscriptionLogic(env def.SessionData) def.SessionData {
	// Check all users
	users := utils.GetAllUsers(env.Secrets.PROJECT_ID)
	for _, user := range users {
		env.User = user
		env.User.Action = app.CMD_DEVO

		config := utils.DeserializeUserConfig(user.Config)

		subscriptions := strings.Split(config.Subscriptions, ",")

		for _, devo := range subscriptions {
			env.Res.Affordances.Remove = true

			// Retrieve devotional
			env.Res = app.GetDevotionalData(env, devo)

			env.User.Action = ""
		}

		if !platform.PostFromProps(env) {
			log.Printf("This message was not translatable from bot language")
			continue
		}
	}

	return env
}

func SubscriptionHandler(secrets *secrets.SecretsData) {
	env := def.SessionData{}

	env.Secrets = *secrets
	log.Printf("Loaded secrets...")

	env.ResourcePath = "/go/bin/"

	env = HandleSubscriptionLogic(env)
	log.Printf("Handled bot logic...")

	if !platform.PostFromProps(env) {
		log.Printf("This message was not translatable from bot language")
		return
	}
}
