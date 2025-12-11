package utils

import "github.com/julwrites/BotPlatform/pkg/def"

type User struct {
	Id        string `datastore:"-"` // ID is the key
	Username  string `datastore:""`
	Firstname string `datastore:""`
	Lastname  string `datastore:""`
	Type      string `datastore:""`
	Action    string `datastore:""`
	Config    string `datastore:""`
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
