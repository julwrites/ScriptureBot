package scripturebot

import botsecrets "github.com/julwrites/BotSecrets"

// Struct definitions for bot

type UserConfig struct {
	Version       string `datastore:""`
	Timezone      string `datastore:""`
	Subscriptions string `datastore:""`
}

type UserData struct {
	Firstname string     `datastore:""`
	Lastname  string     `datastore:""`
	Username  string     `datastore:""`
	Id        string     `datastore:""`
	Type      string     `datastore:""`
	Config    UserConfig `datastore:",flatten"`
}

type MessageData struct {
	Command string
	Message string
}

type ResponseData struct {
	Message string
	Options string
}

type SessionData struct {
	Secrets botsecrets.SecretsData
	Type    string
	User    UserData
	Msg     MessageData
	Res     ResponseData
}
