package scripturebot

func HelpMessage(env *SessionData) {
	env.Res.Message = "Hi, this message comes from the bot"
}

func RunCommands(env *SessionData) {
	switch env.Msg.Command {
	default:
		if !GetBiblePassage(env) {
			HelpMessage(env)
		}
	}
}

func HandleBotLogic(env *SessionData) bool {
	RunCommands(env)

	if len(env.Res.Message) > 0 {
		return true
	}

	return false
}
