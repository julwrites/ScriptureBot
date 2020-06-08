package app

import (
	"testing"

	"github.com/julwrites/BotMultiplexer/pkg/def"
	"github.com/julwrites/ScriptureBot/pkg/utils"
)

func TestGetTMSVerse(t *testing.T) {
	var env def.SessionData
	var conf utils.UserConfig
	conf.Version = "NIV"
	env.User.Config = utils.SerializeUserConfig(conf)

	env.Msg.Message = "A1"
	env = GetTMSVerse(env)

	if len(env.Res.Message) == 0 {
		t.Errorf("Failed TestGetTMSVerse basic scenario")
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

	env.Msg.Message = "1 John 3:16"
	env = GetTMSVerse(env)

	if len(env.Res.Message) == 0 {
		t.Errorf("Failed TestGetTMSVerse error scenario")
	}
}
