package bot

import (
	"strings"
	"testing"

	"github.com/julwrites/BotPlatform/pkg/def"
	"github.com/julwrites/ScriptureBot/pkg/utils"
)

func TestRunCommands(t *testing.T) {
	var env def.SessionData
	var conf utils.UserConfig
	conf.Version = "NIV"
	env.User.Config = utils.SerializeUserConfig(conf)
	env.Msg.Message = "psalm 1"

	env = RunCommands(env)

	if !strings.Contains(env.Res.Message, "Not so the wicked!") {
		t.Errorf("Failed TestRunCommands Passage command")
	}
}

func TestUserCheck(t *testing.T) {
	var user def.UserData

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
