package scripturebot

import (
	"encoding/json"
	"log"
	"strconv"
	"strings"
)

type TelegramSender struct {
	Id        int    `json:"id"`
	Bot       bool   `json:"is_bot"`
	Firstname string `json:"first_name"`
	Lastname  string `json:"last_name"`
	Username  string `json:"username"`
	Language  string `json:"langauge_code"`
}

type TelegramChat struct {
	Id        int    `json:"id"`
	Firstname string `json:"first_name"`
	Lastname  string `json:"last_name"`
	Username  string `json:"username"`
	Type      string `json:"type"`
}

type TelegramMessage struct {
	Sender TelegramSender `json:"from"`
	Chat   TelegramChat   `json:"chat"`
	Text   string         `json:"text"`
}

type TelegramRequest struct {
	Message TelegramMessage `json:"message"`
}

func TelegramTranslate(body []byte, env *SessionData) bool {
	log.Printf("Parsing Telegram message")

	var data TelegramRequest
	err := json.Unmarshal(body, &data)
	if err != nil {
		log.Fatalf("Failed to unmarshal request body: %v", err)
		return false
	}

	env.User.Firstname = data.Message.Sender.Firstname
	env.User.Lastname = data.Message.Sender.Lastname
	env.User.Username = data.Message.Sender.Username
	env.User.Id = strconv.Itoa(data.Message.Sender.Id)
	env.User.Type = TYPE_TELEGRAM

	log.Printf("User: %s %s | %s : %s", env.User.Firstname, env.User.Lastname, env.User.Username, env.User.Id)

	tokens := strings.Split(data.Message.Text, " ")
	if strings.Index(tokens[0], "/") == 0 {
		env.Msg.Command = tokens[0]
	}

	env.Msg.Message = strings.Replace(data.Message.Text, env.Msg.Command, "", 1)

	log.Printf("Message: %s | %s", env.Msg.Command, env.Msg.Message)

	return true
}

func TranslateTelegram(env *SessionData) bool {
	return true
}
