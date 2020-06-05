package app

import (
	bmul "github.com/julwrites/BotMultiplexer"

	"github.com/julwrites/ScriptureBot/pkg/api"
)

func GetBibleWord(env *bmul.SessionData) {
	if len(env.Msg.Message) > 0 {
		doc := api.QueryBibleLexicon(env.Msg.Message, api.DeserializeUserConfig(env.User.Config).Version)

		if doc != nil {
			// Do something
		}
	}
}
