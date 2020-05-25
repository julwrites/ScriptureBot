package main

import (
	"errors"
	"fmt"
	"log"
	"strings"

	bmul "github.com/julwrites/BotMultiplexer"
)

var VERSIONS = map[string]string{
	"NIV":  "NIV",
	"ESV":  "ESV",
	"KJV":  "KJV",
	"NASB": "NASB",
	"NLT":  "NLT",
	"AMP":  "AMP",
}

func SanitizeVersion(msg string) (string, error) {
	msg = strings.ToUpper(strings.Trim(msg, " "))
	_, ok := VERSIONS[msg]
	if ok {
		return msg, nil
	}
	return "", errors.New(fmt.Sprintf("Version could not be recognized %s", msg))
}

func SetVersion(env bmul.SessionData) bmul.SessionData {
	config := DeserializeUserConfig(env.User.Config)

	if env.User.Action == CMD_VERSION {
		log.Printf("Detected existing action /version")

		version, err := SanitizeVersion(env.Msg.Message)
		if err == nil {
			log.Printf("Version is valid, setting to %s", version)

			config.Version = version
			env.User.Config = SerializeUserConfig(config)

			env.User.Action = ""
			env.Res.Message = fmt.Sprintf("Got it, I've changed your version to %s", config.Version)
			env.Res.Affordances.Remove = true
		} else {
			log.Printf("SanitizeVersion failed %v", err)
			env.Res.Message = "I didn't recognize that version, please try again"
		}
	} else {
		log.Printf("Activating action /version")

		var options []bmul.Option

		for _, v := range VERSIONS {
			options = append(options, bmul.Option{Text: v})
		}

		env.Res.Affordances.Options = options

		env.User.Action = CMD_VERSION

		env.Res.Message = fmt.Sprintf("Your current version is %s, what would you like to change it to?", config.Version)
	}

	return env
}
