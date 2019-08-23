package main

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
}

func botHandler(res http.ResponseWriter, req *http.Request) {
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
		return
	}
	log.Printf("This message was not handled")
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
