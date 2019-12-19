package scripturebot

import botsecrets "github.com/julwrites/BotSecrets"

// Struct definitions for bot

type SessionData struct {
	Secrets botsecrets.SecretsData

	Props *map[string]string
}
