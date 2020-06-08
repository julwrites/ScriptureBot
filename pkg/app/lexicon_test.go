package app

import (
	"testing"

	"github.com/julwrites/BotMultiplexer/pkg/def"
	"github.com/julwrites/ScriptureBot/pkg/utils"
)

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
