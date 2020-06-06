package app

import (
	"testing"

	"github.com/julwrites/BotMultiplexer/pkg/def"

	"github.com/julwrites/ScriptureBot/pkg/utils"
)

func TestGetReference(t *testing.T) {
	doc := utils.QueryBiblePassage("gen 1", "NIV")

	ref := GetReference(doc)

	if ref != "Genesis 1" {
		t.Errorf("Failed TestGetReference")
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
