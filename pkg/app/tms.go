package app

import (
	"io/ioutil"
	"log"

	"github.com/julwrites/BotMultiplexer/pkg/def"
	"gopkg.in/yaml.v2"
)

type TMSVerse struct {
	ID        string   `yaml:"ID"`
	Title     string   `yaml:"Title"`
	Reference string   `yaml:"Reference"`
	Tags      []string `yaml:"Tags,flow"`
}

type TMSPack struct {
	ID     string     `yaml:"ID"`
	Verses []TMSVerse `yaml:"Verses"`
}

type TMSDatabase struct {
	Packs []TMSPack `yaml:"Packs"`
}

func GetTMSData() TMSDatabase {
	data, readErr := ioutil.ReadFile("./tms/tms_data.yaml")
	if readErr != nil {
		log.Printf("Error reading TMS data file: %v", readErr)
	}

	var tmsDB TMSDatabase
	yamlErr := yaml.Unmarshal(data, &tmsDB)
	if yamlErr != nil {
		log.Printf("Error reading TMS data from yaml: %v", yamlErr)
	}

	return tmsDB
}

func GetTMSVerse(env def.SessionData) def.SessionData {
	if len(env.Msg.Message) > 0 {
		tmsDB := GetTMSData()

		// Identify the type of query

		// Check in tmsDB
		for _, pack := range tmsDB.Packs {
			for _, verse := range pack.Verses {
				if verse.ID == env.Msg.Message {
					log.Printf("ID match with %v", verse)
				}
			}
		}

		// TODO: Unset action
	} else {
		// TODO: Set action
		log.Printf("Could not find any message")
	}

	return env
}
