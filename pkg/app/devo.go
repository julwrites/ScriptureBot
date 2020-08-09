package app

import (
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/julwrites/BotPlatform/pkg/def"
)

const (
	MCBRP string = "MCBRP"
	DJBRP string = "DJBRP"
	DGORG string = "DGORG"
)

var DEVOS = map[string]string{
	"M'Cheyne Bible Reading Plan":             MCBRP,
	"Discipleship Journal Bible Reading Plan": DJBRP,
	"Desiring God Articles":                   DGORG,
}

func SanitizeDevo(msg string) (string, error) {
	msg = strings.Trim(msg, " ")
	devo, ok := DEVOS[msg]
	if ok {
		return devo, nil
	}
	return "", errors.New(fmt.Sprintf("Devo could not be recognized %s", msg))
}

func GetDevotionalText(devo string) string {
	var text string

	switch devo {
	case MCBRP:
		fallthrough
	case DJBRP:
		text = "Here are today's Bible Reading passages, tap on any one to get the passage!"
		break
	case DGORG:
		text = "Here are today's DesiringGod.org articles, tap on any one to open the article!"
		break
	}

	return text
}

func GetDevotionalData(env def.SessionData, devo string) def.ResponseData {
	var response def.ResponseData

	response.Message = GetDevotionalText(devo)

	switch devo {
	case MCBRP:
		response.Affordances.Options = GetMCheyneReferences()
		break
	case DJBRP:
		response.Affordances.Options = GetDiscipleshipJournalReferences(env)
		if len(response.Affordances.Options) == 0 {
			response.Message = "Take this time today to reflect over this week's devotions"
		}
		break
	case DGORG:
		response.Affordances.Options = GetDesiringGodArticles()
		response.Affordances.Inline = true
		break
	default:
		response.Affordances.Remove = true
		break
	}

	if len(response.Affordances.Options) > 0 {
		response.Affordances.Options = append(response.Affordances.Options, def.Option{Text: CMD_CLOSE})
	}

	return response
}

func GetDevo(env def.SessionData) def.SessionData {
	switch env.User.Action {
	case CMD_DEVO:
		log.Printf("Detected existing action /devo")

		devo, err := SanitizeDevo(env.Msg.Message)
		if err == nil {
			log.Printf("Devotional is valid, retrieving %s", devo)

			// Retrieve devotional
			env.Res = GetDevotionalData(env, devo)

			env.User.Action = ""
		} else {
			log.Printf("SanitizeDevo failed %v", err)
			env.Res.Message = "I didn't recognize that devo, please try again"
		}

		break
	default:
		log.Printf("Activating action /devo")

		var options []def.Option

		for k, _ := range DEVOS {
			options = append(options, def.Option{Text: k})
		}

		env.Res.Affordances.Options = options

		env.User.Action = CMD_DEVO

		env.Res.Message = "Choose a Devotional to read!"

		break
	}

	return env
}
