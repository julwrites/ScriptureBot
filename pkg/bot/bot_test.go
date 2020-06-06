package bot

import (
	"testing"

	"github.com/julwrites/BotMultiplexer/pkg/def"
	"github.com/julwrites/ScriptureBot/pkg/utils"
)

func TestRunCommands(t *testing.T) {
	var env def.SessionData
	var conf utils.UserConfig
	conf.Version = "NIV"
	env.User.Config = utils.SerializeUserConfig(conf)
	env.Msg.Message = "psalm 1"

	env = RunCommands(env)

	if len(env.Res.Message) < 100 {
		t.Errorf("Failed TestRunCommands Passage command")
	}
}
