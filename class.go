package scripturebot

// Struct definitions for bot

type SecretsData struct {
	TELEGRAM_ID string

	ADMIN_ID string

	PROJECT_ID string
}

type SessionData struct {
	Secrets SecretsData

	Props *map[string]string
}
