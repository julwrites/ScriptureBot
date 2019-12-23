package scripturebot

func HandleBotLogic(env *SessionData) bool {
	env.Res.Message = "Hi, this message comes from the bot"

	return false
}
