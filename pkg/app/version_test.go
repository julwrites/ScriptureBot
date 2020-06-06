package app

import (
	"testing"

	"github.com/julwrites/BotMultiplexer/pkg/def"
	"github.com/julwrites/ScriptureBot/pkg/utils"
)

func TestSanitizeVersion(t *testing.T) {
	{
		msg := "Niv"
		ver, _ := SanitizeVersion(msg)
		if ver != "NIV" {
			t.Errorf("Failed Sanitize Version basic scenario")
		}
	}
	{
		msg := "asdkfhdsakjfh"
		_, err := SanitizeVersion(msg)
		if err == nil {
			t.Errorf("Failed Sanitize Version error scenario")
		}
	}
}

func TestSetVersion(t *testing.T) {
	var env def.SessionData
	var conf utils.UserConfig
	conf.Version = "NIV"
	env.User.Config = utils.SerializeUserConfig(conf)

	env = SetVersion(env)
	if len(env.Res.Affordances.Options) < 1 {
		t.Errorf("Failed TestSetVersion initial scenario options")
	}
	if len(env.Res.Message) == 0 {
		t.Errorf("Failed TestSetVersion initial scenario message")
	}

	env.User.Action = CMD_VERSION
	env = SetVersion(env)
	if len(env.Res.Message) == 0 {
		t.Errorf("Failed TestSetVersion error scenario message")
	}

	env.Msg.Message = "ESV"
	env = SetVersion(env)
	if len(env.Res.Affordances.Options) < 1 {
		t.Errorf("Failed TestSetVersion fulfillment scenario options")
	}
	if len(env.Res.Message) == 0 {
		t.Errorf("Failed TestSetVersion fulfillment scenario message")
	}
}
