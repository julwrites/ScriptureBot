package app

import (
	"testing"

	"github.com/julwrites/BotPlatform/pkg/def"
	"github.com/julwrites/ScriptureBot/pkg/utils"
)

func TestGetBibleLexicon(t *testing.T) {
	doc := GetBibleLexicon("beginning", "NIV")

	if doc == nil {
		t.Errorf("Could not retrieve bible lexicon")
	}
}

func TestGetBibleWord(t *testing.T) {
	var env def.SessionData
	env.Msg.Message = "grace"
	var conf utils.UserConfig
	conf.Version = "KJV"
	env.User.Config = utils.SerializeUserConfig(conf)
	env = GetBibleWord(env)

	// if len(env.Res.Message) == 0 {
	// 	t.Errorf("Failed TestGetBibleWord")
	// }
}
