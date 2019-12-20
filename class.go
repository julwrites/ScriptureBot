package scripturebot

import botsecrets "github.com/julwrites/BotSecrets"

// Struct definitions for bot

type UserConfig struct {
	Version       string
	Timezone      string
	Subscriptions string
}

type UserData struct {
	Firstname string
	Lastname  string
	Username  string
	Id        string
	Type      string
	Config    UserConfig `datastore:",flatten"`
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
