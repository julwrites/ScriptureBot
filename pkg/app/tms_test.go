package app

import (
	"strings"
	"testing"

	"github.com/julwrites/BotPlatform/pkg/def"
	"github.com/julwrites/ScriptureBot/pkg/utils"
)

func TestGetTMSData(t *testing.T) {
	db := GetTMSData("../../resource")

	if len(db.Series) == 0 {
		t.Errorf("Failed to Get TMS Data Series")
	}

	for _, series := range db.Series {
		if len(series.ID) == 0 {
			t.Errorf("Failed to Get TMS Data Series ID")
		}

		if len(series.Title) == 0 {
			t.Errorf("Failed to Get TMS Data Series Title")
		}

		if len(series.Packs) == 0 {
			t.Errorf("Failed to Get TMS Data Series Packs")
			break
		}

		for _, pack := range series.Packs {
			if len(pack.ID) == 0 {
				t.Errorf("Failed to Get TMS Data Pack ID")
			}

			if len(pack.Title) == 0 {
				t.Errorf("Failed to Get TMS Data Pack Title")
			}

			if len(pack.Verses) == 0 {
				t.Errorf("Failed to Get TMS Data Pack Verses")
				break
			}

			for _, verse := range pack.Verses {
				if len(verse.ID) == 0 {
					t.Errorf("Failed to Get TMS Data Verse ID")
				}

				if len(verse.Title) == 0 {
					t.Errorf("Failed to Get TMS Data Verse Title")
				}

				if len(verse.Tags) == 0 {
					t.Errorf("Failed to Get TMS Data Verse Tags")
					break
				}
			}
		}
	}
}

func TestQueryTMSDatabase(t *testing.T) {
	db := GetTMSData("../../resource")

	var pack TMSPack
	var verse TMSVerse
	var err error

	pack, verse, err = QueryTMSSeries(db,
		func(s TMSSeries) bool {
			if strings.Contains(s.Title, "A") {
				return true
			}
			return false
		})

	if err != nil {
		t.Errorf("Failed TestQueryTMSDatabase series query")
	}

	pack, verse, err = QueryTMSPack(db,
		func(p TMSPack) bool {
			if strings.Contains(p.Title, "A") {
				return true
			}
			return false
		})

	if err != nil {
		t.Errorf("Failed TestQueryTMSDatabase pack query")
	}
	if !strings.Contains(pack.Title, "A") {
		t.Errorf("Failed TestQueryTMSDatabase pack query validity")
	}

	pack, verse, err = QueryTMSVerse(db,
		func(v TMSVerse) bool {
			if strings.Contains(v.Reference, "2 Corinthians 5 :17") {
				return true
			}
			return false
		})

	if err != nil {
		t.Errorf("Failed TestQueryTMSDatabase verse query")
	}
	if !strings.Contains(verse.Reference, "Corinthians") {
		t.Errorf("Failed TestQueryTMSDatabase verse query validity")
	}

	pack, verse, err = QueryTMSVerse(db,
		func(v TMSVerse) bool {
			for _, tag := range v.Tags {
				if strings.Contains(tag, "Prosperous") {
					return true
				}
			}
			return false
		})

	if err != nil {
		t.Errorf("Failed TestQueryTMSDatabase tag query")
	}
	if !strings.Contains(verse.Reference, "Joshua") {
		t.Errorf("Failed TestQueryTMSDatabase tag query validity")
	}
}

func TestIdentifyQuery(t *testing.T) {
	db := GetTMSData("../../resource")

	var queryType TMSQueryType

	queryType = IdentifyQuery(db, "E5")

	if queryType != ID {
		t.Errorf("Failed TestIdentifyQuery ID")
	}

	queryType = IdentifyQuery(db, "Gal 2:20")

	if queryType != Reference {
		t.Errorf("Failed TestIdentifyQuery Reference")
	}

	queryType = IdentifyQuery(db, "grace")

	if queryType != Tag {
		t.Errorf("Failed TestIdentifyQuery Word")
	}
}

func TestGetTMSVerse(t *testing.T) {
	var env def.SessionData
	var conf utils.UserConfig
	conf.Version = "NIV"
	env.User.Config = utils.SerializeUserConfig(conf)
	env.ResourcePath = "../../resource"

	env.Msg.Message = "A1"
	env = GetTMSVerse(env)

	if len(env.Res.Message) == 0 {
		t.Errorf("Failed TestGetTMSVerse basic scenario")
	}

	env.Msg.Message = "2 Corinthians 5:17"
	env = GetTMSVerse(env)

	if len(env.Res.Message) == 0 {
		t.Errorf("Failed TestGetTMSVerse reference scenario")
	}

	env.Msg.Message = "John 13:34-35"
	env = GetTMSVerse(env)

	if len(env.Res.Message) == 0 {
		t.Errorf("Failed TestGetTMSVerse reference scenario")
	}

	env.Msg.Message = "grace"
	env = GetTMSVerse(env)

	if len(env.Res.Message) == 0 {
		t.Errorf("Failed TestGetTMSVerse word scenario")
	}

	env.Msg.Message = "F5"
	env = GetTMSVerse(env)

	if len(env.Res.Message) == 0 {
		t.Errorf("Failed TestGetTMSVerse error scenario")
	}
}
