package app

import (
	"log"

	bmul "github.com/julwrites/BotMultiplexer"

	"github.com/julwrites/ScriptureBot/pkg/api"
)

func GetBibleWord(env *bmul.SessionData) {
	if len(env.Msg.Message) > 0 {

		doc := api.QueryBibleLexicon(env.Msg.Message, env)

		ref := api.GetReference(doc, env)
		log.Printf("Reference retrieved: %s", ref)

		if len(ref) > 0 {
			log.Printf("Getting passage")
			env.Res.Message = api.GetPassage(doc, env)
		}
	}
}
