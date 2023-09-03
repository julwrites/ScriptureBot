package app

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"strings"

	"github.com/julwrites/BotPlatform/pkg/def"
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

func QueryTMSSeries(db TMSDatabase, s SeriesSelector) (TMSPack, TMSVerse, error) {
	for _, series := range db.Series {
		if s(series) {
			pack := series.Packs[rand.Int()%len(series.Packs)]
			// TODO: Select one randomly from the series
			verse := pack.Verses[rand.Int()%len(pack.Verses)]

			return pack, verse, nil
		}
	}

	return TMSPack{}, TMSVerse{}, errors.New("Could not find any associated verse")
}

func QueryTMSPack(db TMSDatabase, p PackSelector) (TMSPack, TMSVerse, error) {
	for _, series := range db.Series {
		for _, pack := range series.Packs {
			if p(pack) {
				verse := pack.Verses[rand.Int()%len(pack.Verses)]
				return pack, verse, nil
			}
		}
	}

	return TMSPack{}, TMSVerse{}, errors.New("Could not find any associated verse")
}

func QueryTMSVerse(db TMSDatabase, v VerseSelector) (TMSPack, TMSVerse, error) {
	for _, series := range db.Series {
		for _, pack := range series.Packs {
			for _, verse := range pack.Verses {
				if v(verse) {
					return pack, verse, nil
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
		doc := GetPassageHtml(query, "NIV")
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
		series = append(series, "")
		for _, s := range tmsDB.Series {
			series = append(series, s.ID)
		}

		env.User.Action = CMD_TMS
		env.Res.Message = fmt.Sprintf("Tell me which TMS verse you would like using the number (e.g. A1) the reference (e.g. 2 Corinthians 5 : 17)\nAlternatively, give me a topic and I'll try to find a suitable verse!\nSupported TMS Series:\n%s", strings.Join(series, "\n- "))
	} else {
		// Identify the type of query
		queryType := IdentifyQuery(tmsDB, env.Msg.Message)

		query := FormatQuery(env.Msg.Message, queryType)

		var pack TMSPack
		var verse TMSVerse
		var err error

		switch queryType {
		case ID:
			pack, verse, err = QueryTMSPack(tmsDB,
				func(tPack TMSPack) bool {
					return strings.Contains(strings.ToLower(query), strings.ToLower(pack.ID))
				})
			break
		case Tag:
			pack, verse, err = QueryTMSVerse(tmsDB,
				func(tVerse TMSVerse) bool {
					for _, tag := range tVerse.Tags {
						if strings.Contains(strings.ToLower(query), strings.ToLower(tag)) {
							return true
						}
					}
					return false
				})
			break
		case Reference:
			pack, verse, err = QueryTMSVerse(tmsDB,
				func(tVerse TMSVerse) bool {
					qry := strings.ReplaceAll(strings.ToLower(query), " ", "")
					ref := strings.ReplaceAll(strings.ToLower(tVerse.Reference), " ", "")
					return strings.Contains(qry, ref)
				})
			break
		}

		if err != nil {
			log.Printf("Query TMS Database failed %v", err)
		}

		env.User.Action = ""
		env.Msg.Message = verse.Reference
		env = GetBiblePassage(env)

		if len(env.Res.Message) != 0 {
			env.Res.Message = fmt.Sprintf("_%s_\n*%s*\n%s*%s*", pack.Title, verse.Title, env.Res.Message, verse.Reference)

			log.Printf("%s", env.Res.Message)
		} else {
			env.Res.Message = fmt.Sprintf("I couldn't find a relevant verse")
			log.Printf("Failed to retrieve the verse")
		}

	}

	return env
}
