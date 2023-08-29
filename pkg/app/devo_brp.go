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

type BibleReadingPlanDevo struct {
	Verses []string `yaml:"Verses,flow"`
}

type DailyBRP struct {
	BibleReadingPlan []BibleReadingPlanDevo `yaml:"BibleReadingPlan"`
}

func GetDiscipleshipJournalDatabase(dataPath string) DailyBRP {

	var path []string
	path = append(path, dataPath)
	path = append(path, "djbr_data.yaml")

	data, readErr := ioutil.ReadFile(strings.Join(path, "/"))
	if readErr != nil {
		log.Printf("Error reading DJBR data file: %v", readErr)
	}

	var djBRP DailyBRP
	yamlErr := yaml.Unmarshal(data, &djBRP)
	if yamlErr != nil {
		log.Printf("Error reading DJBR data from yaml: %v", yamlErr)
	}

	return djBRP
}

func GetDailyNewTestamentDatabase(dataPath string) DailyBRP {

	var path []string
	path = append(path, dataPath)
	path = append(path, "dailynt_data.yaml")

	data, readErr := ioutil.ReadFile(strings.Join(path, "/"))
	if readErr != nil {
		log.Printf("Error reading DNTBR data file: %v", readErr)
	}

	var dntBRP DailyBRP
	yamlErr := yaml.Unmarshal(data, &dntBRP)
	if yamlErr != nil {
		log.Printf("Error reading DNTBR data from yaml: %v", yamlErr)
	}

	return dntBRP
}

type DailyChapterNTBRP struct {
	Prompt           string                 `yaml:"Prompt"`
	BibleReadingPlan []BibleReadingPlanDevo `yaml:"BibleReadingPlan"`
}

func GetNavigators5xDatabase(dataPath string) DailyChapterNTBRP {

	var path []string
	path = append(path, dataPath)
	path = append(path, "n5br_data.yaml")

	data, readErr := ioutil.ReadFile(strings.Join(path, "/"))
	if readErr != nil {
		log.Printf("Error reading NTBR data file: %v", readErr)
	}

	var ntBRP DailyChapterNTBRP
	yamlErr := yaml.Unmarshal(data, &ntBRP)
	if yamlErr != nil {
		log.Printf("Error reading NTBR data from yaml: %v", yamlErr)
	}

	return ntBRP
}

func GetDiscipleshipJournalReferences(env def.SessionData) []def.Option {
	var options []def.Option

	djBRP := GetDiscipleshipJournalDatabase(env.ResourcePath)

	// We will read the entry using the date, format: Year, Month, Day

	brp := djBRP.BibleReadingPlan[time.Now().YearDay()]

	for _, r := range brp.Verses {
		if r == "Reflection" {
			continue
		}
		options = append(options, def.Option{Text: r})
	}

	return options
}
func GetDailyNewTestamentReadingReferences(env def.SessionData) []def.Option {
	var options []def.Option

	dntBRP := GetDailyNewTestamentDatabase(env.ResourcePath)

	// We will read the entry using the date, format: Year, Month, Day
	baseline := time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC)
	day := int64(time.Now().Sub(baseline).Hours() / 24)
	day = day % 260
	brp := dntBRP.BibleReadingPlan[day]

	for _, r := range brp.Verses {
		options = append(options, def.Option{Text: r})
	}

	return options
}

func GetNavigators5xPrompt(env def.SessionData) string {
	ntBRP := GetNavigators5xDatabase(env.ResourcePath)

	return ntBRP.Prompt
}

func GetNavigators5xReferences(env def.SessionData) []def.Option {
	var options []def.Option

	ntBRP := GetNavigators5xDatabase(env.ResourcePath)

	length := len(ntBRP.BibleReadingPlan) / 12

	// We will read the entry using the date, format: Year, Month, Day
	_, month, day := time.Now().Date()
	brp := ntBRP.BibleReadingPlan[(int(month)-1)*length+(day-1)]

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
