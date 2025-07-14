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

		if len(brp.Verses[0]) < 1 {
			t.Errorf("Failed to get DiscipleshipJournal References")
		}
	}
}

func TestGetDesiringGodArticles(t *testing.T) {
	articles := GetDesiringGodArticles()

	if len(articles) == 0 {
		t.Errorf("Failed TestGetDesiringGodArticles, no articles found")
	}
}

func TestGetUtmostForHisHighestArticles(t *testing.T) {
	articles := GetUtmostForHisHighestArticles()

	if len(articles) == 0 {
		t.Errorf("Failed TestGetUtmostForHisHighestArticles, no articles found")
	}

	// Print the content of the first article for verification
	if len(articles) > 0 {
		t.Logf("First Utmost For His Highest Article Link: %s", articles[0].Link)
	}
}

func TestGetDevotionalData(t *testing.T) {
	var env def.SessionData

	env.ResourcePath = "../../resource"

	env.Res = GetDevotionalData(env, "DTMSV")

	if len(env.Res.Message) == 0 {
		t.Errorf("Failed TestGetDevotionalData for DTMSV")
	}
}

func TestGetDevo(t *testing.T) {
	// Test initial devo command (no specific devo chosen)
	var env def.SessionData
	env.User.Action = ""
	env.Msg.Message = CMD_DEVO // Simulate user typing /devo

	env = GetDevo(env)
	if len(env.Res.Message) == 0 {
		t.Errorf("Failed TestGetDevo initial, no message")
	}
	if len(env.Res.Affordances.Options) == 0 {
		t.Errorf("Failed TestGetDevo initial, no affordances")
	}

	// Test each specific devotional option
	for devoName, devoCode := range DEVOS {
		t.Run(devoName, func(t *testing.T) {
			var env def.SessionData
			env.User.Action = CMD_DEVO
			env.Msg.Message = devoName
			env.ResourcePath = "../../resource" // Needed for some devo types

			env = GetDevo(env)

			// Check if a message or options are returned
			if len(env.Res.Message) == 0 && len(env.Res.Affordances.Options) == 0 {
				t.Errorf("Failed TestGetDevo for %s: no message or affordances", devoName)
			}

			// Specific checks based on dispatch method
			switch GetDevotionalDispatchMethod(devoCode) {
			case Passage:
				if len(env.Res.Message) == 0 {
					t.Errorf("Failed TestGetDevo for %s (Passage): no message returned", devoName)
				}
			case Keyboard:
				// For Keyboard types, either options should be present or a message (e.g., for N5XBRP rest day)
				if len(env.Res.Affordances.Options) == 0 && len(env.Res.Message) == 0 {
					t.Errorf("Failed TestGetDevo for %s (Keyboard): no affordances or message returned", devoName)
				}
			}
		})
	}
}
