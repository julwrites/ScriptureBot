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
	log.Printf("Retrieved %d users", len(users))
	for _, user := range users {
		env.User = user
		env.User.Action = app.CMD_DEVO

		config := utils.DeserializeUserConfig(user.Config)

		if len(config.Subscriptions) > 0 {
			subscriptions := strings.Split(config.Subscriptions, ",")

			log.Printf("Found subscriptions for %s: %s", user.Firstname+" "+user.Lastname, subscriptions)

			// Clear existing affordances and send a preliminary message
			env.Res.Message = "Here are today's devotions!"
			env.Res.Affordances.Remove = true

			platform.PostFromProps(env)

			// Send the devotional
			for _, devo := range subscriptions {
				// Retrieve devotional
				log.Printf("Getting data for (%s)", devo)
				env.Res = app.GetDevotionalData(env, devo)

				platform.PostFromProps(env)
			}
		}
	}

	return env
}

func SubscriptionHandler(secrets *secrets.SecretsData) {
	env := def.SessionData{}

	env.Secrets = *secrets
	log.Printf("Loaded secrets...")

	env.ResourcePath = "/go/bin/"

	// TODO: Iterate through types
	env.Type = def.TYPE_TELEGRAM

	env = HandleSubscriptionLogic(env)
	log.Printf("Handled bot logic...")

	if !platform.PostFromProps(env) {
		log.Printf("This message was not translatable from bot language")
		return
	}
}
