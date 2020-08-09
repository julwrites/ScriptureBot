package app

import (
	"testing"
	"time"

	"github.com/julwrites/BotPlatform/pkg/def"
)

func TestGetMCheyneHtml(t *testing.T) {
	doc := GetMCheyneHtml()

	if doc == nil {
		t.Errorf("Failed TestGetMCheyneHtml, no RSS retrieved")
	}
}

func TestGetMCheyneReferences(t *testing.T) {
	options := GetMCheyneReferences()

	if len(options) == 0 {
		t.Errorf("Failed TestGetMCheyneReferences, no options retrieved")
	}
}

func TestGetDiscipleshipJournalDatabase(t *testing.T) {
	db := GetDiscipleshipJournalDatabase("../../resource")

	if len(db.BibleReadingPlan) == 0 {
		t.Errorf("Failed to Get DiscipleshipJournal BibleReadingPlan Data")
	}
}

func TestGetDiscipleshipJournalReferences(t *testing.T) {
	var env def.SessionData

	env.ResourcePath = "../../resource"

	options := GetDiscipleshipJournalReferences(env)

	if len(options) == 0 {
		djBRP := GetDiscipleshipJournalDatabase(env.ResourcePath)

		length := len(djBRP.BibleReadingPlan) / 12

		// We will read the entry using the date, format: Year, Month, Day
		_, month, day := time.Now().Date()
		day = day % length
		brp := djBRP.BibleReadingPlan[(int(month)-1)*length+(day-1)]

		if brp.Verses[0] != "Reflection" {
			t.Errorf("Failed to get DiscipleshipJournal References")
		}
	}
}

func TestGetDesiringGodHtml(t *testing.T) {
	doc := GetDesiringGodHtml()

	if doc == nil {
		t.Errorf("Failed TestGetDesiringGodHtml, no RSS retrieved")
	}
}

func TestGetDesiringGodArticles(t *testing.T) {
	articles := GetDesiringGodArticles()

	if len(articles) == 0 {
		t.Errorf("Failed TestGetDesiringGodArticles, no articles found")
	}
}
