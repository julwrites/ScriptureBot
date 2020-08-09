package app

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"time"

	"github.com/julwrites/BotPlatform/pkg/def"
	"github.com/julwrites/ScriptureBot/pkg/utils"
	"golang.org/x/net/html"
	"gopkg.in/yaml.v2"
)

func GetMCheyneHtml() *html.Node {
	query := fmt.Sprintf("http://www.edginet.org/mcheyne/rss_feed.php?type=rss_2.0&tz=0&cal=classic&bible=esv")

	return utils.QueryHtml(query)
}

func GetMCheyneReferences() []def.Option {
	var options []def.Option

	doc := GetMCheyneHtml()

	titleNodes := utils.FilterTree(doc, func(node *html.Node) bool {
		if node.Data == "title" {
			return true
		}
		return false
	})

	for _, node := range titleNodes {
		for child := node.FirstChild; child != nil; child = child.NextSibling {
			ref := child.Data
			ref = strings.Split(ref, " (")[0]
			if CheckBibleReference(ref) {
				options = append(options, def.Option{Text: ref})
			}
		}
	}

	return options
}

type DiscipleshipJournalDevo struct {
	Verses []string `yaml:"Verses,flow"`
}

type DiscipleshipJournalBRP struct {
	BibleReadingPlan []DiscipleshipJournalDevo `yaml:"BibleReadingPlan"`
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
		for _, r := range brp.Verses {
			if r == "Reflection" {
				continue
			}
			options = append(options, def.Option{Text: r})
		}
	}

	return options
}
