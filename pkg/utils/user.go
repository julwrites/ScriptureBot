package utils

import (
	"log"

	"github.com/julwrites/BotPlatform/pkg/def"
	"github.com/julwrites/ScriptureBot/pkg/secrets"
)

type User struct {
	// Id is the user's identifier on the chat platform (e.g., Telegram user ID).
	// This field is not used as a database index - the datastore key serves as the primary identifier.
	Id        string `datastore:",noindex"`
	Username  string `datastore:",noindex"`
	Firstname string `datastore:",noindex"`
	Lastname  string `datastore:",noindex"`
	Type      string `datastore:",noindex"`
	Action    string `datastore:",noindex"`
	Config    string `datastore:",noindex"`
}

func GetUserFromSession(env def.SessionData) User {
	if u, ok := env.Props["User"].(User); ok {
		return u
	}
	return User{}
}

func UpdateUserInSession(env def.SessionData, user User) def.SessionData {
	if env.Props == nil {
		env.Props = make(map[string]interface{})
	}
	env.Props["User"] = user
	return env
}

func SetUserAction(env def.SessionData, action string) def.SessionData {
	user := GetUserFromSession(env)
	user.Action = action
	return UpdateUserInSession(env, user)
}

func SetUserConfig(env def.SessionData, config string) def.SessionData {
	user := GetUserFromSession(env)
	user.Config = config
	return UpdateUserInSession(env, user)
}

func GetUserAction(env def.SessionData) string {
	return GetUserFromSession(env).Action
}

func GetUserConfig(env def.SessionData) string {
	return GetUserFromSession(env).Config
}

func GetResourcePath(env def.SessionData) string {
	if s, ok := env.Props["ResourcePath"].(string); ok {
		return s
	}
	return ""
}

func IsAdmin(env def.SessionData) bool {
	adminID, err := secrets.Get("TELEGRAM_ADMIN_ID")
	if err != nil {
		log.Printf("Failed to get admin ID: %v", err)
		return false
	}
	return env.User.Id == adminID
}
