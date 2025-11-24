package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type TelegramResponse struct {
	Ok          bool   `json:"ok"`
	Result      bool   `json:"result"`
	Description string `json:"description"`
}

func main() {
	serviceURL := flag.String("url", "", "The Cloud Run Service URL")
	telegramToken := flag.String("token", "", "The Telegram Bot Token")

	flag.Parse()

	// Fallback to env var for token if not provided
	if *telegramToken == "" {
		*telegramToken = os.Getenv("TELEGRAM_ID")
	}

	if *serviceURL == "" {
		log.Fatal("Error: -url flag is required")
	}
	if *telegramToken == "" {
		log.Fatal("Error: -token flag or TELEGRAM_ID env var is required")
	}

	// Clean up URL
	baseURL := strings.TrimRight(*serviceURL, "/")

	// Construct Webhook URL
	// Based on main.go, the handler is at /<token>
	webhookURL := fmt.Sprintf("%s/%s", baseURL, *telegramToken)

	log.Printf("Setting webhook to: %s/<HIDDEN_TOKEN>", baseURL)

	if err := setWebhook(*telegramToken, webhookURL); err != nil {
		log.Fatalf("Failed to set webhook: %v", err)
	}

	log.Println("Webhook set successfully")
}

var telegramAPIBase = "https://api.telegram.org/bot"

func setWebhook(token, webhookURL string) error {
	apiURL := fmt.Sprintf("%s%s/setWebhook?url=%s", telegramAPIBase, token, url.QueryEscape(webhookURL))

	resp, err := http.Get(apiURL)
	if err != nil {
		return fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("telegram API returned status: %s", resp.Status)
	}

	var result TelegramResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return fmt.Errorf("failed to decode response: %w", err)
	}

	if !result.Ok {
		return fmt.Errorf("telegram API error: %s", result.Description)
	}

	return nil
}
