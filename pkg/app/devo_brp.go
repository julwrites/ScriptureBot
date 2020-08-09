package app

import (
	"io/ioutil"
	"log"
	"strings"
	"time"

	"github.com/julwrites/BotPlatform/pkg/def"
	"github.com/julwrites/ScriptureBot/pkg/utils"
	"golang.org/x/net/html"
	"gopkg.in/yaml.v2"
)

func GetMCheyneReferences() []def.Option {
	var options []def.Option

	doc := utils.QueryMCheyne()

	titleNodes := utils.FilterTree(doc, func(node *html.Node) bool {
		for _, attr := range node.Attr {
			if attr.Key == "class" && strings.Contains(attr.Val, "title") {
				return true
			}
		}
		return false
	})

	for _, node := range titleNodes {
		ref := node.Data
		ref = strings.ReplaceAll(ref, "(", "")
		ref = strings.ReplaceAll(ref, ")", "")
		options = append(options, def.Option{Text: ref})
	}

	return options
}

type DiscipleshipJournalBRP struct {
	BibleReadingPlan [][]string `yaml:"Tags,flow"`
}

func GetDiscipleshipJournalDatabase(dataPath string) DiscipleshipJournalBRP {

	var path []string
	path = append(path, dataPath)
	path = append(path, "djbr_data.yaml")

	data, readErr := ioutil.ReadFile(strings.Join(path, "/"))
	if readErr != nil {
		log.Printf("Error reading DJBR data file: %v", readErr)
	}

	var djBRP DiscipleshipJournalBRP
	yamlErr := yaml.Unmarshal(data, &djBRP)
	if yamlErr != nil {
		log.Printf("Error reading DJBR data from yaml: %v", yamlErr)
	}

	return djBRP
}

func GetDiscipleshipJournalReferences(env def.SessionData) []def.Option {
	var options []def.Option

	djBRP := GetDiscipleshipJournalDatabase(env.ResourcePath)

	length := len(djBRP.BibleReadingPlan) / 12

	// We will read the entry using the date, format: Year, Month, Day
	_, month, day := time.Now().Date()
	brp := djBRP.BibleReadingPlan[(int(month)-1)*length+(day-1)]

	if day < length {
		for _, r := range brp {
			if r == "Reflection" {
				continue
			}
			options = append(options, def.Option{Text: r})
		}
	}

	return options
}
