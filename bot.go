// Brief: Bot logic
// Primary responsibility: Top level logic layer for bot

package scripturebot

import botmultiplexer "github.com/julwrites/BotMultiplexer"

func HelpMessage(env *botmultiplexer.SessionData) {
	env.Res.Message = "Hi, this message comes from the bot"
}

func RunCommands(env *botmultiplexer.SessionData) {
	switch env.Msg.Command {
	default:
		if !GetBiblePassage(env) {
			HelpMessage(env)
		}
	}
}

func HandleBotLogic(env *botmultiplexer.SessionData) bool {
	RunCommands(env)

	if len(env.Res.Message) > 0 {
		return true
	}

	return false
}
