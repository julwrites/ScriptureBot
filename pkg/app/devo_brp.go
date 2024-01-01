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

type BibleReadingPlanChapter struct {
	Verses string `yaml:"Verses,flow"`
}

type DailyChapterBRP struct {
	BibleReadingPlan []BibleReadingPlanChapter `yaml:"BibleReadingPlan"`
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

func GetDailyNewTestamentDatabase(dataPath string) DailyChapterBRP {

	var path []string
	path = append(path, dataPath)
	path = append(path, "dailynt_data.yaml")

	data, readErr := ioutil.ReadFile(strings.Join(path, "/"))
	if readErr != nil {
		log.Printf("Error reading DNTBR data file: %v", readErr)
	}

	var DNTBRP DailyChapterBRP
	yamlErr := yaml.Unmarshal(data, &DNTBRP)
	if yamlErr != nil {
		log.Printf("Error reading DNTBR data from yaml: %v", yamlErr)
	}

	return DNTBRP
}

func GetNavigators5xDatabase(dataPath string) DailyChapterBRP {

	var path []string
	path = append(path, dataPath)
	path = append(path, "n5br_data.yaml")

	data, readErr := ioutil.ReadFile(strings.Join(path, "/"))
	if readErr != nil {
		log.Printf("Error reading N5BR data file: %v", readErr)
	}

	var N5XBRP DailyChapterBRP
	yamlErr := yaml.Unmarshal(data, &N5XBRP)
	if yamlErr != nil {
		log.Printf("Error reading N5BR data from yaml: %v", yamlErr)
	}

	return N5XBRP
}

func GetDiscipleshipJournalReferences(env def.SessionData) []def.Option {
	var options []def.Option

	djBRP := GetDiscipleshipJournalDatabase(env.ResourcePath)

	// We will read the entry using the date, format: Year, Month, Day

	dateIndex := time.Now().YearDay()

	if dateIndex >= len(djBRP.BibleReadingPlan) {
		log.Printf("No References found in DJBRP for date: %d", dateIndex)

		return options
	}

	brp := djBRP.BibleReadingPlan[dateIndex]

	for _, r := range brp.Verses {
		if r == "Reflection" {
			continue
		}
		options = append(options, def.Option{Text: r})
	}

	return options
}
func GetDailyNewTestamentReadingReferences(env def.SessionData) string {
	DNTBRP := GetDailyNewTestamentDatabase(env.ResourcePath)

	// We will read the entry using the date, format: Year, Month, Day
	baseline := time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC)
	day := int64(time.Now().Sub(baseline).Hours() / 24)
	day = day % 260
	brp := DNTBRP.BibleReadingPlan[day]

	return brp.Verses
}

func GetNavigators5xRestDayPrompt(env def.SessionData) (string, []def.Option) {
	var options []def.Option

	N5XBRP := GetNavigators5xDatabase(env.ResourcePath)

	// We will read the entry using the date, format: Year, Month, Day
	dateIndex := time.Now().YearDay()

	// This prompt should only be called on the rest days, so we should get back 5 or 6
	weekday := dateIndex % 7
	weekstart := dateIndex - weekday
	for i := 0; i <= weekday; i++ {
		brp := N5XBRP.BibleReadingPlan[weekstart+i]
		options = append(options, def.Option{Text: brp.Verses})
	}

	prompt := `Today is a rest day! Take some time today to dig deeper. 

As a reminder, here are 5 ways to dig deeper:
Pause in your reading to dig into the Bible. Below are 5 different ways to dig deeper each day. These exercises will encourage meditation. Try a single idea for a week to find what works best for you. Remember to keep a pen and paper ready to capture God's insights.

1. Underline or highlight key words or phrases in the Bible passage. Use a pen or highlighter to mark new discoveries from the text.

2. Put it in your own words. Read the passage or verse slowly, then rewrite each phrase or sentence using your own words.

3. Ask and answer questions. Questions unlock new discoveries and meanings. Ask questions about the passage using these words: who, what, why, when, where, or how. Jot down your answers to these questions.

4. Capture the big idea. God's Word communicates big ideas. Periodically ask: What's the big idea in this sentence, paragraph, or chapter?

5. Personalize the meaning. Respond as God speaks to you through the Scriptures. Ask: How could my life be different today as I respond to what I'm reading?

Here are this week's passages!
`

	return prompt, options
}

func GetNavigators5xReferences(env def.SessionData) string {
	N5XBRP := GetNavigators5xDatabase(env.ResourcePath)

	// We will read the entry using the date, format: Year, Month, Day
	day := time.Now().YearDay()
	brp := N5XBRP.BibleReadingPlan[day]

	return brp.Verses
}
