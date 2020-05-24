package main

import bmul "github.com/julwrites/BotMultiplexer"

func SetVersion(env *bmul.SessionData) {
	if env.User.Action == CMD_VERSION {
		env.User.Action = ""
		UpdateUser(&env.User, env)
	} else {
		env.User.Action = CMD_VERSION
		UpdateUser(&env.User, env)
	}
}
