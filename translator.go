package scripturebot

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/google/go-cmp/cmp"
)

// Translator component to handle translation of a HTTP payload into a
// consistent format, and to translate that format back into a HTTP payload
// for posting

func TranslateTelegram(body *string, env *SessionData) bool {
	log.Printf("Parsing Telegram message")

	var data TelegramRequest
	err := json.Unmarshal(body, &data)
	if err != nil {
		log.Fatalf("Failed to unmarshal request body: %v", err)
		return false
	}

	env.Props.User.Firstname = data.Message.Sender.Firstname
	env.Props.User.Lastname = data.Message.Sender.Lastname
	env.Props.User.Username = data.Message.Sender.Username
	env.Props.User.Id = strconv.Itoa(data.Message.Sender.Id)
	env.Props.User.Type = TYPE_TELEGRAM

	log.Printf("User: %s %s | %s : %s", env.Props.User.Firstname, env.Props.User.Lastname, env.Props.User.Username, env.Props.User.Id)

	tokens := strings.Split(data.Message.Text, " ")
	if strings.Index(tokens[0], "/") == 0 {
		env.Props.Command = tokens[0]
	}

	env.Props.Message = strings.Replace(data.Message.Text, env.Props.Command, "", 1)

	log.Printf("Message: %s | %s", env.Props.Command, env.Props.Message)

	return true
}

func CompareAndUpdateUser(env *SessionData) {
	storedUser := QueryUser(env)

	if !cmp.Equal(storedUser, env.Props.User) {
		env.Props.User.Config = storedUser.Config

		UpdateUser(&env.Props.User, env)
	}
}

func TranslateToProps(req *http.Request, env *SessionData) bool {
	reqBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Fatalf("Error occurred reading http request: %v", err)
		return false
	}
	log.Printf("Request body: %s", strings.ReplaceAll(string(reqBody), "\n", ""))

	switch env.Type {
	case TYPE_TELEGRAM:
		return TranslateTelegram(reqBody, env)
	}

	CompareAndUpdateUser(env)

	return false
}

func TranslateToHttp(props *SessionData) bool {
	return false
}
