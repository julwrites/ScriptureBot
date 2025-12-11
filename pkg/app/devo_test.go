package app

import (
	"testing"
	"time"

	"github.com/julwrites/BotPlatform/pkg/def"
	"github.com/julwrites/ScriptureBot/pkg/utils"
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
	env.Props = map[string]interface{}{"ResourcePath": "../../resource"}

	options := GetDiscipleshipJournalReferences(env)

	if len(options) == 0 {
		djBRP := GetDiscipleshipJournalDatabase(utils.GetResourcePath(env))

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
		t.Logf("Utmost For His Highest Article data: %v", articles)
	}
}

func TestGetDevotionalData(t *testing.T) {
	t.Run("DTMSV", func(t *testing.T) {
		defer UnsetEnv("BIBLE_API_URL")()
		defer UnsetEnv("BIBLE_API_KEY")()
		ResetAPIConfigCache()

		var env def.SessionData
		env.Props = map[string]interface{}{"ResourcePath": "../../resource"}
		env.Res = GetDevotionalData(env, "DTMSV")

		if len(env.Res.Message) == 0 {
			t.Errorf("Failed TestGetDevotionalData for DTMSV")
		}
	})
}

func TestGetDevo(t *testing.T) {
	t.Run("Initial Devo", func(t *testing.T) {
		defer UnsetEnv("BIBLE_API_URL")()
		defer UnsetEnv("BIBLE_API_KEY")()
		ResetAPIConfigCache()

		var env def.SessionData
		env = utils.SetUserAction(env, "")
		env.Msg.Message = CMD_DEVO

		env = GetDevo(env, &MockBot{})
		if len(env.Res.Message) == 0 {
			t.Error("Failed TestGetDevo initial, no message")
		}
		if len(env.Res.Affordances.Options) == 0 {
			t.Error("Failed TestGetDevo initial, no affordances")
		}
	})

	for devoName, devoCode := range DEVOS {
		devoName := devoName
		devoCode := devoCode
		t.Run(devoName, func(t *testing.T) {
			defer UnsetEnv("BIBLE_API_URL")()
			defer UnsetEnv("BIBLE_API_KEY")()
			ResetAPIConfigCache()

			var env def.SessionData
			env = utils.SetUserAction(env, CMD_DEVO)
			env.Msg.Message = devoName
			env.Props = map[string]interface{}{"ResourcePath": "../../resource"}

			env = GetDevo(env, &MockBot{})

			if len(env.Res.Message) == 0 && len(env.Res.Affordances.Options) == 0 {
				t.Fatalf("Failed TestGetDevo for %s: no message or affordances", devoName)
			}

			switch GetDevotionalDispatchMethod(devoCode) {
			case Passage:
				if len(env.Res.Message) == 0 {
					t.Errorf("Expected a message for Passage type devo, got none")
				}
			case Keyboard:
				if len(env.Res.Affordances.Options) == 0 && len(env.Res.Message) == 0 {
					t.Errorf("Expected affordances or a message for Keyboard type devo, got none")
				}
			}
		})
	}
}
