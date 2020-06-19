package app

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/julwrites/BotMultiplexer/pkg/def"
	"github.com/julwrites/ScriptureBot/pkg/utils"
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
	Title  string     `yaml:"Title"`
	Verses []TMSVerse `yaml:"Verses"`
}

type TMSSeries struct {
	ID    string    `yaml:"ID"`
	Title string    `yaml:"Title"`
	Packs []TMSPack `yaml:"Packs"`
}

type TMSDatabase struct {
	Series []TMSSeries `yaml:"Series"`
}

type SeriesSelector func(TMSSeries) bool
type PackSelector func(TMSPack) bool
type VerseSelector func(TMSVerse) bool

func QueryTMSDatabase(db TMSDatabase, s SeriesSelector, p PackSelector, v VerseSelector) (TMSPack, TMSVerse, error) {
	for _, series := range db.Series {
		if s(series) {
			for _, pack := range series.Packs {
				if p(pack) {
					for _, verse := range pack.Verses {
						if v(verse) {
							return pack, verse, nil
						}
					}
				}
			}
		}
	}

	return TMSPack{}, TMSVerse{}, errors.New("Could not find any associated verse")
}

func GetTMSData(dataPath string) TMSDatabase {
	var path []string
	path = append(path, dataPath)
	path = append(path, "tms_data.yaml")

	data, readErr := ioutil.ReadFile(strings.Join(path, "/"))
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

type TMSQueryType string

const (
	ID        TMSQueryType = "ID"
	Tag       TMSQueryType = "Tag"
	Reference TMSQueryType = "Reference"
	Null      TMSQueryType = "0"
)

func IdentifyQuery(db TMSDatabase, query string) TMSQueryType {
	query = strings.Trim(query, " \t\n")

	if strings.Contains(query, ":") {
		parts := strings.Split(query, ":")
		if len(parts) == 2 {
			if strings.ContainsAny(parts[1], "1234567890") {
				return Reference
			}
		}
	}

	for _, series := range db.Series {
		for _, pack := range series.Packs {
			if strings.Contains(strings.ToUpper(query), pack.ID) && strings.ContainsAny(query, "1234567890") {
				return ID
			}
		}
	}

	if len(query) > 0 && !strings.ContainsAny(query, "1234567890") {
		return Tag
	}

	return Null
}

func FormatQuery(query string, t TMSQueryType) string {
	switch t {
	case ID:
		query = strings.ToUpper(query)
		query = strings.ReplaceAll(query, " \t\n", "")
		break
	case Tag:
		query = strings.ToUpper(query)
		query = strings.ReplaceAll(query, " \t\n", "")
		break
	case Reference:
		doc := utils.QueryBiblePassage(query, "NIV")
		query = GetReference(doc)
		break
	}

	return query
}

func GetTMSVerse(env def.SessionData) def.SessionData {
	tmsDB := GetTMSData(env.ResourcePath)

	if len(env.Msg.Message) == 0 {
		log.Printf("Activating action /tms")

		var series []string
		for _, s := range tmsDB.Series {
			series = append(series, s.ID)
		}

		env.User.Action = CMD_TMS
		env.Res.Message = fmt.Sprintf("Tell me which TMS verse you would like using the number (e.g. A1) the reference (e.g. 2 Corinthians 5 : 17)\nAlternatively, give me a topic and I'll try to find a suitable verse!\n\nSupported TMS Series':%s", strings.Join(series, "\n-"))
	} else {
		// Identify the type of query
		queryType := IdentifyQuery(tmsDB, env.Msg.Message)

		query := FormatQuery(env.Msg.Message, queryType)

		var pack TMSPack
		var verse TMSVerse
		var err error

		switch queryType {
		case ID:
			pack, verse, err = QueryTMSDatabase(tmsDB, func(TMSSeries) bool { return true },
				func(pack TMSPack) bool { return strings.Contains(query, pack.ID) },
				func(verse TMSVerse) bool { return strings.Compare(query, verse.ID) == 0 })
			break
		case Tag:
			pack, verse, err = QueryTMSDatabase(tmsDB, func(TMSSeries) bool { return true },
				func(TMSPack) bool { return true },
				func(verse TMSVerse) bool {
					for _, tag := range verse.Tags {
						if strings.Contains(query, tag) {
							return true
						}
					}
					return false
				})
			break
		case Reference:
			pack, verse, err = QueryTMSDatabase(tmsDB, func(TMSSeries) bool { return true },
				func(TMSPack) bool { return true },
				func(verse TMSVerse) bool { return strings.Compare(query, verse.Reference) == 0 })
			break
		}

		if err != nil {
			log.Printf("Query TMS Database failed %v", err)
		}

		env.User.Action = ""
		env.Msg.Message = verse.Reference
		env = GetBiblePassage(env)

		if len(env.Res.Message) != 0 {
			env.Res.Message = fmt.Sprintf("_%s_\n*%s*\n%s\n*%s*", pack.Title, verse.Title, env.Res.Message, verse.Reference)

			log.Printf("%s", env.Res.Message)
		} else {
			env.Res.Message = fmt.Sprintf("I couldn't find a relevant verse")
			log.Printf("Failed to retrieve the verse")
		}

	}

	return env
}
