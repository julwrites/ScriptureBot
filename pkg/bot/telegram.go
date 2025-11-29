// Brief: Handler for Telegram-specific messages
// Primary responsibility: Receive and handle Telegram messages from Request to Response

package bot

import (
	"io"
	"log"
	"net/http"

	"github.com/julwrites/BotPlatform/pkg/platform"
	"github.com/julwrites/ScriptureBot/pkg/secrets"
	"github.com/julwrites/ScriptureBot/pkg/utils"
)

func TelegramHandler(res http.ResponseWriter, req *http.Request, secrets *secrets.SecretsData) {
	log.Printf("Loading session data...")

	body, err := io.ReadAll(req.Body)
	if err != nil {
		log.Printf("Failed to read request body: %v", err)
		return
	}

	bot := platform.NewTelegram(secrets.TELEGRAM_ID, secrets.TELEGRAM_ADMIN_ID)

	env, err := bot.Translate(body)
	if err != nil {
		log.Printf("This message was not translatable to bot language: %v", err)
		return
	}

	// log.Printf("Loaded secrets...")

	env.ResourcePath = "/go/bin/"

	user := utils.RegisterUser(env.User, secrets.PROJECT_ID)
	env.User = user
	// log.Printf("Loaded user...")

	env = HandleBotLogic(env, bot)
	// log.Printf("Handled bot logic...")

	if !bot.Post(env) {
		log.Printf("This message was not translatable from bot language")
		return
	}

	if env.User != user {
		log.Printf("Updating user %v", env.User)
		utils.PushUser(env.User, secrets.PROJECT_ID) // Any change to the user throughout the commands should be put to database
	}
}
