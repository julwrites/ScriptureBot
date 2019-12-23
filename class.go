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
	Id      string
	Chat    string
	Command string
	Message string
}

type Option struct {
	Text string
	Link string
}

type ResponseOptions struct {
	Inline  bool
	Options []Option
	Remove  bool
}

type ResponseData struct {
	Message     string
	Affordances *ResponseOptions
}

type SessionData struct {
	Secrets botsecrets.SecretsData
	Type    string
	Channel string
	User    UserData
	Msg     MessageData
	Res     ResponseData
}
