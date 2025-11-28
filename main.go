// Brief: Main entry point of the bot Web App
// Primary responsibility: Receive and identify handler to delegate to

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/julwrites/ScriptureBot/pkg/bot"
	"github.com/julwrites/ScriptureBot/pkg/secrets"
)

func bothandler(res http.ResponseWriter, req *http.Request) {
	secretsData, err := secrets.LoadSecrets()
	if err != nil {
		log.Fatalf("Failed to load secrets: %v", err)
	}

	switch strings.Trim(req.URL.EscapedPath(), "\n") {
	case strings.Trim("/"+secretsData.TELEGRAM_ID, "\n"):
		log.Printf("Incoming telegram message")
		bot.TelegramHandler(res, req, &secretsData)
		break
	default:
		log.Printf("No appropriate handler")
	}
}

func subscriptionhandler() {
	secretsData, err := secrets.LoadSecrets()
	if err != nil {
		log.Fatalf("Failed to load secrets: %v", err)
	}

	bot.SubscriptionHandler(&secretsData)
}

func main() {
	if os.Getenv("MODE") == "subscription" {
		subscriptionhandler()
	} else {
		http.HandleFunc("/", bothandler)

		port := os.Getenv("PORT")
		if port == "" {
			port = "8080"
			log.Printf("Defaulting to port %s", port)
		}

		log.Printf("Listening on port %s", port)
		log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
	}
}
