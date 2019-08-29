package main

// Translator methods
import (
	"gopkg.in/yaml.v2"

	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Secrets struct {
	TELEGRAM_ID string

	ADMIN_ID string

	PROJECT_ID string
}

func TranslateToProps(req *http.Request, props *map[string]string) bool {
	data, err := ioutil.ReadFile("secrets.yaml")
	if err != nil {
		log.Fatalf("Error reading secrets: %v", err)
	}

	env := Secrets{}

	err = yaml.Unmarshal([]byte(data), &env)
	if err != nil {
		log.Fatalf("Error unmarshaling secrets: %v", err)
	}

	if req.URL.Path == ("/" + env.TELEGRAM_ID) {
		log.Printf("Telegram message")
		return true
	}

	return false
}

func TranslateToHttp(props *map[string]string) bool {
	return false
}

// Bot methods
func HandleBotLogic(props *map[string]string) bool {
	return false
}

func botHandler(res http.ResponseWriter, req *http.Request) {
	props := map[string]string{}

	if !TranslateToProps(req, &props) {
		log.Printf("This message was not translatable to bot language")
		return
	}

	if !HandleBotLogic(&props) {
		log.Printf("This message was not handled by bot")
		return
	}

	if !TranslateToHttp(&props) {
		log.Printf("This message was not translatable from bot language")
		return
	}
}

func main() {
	http.HandleFunc("/", botHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
