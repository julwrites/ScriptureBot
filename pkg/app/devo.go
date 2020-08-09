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
	msg = strings.ToUpper(strings.Trim(msg, " "))
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
	case DJBRP:
		text = "Here are today's Bible Reading passages, tap on any one to get the passage!"
		break
	case DGORG:
		break
	}

	return text
}

func GetDevotionalAffordances(devo string) def.ResponseOptions {
	var affordances def.ResponseOptions

	switch devo {
	case MCBRP:
		var options []def.Option
		break
	case DJBRP:
		var options []def.Option
		break
	case DGORG:
		var options []def.Option
		break
	default:
		affordances.Remove = true
		break
	}

	return affordances
}

func GetDevo(env def.SessionData) def.SessionData {
	switch env.User.Action {
	case CMD_DEVO:
		log.Printf("Detected existing action /devo")

		devo, err := SanitizeDevo(env.Msg.Message)
		if err == nil {
			log.Printf("Devotional is valid, retrieving %s", devo)

			// Retrieve devotional
			env.Res.Message = GetDevotionalText(devo)
			env.Res.Affordances = GetDevotionalAffordances(devo)

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
