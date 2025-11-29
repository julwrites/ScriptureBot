package app

import (
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/julwrites/BotPlatform/pkg/platform"

	"github.com/julwrites/BotPlatform/pkg/def"
)

const (
	MCBRP  string = "MCBRP"
	DJBRP  string = "DJBRP"
	DNTBRP string = "DNTBRP"
	N5XBRP string = "N5XBRP"
	DGORG  string = "DGORG"
	DTMSV  string = "DTMSV"
	MUFHH  string = "MUFHH"
)

const (
	Keyboard    string = "Keyboard"
	Passage     string = "Passage"
	MemoryVerse string = "MemoryVerse"
)

var DEVO_NAMES = map[string]string{
	MCBRP:  "M'Cheyne Bible Reading Plan",
	DJBRP:  "Discipleship Journal Bible Reading Plan",
	DNTBRP: "Daily New Testament Reading Plan",
	N5XBRP: "Navigators 5x5x5 New Testament Reading Plan",
	DGORG:  "Desiring God Articles",
	DTMSV:  "Daily Topical Memory System Verse",
	MUFHH:  "My Utmost For His Highest Articles",
}

var DEVOS = map[string]string{
	"M'Cheyne Bible Reading Plan":                 MCBRP,
	"Discipleship Journal Bible Reading Plan":     DJBRP,
	"Daily New Testament Reading Plan":            DNTBRP,
	"Navigators 5x5x5 New Testament Reading Plan": N5XBRP,
	"Desiring God Articles":                       DGORG,
	"Daily Topical Memory System Verse":           DTMSV,
	"My Utmost For His Highest Articles":          MUFHH,
}

func AcronymizeDevo(msg string) (string, error) {
	msg = strings.Trim(msg, " ")
	devo, ok := DEVOS[msg]
	if ok {
		return devo, nil
	}
	return "", errors.New(fmt.Sprintf("Devo could not be recognized %s", msg))
}

func ExpandDevo(msg string) (string, error) {
	msg = strings.Trim(msg, " ")
	devo, ok := DEVO_NAMES[msg]
	if ok {
		return devo, nil
	}
	return "", errors.New(fmt.Sprintf("Devo could not be recognized %s", msg))
}

func GetDevotionalDispatchMethod(devo string) string {
	switch devo {
	case MCBRP:
		return Keyboard
	case DJBRP:
		return Keyboard
	case DNTBRP:
		return Passage
	case N5XBRP:
		return Keyboard
	case DGORG:
		return Keyboard
	case DTMSV:
		return Passage
	case MUFHH:
		return Keyboard
	}

	return ""
}

func GetDevotionalText(devo string) string {
	var text string

	switch devo {
	case MUFHH:
		fallthrough // Same as DGORG
	case DGORG:
		text = "Here are the current articles, tap on any one to open the article!"
	case MCBRP:
		fallthrough // Same as DJBRP
	case DJBRP:
		text = "Here are today's Bible Reading passages, tap on any one to get the passage!"
	case DNTBRP:
		fallthrough
	case N5XBRP:
		fallthrough
	case DTMSV:
		break // No text because we send the text directly
	}

	return text
}

func GetDevotionalData(env def.SessionData, devo string) def.ResponseData {
	var response def.ResponseData

	response.Message = GetDevotionalText(devo)
	log.Printf("Devotional text: %s", response.Message)

	switch devo {
	case MCBRP:
		log.Printf("Retrieving MCheyne Bible Reading Plan")
		response.Affordances.Options = GetMCheyneReferences()
	case DJBRP:
		log.Printf("Retrieving Discipleship Journal Bible Reading Plan")
		response.Affordances.Options = GetDiscipleshipJournalReferences(env)
		if len(response.Affordances.Options) == 0 {
			response.Message = "Take this time today to reflect over this week's devotions"
		}
	case DNTBRP:
		log.Printf("Retrieving Daily New Testament Bible Reading Plan")
		env.Msg.Message = GetDailyNewTestamentReadingReferences(env)
		env = GetBiblePassage(env)
		response = env.Res
	case N5XBRP:
		log.Printf("Retrieving Navigators 5x5x5 New Testament Bible Reading Plan")
		env.Msg.Message = GetNavigators5xReferences(env)
		if len(env.Msg.Message) > 0 {
			env = GetBiblePassage(env)
			response = env.Res
		} else {
			prompt, options := GetNavigators5xRestDayPrompt(env)
			response.Message = prompt
			response.Affordances.Options = options
		}
	case DGORG:
		log.Printf("Retrieving Desiring God Articles")
		response.Affordances.Options = GetDesiringGodArticles()
		response.Affordances.Inline = true
	case DTMSV:
		log.Printf("Retrieving Daily Topical Memory System Verse")
		env.Msg.Message = GetRandomTMSVerse(env)
		env = GetTMSVerse(env)
		response = env.Res
	case MUFHH:
		log.Printf("Retrieving My Utmost For His Highest Articles")
		response.Affordances.Options = GetUtmostForHisHighestArticles()
		response.Affordances.Inline = true
	default:
		response.Affordances.Remove = true
	}

	if len(response.Affordances.Options) > 0 && response.Affordances.Inline == false {
		response.Affordances.Options = append(response.Affordances.Options, def.Option{Text: CMD_CLOSE})
	}

	log.Printf("Devotional response: %+v", response)

	return response
}

func GetDevo(env def.SessionData, bot platform.Platform) def.SessionData {
	switch env.User.Action {
	case CMD_DEVO:
		log.Printf("Detected existing action /devo")

		devo, err := AcronymizeDevo(env.Msg.Message)
		if err == nil {
			log.Printf("Devotional is valid, retrieving %s", devo)

			env.Res.Affordances.Remove = true
			env.Res.Message = "Just a moment..."

			log.Printf("Affordances before posting: %+v", env.Res.Affordances)
			bot.Post(env)

			// Retrieve devotional
			env.Res = GetDevotionalData(env, devo)

			env.User.Action = ""
		} else {
			log.Printf("AcronymizeDevo failed %v", err)
			env.Res.Message = "I didn't recognize that devo, please try again"
		}
	default:
		log.Printf("Activating action /devo")

		var options []def.Option

		for k, _ := range DEVOS {
			options = append(options, def.Option{Text: k})
		}
		options = append(options, def.Option{Text: CMD_CLOSE})

		env.Res.Affordances.Options = options

		env.User.Action = CMD_DEVO

		env.Res.Message = "Choose a Devotional to read!"
	}

	return env
}
