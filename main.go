// Brief: Main entry point of the bot Web App
// Primary responsibility: Receive and identify handler to delegate to

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/julwrites/BotSecrets/pkg/secrets"
	"github.com/julwrites/ScriptureBot/pkg/bot"
)

func bothandler(res http.ResponseWriter, req *http.Request) {
	secretsPath, _ := filepath.Abs("./secrets.yaml")
	secretsData := secrets.LoadSecrets(secretsPath)

	switch strings.Trim(req.URL.EscapedPath(), "\n") {
	case strings.Trim("/"+secretsData.TELEGRAM_ID, "\n"):
		log.Printf("Incoming telegram message")
		bot.TelegramHandler(res, req, &secretsData)
		break
	default:
		log.Printf("No appropriate handler")
	}
}

func main() {
	http.HandleFunc("/", bothandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
