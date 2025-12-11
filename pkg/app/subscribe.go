package app

import (
	"fmt"
	"log"
	"strings"

	"github.com/julwrites/BotPlatform/pkg/def"
	"github.com/julwrites/ScriptureBot/pkg/utils"
)

func UpdateSubscription(env def.SessionData) def.SessionData {
	config := utils.DeserializeUserConfig(utils.GetUserConfig(env))

	switch utils.GetUserAction(env) {
	case CMD_SUBSCRIBE:
		log.Printf("Detected existing action /subscribe")

		env.Msg.Message = strings.ReplaceAll(env.Msg.Message, "(Subscribed)", "")

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
				env.Res.Message = fmt.Sprintf("Got it, I've updated your subscriptions to remove %s", env.Msg.Message)
			} else {
				for _, s := range strings.Split(config.Subscriptions, ",") {
					// If user selected a bible reading plan, we remove any existing bible reading plan
					if GetDevotionalDispatchMethod(devo) == Keyboard && GetDevotionalDispatchMethod(s) == Keyboard {
						continue
					} else {
						subscriptions = append(subscriptions, s)
					}
				}
				// Add the new bible reading plan
				subscriptions = append(subscriptions, devo)
				env.Res.Message = fmt.Sprintf("Got it, I've updated your subscriptions to include %s", env.Msg.Message)
			}

			config.Subscriptions = strings.Join(subscriptions, ",")
			env = utils.SetUserConfig(env, utils.SerializeUserConfig(config))

			env = utils.SetUserAction(env, "")
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
			devo, err := AcronymizeDevo(k)
			if err != nil {
				log.Fatalf("ExpandDevo failed %v", err)
			} else {
				if strings.Contains(config.Subscriptions, devo) {
					options = append(options, def.Option{Text: k + " (Subscribed)"})
				} else {
					options = append(options, def.Option{Text: k})
				}
			}
		}
		options = append(options, def.Option{Text: CMD_CLOSE})

		env.Res.Affordances.Options = options

		env = utils.SetUserAction(env, CMD_SUBSCRIBE)

		env.Res.Message = "Choose a Devotional to receive!"

		break
	}

	return env
}
