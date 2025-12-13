package bot

import (
	"testing"

	"github.com/julwrites/BotPlatform/pkg/def"
	"github.com/julwrites/ScriptureBot/pkg/app"
	"github.com/julwrites/ScriptureBot/pkg/utils"
)

func TestHandleSubscriptionLogic(t *testing.T) {
	var env def.SessionData
	env.Props = map[string]interface{}{"ResourcePath": "../../resource"}

	var conf utils.UserConfig
	conf.Version = "NIV"
	conf.Subscriptions = "DTMSV"
	env = utils.SetUserConfig(env, utils.SerializeUserConfig(conf))

	env = HandleSubscriptionLogic(env, &app.MockBot{})

	if len(env.Res.Message) == 0 {
		t.Errorf("Failed TestHandleSubscriptionLogic")
	}
}
