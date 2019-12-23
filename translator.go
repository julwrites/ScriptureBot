package scripturebot

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/google/go-cmp/cmp"
)

// Translator component to handle translation of a HTTP payload into a
// consistent format, and to translate that format back into a HTTP payload
// for posting

func CompareAndUpdateUser(env *SessionData) {
	storedUser := QueryUser(env)

	if !cmp.Equal(storedUser, env.User) {
		env.User.Config = storedUser.Config

		log.Printf("Updating user %s", env.User.Username)

		UpdateUser(&env.User, env)
	}
}

func TranslateToProps(req *http.Request, env *SessionData) bool {
	reqBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Fatalf("Error occurred reading http request: %v", err)
		return false
	}
	log.Printf("Request body: %s", strings.ReplaceAll(string(reqBody), "\n", ""))

	translated := false

	switch env.Type {
	case TYPE_TELEGRAM:
		translated = TelegramTranslate(reqBody, env)
	}

	if translated {
		CompareAndUpdateUser(env)
		return translated
	}

	return false
}

func TranslateToHttp(props *SessionData) bool {
	return false
}
