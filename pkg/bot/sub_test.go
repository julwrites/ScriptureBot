package bot

import (
	"testing"

	"github.com/julwrites/BotPlatform/pkg/def"
	"github.com/julwrites/ScriptureBot/pkg/utils"
)

func TestHandleSubscriptionLogic(t *testing.T) {
	var env def.SessionData
	env.ResourcePath = "../../resource"

	var conf utils.UserConfig
	conf.Version = "NIV"
	conf.Subscriptions = "DTMSV"
	env.User.Config = utils.SerializeUserConfig(conf)

	env = HandleSubscriptionLogic(env)

	if len(env.Res.Message) == 0 {
		t.Errorf("Failed TestHandleSubscriptionLogic")
	}
}
