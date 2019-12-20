package scripturebot

import botsecrets "github.com/julwrites/BotSecrets"

// Struct definitions for bot

type UserData struct {
	Firstname string
	Lastname  string
	Username  string
	Id        string
	Options   string
}

type MessageData struct {
	User    UserData
	Command string
	Message string
}

type SessionData struct {
	Secrets botsecrets.SecretsData
	Type    string
	Props   MessageData
}
