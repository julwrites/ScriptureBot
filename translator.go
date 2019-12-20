package scripturebot

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

// Translator component to handle translation of a HTTP payload into a
// consistent format, and to translate that format back into a HTTP payload
// for posting

func TranslateToProps(req *http.Request, env *SessionData) bool {
	reqBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Fatalf("Error occurred reading http request: %s", err)
		return false
	}
	log.Printf("Request body: %s", reqBody)

	switch env.Type {
	case TYPE_TELEGRAM:
		log.Printf("Parsing Telegram message")

		var data TelegramRequest
		json.Unmarshal(reqBody, &data)

		env.Props.User.Firstname = data.Message.Sender.Firstname
		env.Props.User.Lastname = data.Message.Sender.Lastname
		env.Props.User.Username = data.Message.Sender.Username
		env.Props.User.Id = data.Message.Sender.Id

		log.Printf("User: %s %s | %s : %s", env.Props.User.Firstname, env.Props.User.Lastname, env.Props.User.Username, env.Props.User.Id)

		tokens := strings.Split(data.Message.Text, " ")
		if strings.Index(tokens[0], "/") == 0 {
			env.Props.Command = tokens[0]
		}

		env.Props.Message = strings.Replace(data.Message.Text, env.Props.Command, "", 1)

		log.Printf("Message: %s | %s", env.Props.Command, env.Props.Message)

		return true
	}

	return false
}

func TranslateToHttp(props *SessionData) bool {
	return false
}
