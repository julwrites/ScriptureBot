package app

import (
	"testing"

	"github.com/julwrites/BotPlatform/pkg/def"
	"github.com/julwrites/ScriptureBot/pkg/utils"
)

func TestGetBiblePassageHtml(t *testing.T) {
	doc := GetPassageHtml("gen 1", "NIV")

	if doc == nil {
		t.Errorf("Could not retrieve bible passage")
	}
}

func TestGetReference(t *testing.T) {
	doc := GetPassageHtml("gen 1", "NIV")

	ref := GetReference(doc)

	if ref != "Genesis 1" {
		t.Errorf("Failed TestGetReference")
	}
}

func TestGetPassage(t *testing.T) {
	doc := GetPassageHtml("gen 1", "NIV")

	passage := GetPassage(doc, "NIV")

	if len(passage) == 0 {
		t.Errorf("Failed TestGetPassage")
	}
}

func TestGetBiblePassage(t *testing.T) {
	var env def.SessionData
	env.Msg.Message = "gen 1"
	var conf utils.UserConfig
	conf.Version = "NIV"
	env.User.Config = utils.SerializeUserConfig(conf)
	env = GetBiblePassage(env)

	if len(env.Res.Message) == 0 {
		t.Errorf("Failed TestGetBiblePassage")
	}
}

func TestCheckBibleReference(t *testing.T) {
	if CheckBibleReference("Genesis 1:1") == false {
		t.Errorf("Failed CheckBibleReference positive test")
	}

	if CheckBibleReference("Some terrible other word") == true {
		t.Errorf("Failed CheckBibleReference negative test")
	}
}
