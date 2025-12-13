package bot

import (
	"log"
	"strings"

	"github.com/julwrites/BotPlatform/pkg/def"
	"github.com/julwrites/BotPlatform/pkg/platform"
	"github.com/julwrites/ScriptureBot/pkg/app"
	"github.com/julwrites/ScriptureBot/pkg/secrets"
	"github.com/julwrites/ScriptureBot/pkg/utils"
)

func HandleSubscriptionLogic(env def.SessionData, bot platform.Platform) def.SessionData {
	user := utils.GetUserFromSession(env)
	config := utils.DeserializeUserConfig(user.Config)

	if len(config.Subscriptions) > 0 {
		subscriptions := strings.Split(config.Subscriptions, ",")

		log.Printf("Found subscriptions for %s: %s", user.Firstname+" "+user.Lastname, subscriptions)

		// Clear existing affordances and send a preliminary message
		env.Res.Message = "Here are today's devotions!"
		env.Res.Affordances.Remove = true

		bot.Post(env)

		// Send the devotional
		for _, devo := range subscriptions {
			// Retrieve devotional
			log.Printf("Getting data for <%s>", devo)
			env.Res = app.GetDevotionalData(env, devo)

			bot.Post(env)
		}
	}

	return env
}

func HandleSubscriptionPublish(env def.SessionData, bot platform.Platform, projectID string) def.SessionData {
	// Check all users
	users := utils.GetAllUsers(projectID)
	log.Printf("Retrieved %d users", len(users))
	for _, user := range users {
		env = utils.UpdateUserInSession(env, user)

		// Sync platform identity so the bot knows who to message
		env.User.Id = user.Id
		env.User.Firstname = user.Firstname
		env.User.Lastname = user.Lastname
		env.User.Username = user.Username
		env.User.Type = user.Type

		env = utils.SetUserAction(env, app.CMD_DEVO)

		env = HandleSubscriptionLogic(env, bot)
	}

	return env
}

func SubscriptionHandler(localSecrets *secrets.SecretsData) {
	env := def.SessionData{}

	bot := platform.NewTelegram(localSecrets.TELEGRAM_ID, localSecrets.TELEGRAM_ADMIN_ID)

	// log.Printf("Loaded secrets...")

	env.Props = map[string]interface{}{"ResourcePath": "/go/bin/"}

	// TODO: Iterate through types
	env.Type = def.TYPE_TELEGRAM

	env = HandleSubscriptionPublish(env, bot, localSecrets.PROJECT_ID)
	log.Printf("Handled bot publish...")

	if !bot.Post(env) {
		log.Printf("This message was not translatable from bot language")
		return
	}
}
