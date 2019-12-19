package ScriptureBot

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

// Secrets handling for bot

func LoadSecrets() SecretsData {
	data, err := ioutil.ReadFile("secrets.yaml")
	if err != nil {
		log.Fatalf("Error reading secrets: %v", err)
	}

	env := SecretsData{}

	err = yaml.Unmarshal([]byte(data), &env)
	if err != nil {
		log.Fatalf("Error unmarshaling secrets: %v", err)
	}

	return env
}
