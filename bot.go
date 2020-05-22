// Brief: Bot logic
// Primary responsibility: Top level logic layer for bot

package main

import bmul "github.com/julwrites/BotMultiplexer"

func HelpMessage(env *bmul.SessionData) {
	env.Res.Message = "Hi, this message comes from the bot"
}

func RunCommands(env *bmul.SessionData) {
	switch env.Msg.Command {
	default:
		if !GetBiblePassage(env) {
			HelpMessage(env)
		}
	}
}

func HandleBotLogic(env *bmul.SessionData) bool {
	RunCommands(env)

	if len(env.Res.Message) > 0 {
		return true
	}

	return false
}
