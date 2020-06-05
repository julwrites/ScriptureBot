package app

import (
	"github.com/julwrites/BotMultiplexer/pkg/def"
	"github.com/julwrites/ScriptureBot/pkg/api"
)

func GetBibleWord(env *def.SessionData) {
	if len(env.Msg.Message) > 0 {
		doc := api.QueryBibleLexicon(env.Msg.Message, api.DeserializeUserConfig(env.User.Config).Version)

		if doc != nil {
			// Do something
		}
	}
}
