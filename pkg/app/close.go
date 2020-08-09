package app

import (
	"fmt"
	"math/rand"

	"github.com/julwrites/BotPlatform/pkg/def"
)

var CLOSEMSGS = []string{
	"Okay %s",
	"Got it, %s!",
	"As you wish, %s",
	"Because you said so, %s",
	"I hear and obey, %s",
}

func CloseAction(env def.SessionData) def.SessionData {
	env.Res.Affordances.Remove = true
	env.User.Action = ""

	fmtMessage := CLOSEMSGS[rand.Intn(len(CLOSEMSGS))]

	env.Res.Message = fmt.Sprintf(fmtMessage, env.User.Firstname)

	return env
}
