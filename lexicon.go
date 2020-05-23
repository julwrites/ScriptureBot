package main

import (
	"log"

	bmul "github.com/julwrites/BotMultiplexer"
)

func GetBibleWord(env *bmul.SessionData) {
	if len(env.Msg.Message) > 0 {

		doc := QueryBlueLetterBible(env.Msg.Message, env)

		ref := GetReference(doc, env)
		log.Printf("Reference retrieved: %s", ref)

		if len(ref) > 0 {
			log.Printf("Getting passage")
			env.Res.Message = GetPassage(doc, env)
		}
	}
}
