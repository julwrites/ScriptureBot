package app

import (
	"github.com/julwrites/BotMultiplexer/pkg/def"
	"github.com/julwrites/ScriptureBot/pkg/utils"
)

func GetBibleWord(env *def.SessionData) {
	if len(env.Msg.Message) > 0 {
		doc := utils.QueryBibleLexicon(env.Msg.Message, utils.DeserializeUserConfig(env.User.Config).Version)

		if doc != nil {
			// Do something
		}
	}
}
