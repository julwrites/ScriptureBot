package app

import (
	"fmt"
	"log"
	"strings"

	"github.com/julwrites/BotPlatform/pkg/def"
	"github.com/julwrites/ScriptureBot/pkg/utils"
)

func UpdateSubscription(env def.SessionData) def.SessionData {
	config := utils.DeserializeUserConfig(env.User.Config)

	switch env.User.Action {
	case CMD_SUBSCRIBE:
		log.Printf("Detected existing action /subscribe")

		devo, err := AcronymizeDevo(env.Msg.Message)
		if err == nil {
			log.Printf("Devotional is valid, retrieving %s", devo)

			var subscriptions []string
			if strings.Contains(config.Subscriptions, devo) {
				// Remove by including all except this
				for _, s := range strings.Split(config.Subscriptions, ",") {
					if s != devo {
						subscriptions = append(subscriptions, s)
					}
				}
			} else {
				// Add selectively
				for _, s := range strings.Split(config.Subscriptions, ",") {
					// Only allow one bible reading plan
					if GetDevotionalType(devo) == BibleReadingPlan && GetDevotionalType(s) == BibleReadingPlan {
						subscriptions = append(subscriptions, devo)
					}
				}
			}
			config.Subscriptions = strings.Join(subscriptions, ",")

			println(config.Subscriptions)
			env.User.Config = utils.SerializeUserConfig(config)

			env.User.Action = ""
			env.Res.Message = fmt.Sprintf("Got it, I've updated your subscriptions to include %s", devo)
			env.Res.Affordances.Remove = true
		} else {
			log.Printf("AcronymizeDevo failed %v", err)
			env.Res.Message = "I didn't recognize that devo, please try again"
		}

		break
	default:
		log.Printf("Activating action /subscribe")

		var options []def.Option

		for k, _ := range DEVOS {
			name, err := AcronymizeDevo(k)
			if err != nil {
				log.Fatalf("ExpandDevo failed %v", err)
			} else {
				if strings.Contains(config.Subscriptions, k) {
					options = append(options, def.Option{Text: name + " (Subscribed)"})
				} else {
					options = append(options, def.Option{Text: name})
				}
			}
		}
		options = append(options, def.Option{Text: CMD_CLOSE})

		env.Res.Affordances.Options = options

		env.User.Action = CMD_SUBSCRIBE

		env.Res.Message = "Choose a Devotional to receive!"

		break
	}

	return env
}
