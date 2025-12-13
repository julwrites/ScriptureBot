package bot

import (
	"strings"
	"testing"

	"github.com/julwrites/BotPlatform/pkg/def"
	"github.com/julwrites/ScriptureBot/pkg/app"
	"github.com/julwrites/ScriptureBot/pkg/utils"
)

func TestRunCommands(t *testing.T) {
	defer app.UnsetEnv("BIBLE_API_URL")()
	defer app.UnsetEnv("BIBLE_API_KEY")()
	app.ResetAPIConfigCache()

	var env def.SessionData
	var conf utils.UserConfig
	conf.Version = "NIV"
	env = utils.SetUserConfig(env, utils.SerializeUserConfig(conf))
	env.Msg.Message = "psalm 1"

	env = RunCommands(env, &app.MockBot{})

	if !strings.Contains(env.Res.Message, "wicked") && !strings.Contains(env.Res.Message, "Blessed") {
		t.Errorf("Failed TestRunCommands Passage command. Got: %s", env.Res.Message)
	}
}

func TestUserCheck(t *testing.T) {
	var user utils.User

	user.Firstname = "User"
	user.Lastname = ""
	user.Username = "username"
	user.Id = "id"
	user.Type = "Individual"
	user.Action = ""
	user.Config = ""

	cache := user

	if user != cache {
		t.Errorf("Failed TestUserCheck identical user")
	}

	user.Lastname = "name"

	if user == cache {
		t.Errorf("Failed TestUserCheck changed user")
	}
}
