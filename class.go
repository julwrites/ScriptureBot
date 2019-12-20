package scripturebot

import botsecrets "github.com/julwrites/BotSecrets"

// Struct definitions for bot

type TelegramData struct {
	User    string `json:"user"`
	Message string `json:"text"`
}

type MessageData struct {
	User    string
	Command string
	Message string
}

type UserData struct {
	Firstname string
	Lastname  string
	Username  string
	Id        string
	Options   string
}

type SessionData struct {
	Secrets botsecrets.SecretsData
	Type    string
	Props   MessageData
}
