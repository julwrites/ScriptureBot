// Brief: Handler for Telegram-specific messages
// Primary responsibility: Receive and handle Telegram messages from Request to Response

package bot

import (
	"log"
	"net/http"

	botsecrets "github.com/julwrites/BotSecrets"

	"github.com/julwrites/ScriptureBot/pkg/api"
)

func TelegramHandler(res http.ResponseWriter, req *http.Request, secrets *botsecrets.SecretsData) {
	log.Printf("Loading session data...")
	env, ok := def.TranslateToProps(req, def.TYPE_TELEGRAM)
	if !ok {
		log.Printf("This message was not translatable to bot language")
		return
	}

	env.Secrets = *secrets
	log.Printf("Loaded secrets...")

	env = api.RegisterUser(env.User, env.Secrets.PROJECT_ID)
	log.Printf("Loaded user...")

	env = HandleBotLogic(env)
	log.Printf("Handled bot logic...")

	if !def.PostFromProps(env) {
		log.Printf("This message was not translatable from bot language")
		return
	}

	api.PushUser(env.User, env.Secrets.PROJECT_ID) // Any change to the user throughout the commands should be put to database
}
