package app

import (
	"log"
	"testing"

	"github.com/julwrites/BotPlatform/pkg/def"
	"github.com/julwrites/ScriptureBot/pkg/utils"
)

func TestUpdateSubscription(t *testing.T) {
	var env def.SessionData
	var conf utils.UserConfig
	conf.Subscriptions = "MCBRP"
	env.User.Config = utils.SerializeUserConfig(conf)

	env = UpdateSubscription(env)
	if len(env.Res.Affordances.Options) < 1 {
		t.Errorf("Failed TestUpdateSubscription initial scenario options")
	}
	if len(env.Res.Message) == 0 {
		t.Errorf("Failed TestUpdateSubscription initial scenario message")
	}
	if env.User.Action != CMD_SUBSCRIBE {
		t.Errorf("Failed TestUpdateSubscription initial scenario state")
	}

	env = UpdateSubscription(env)
	if len(env.Res.Message) == 0 {
		t.Errorf("Failed TestUpdateSubscription error scenario message")
	}

	env.User.Action = CMD_SUBSCRIBE
	env.Msg.Message = "Discipleship Journal Bible Reading Plan"
	env = UpdateSubscription(env)
	config := utils.DeserializeUserConfig(env.User.Config)
	log.Printf("Subscriptions: %s", config.Subscriptions)
	if len(env.Res.Affordances.Options) < 1 {
		t.Errorf("Failed TestUpdateSubscription fulfillment scenario options")
	}
	if len(env.Res.Message) == 0 {
		t.Errorf("Failed TestUpdateSubscription fulfillment scenario message")
	}
	if config.Subscriptions != DJBRP {
		t.Errorf("Failed TestUpdateSubscription fulfillment scenario state")
	}
}
