// Brief: Handler for Telegram-specific messages
// Primary responsibility: Receive and handle Telegram messages from Request to Response

package bot

import (
	"log"
	"net/http"

	"github.com/julwrites/BotPlatform/pkg/secrets"
	"github.com/julwrites/ScriptureBot/pkg/utils"

	"github.com/julwrites/BotPlatform/pkg/def"
	"github.com/julwrites/BotPlatform/pkg/platform"
)

func TelegramHandler(res http.ResponseWriter, req *http.Request, secrets *secrets.SecretsData) {
	log.Printf("Loading session data...")
	env, ok := platform.TranslateToProps(req, def.TYPE_TELEGRAM)
	if !ok {
		log.Printf("This message was not translatable to bot language")
		return
	}

	env.Secrets = *secrets
	// log.Printf("Loaded secrets...")

	env.ResourcePath = "/go/bin/"

	user := utils.RegisterUser(env.User, env.Secrets.PROJECT_ID)
	env.User = user
	// log.Printf("Loaded user...")

	env = HandleBotLogic(env)
	// log.Printf("Handled bot logic...")

	if !platform.PostFromProps(env) {
		log.Printf("This message was not translatable from bot language")
		return
	}

	if env.User != user {
		log.Printf("Updating user %v", env.User)
		utils.PushUser(env.User, env.Secrets.PROJECT_ID) // Any change to the user throughout the commands should be put to database
	}
}
