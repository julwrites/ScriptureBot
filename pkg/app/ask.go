package app

import (
	"fmt"
	"log"
	"strings"

	"github.com/julwrites/BotPlatform/pkg/def"
	"github.com/julwrites/ScriptureBot/pkg/utils"
)

func GetBibleAsk(env def.SessionData) def.SessionData {
	return GetBibleAskWithContext(env, nil)
}

func GetBibleAskWithContext(env def.SessionData, contextVerses []string) def.SessionData {
	if len(env.Msg.Message) > 0 {
		config := utils.DeserializeUserConfig(env.User.Config)

		req := QueryRequest{
			Query: QueryObject{
				Prompt: env.Msg.Message,
			},
			Context: QueryContext{
				User: UserContext{
					Version: config.Version,
				},
				Verses: contextVerses,
			},
		}

		var resp OQueryResponse
		err := SubmitQuery(req, &resp)
		if err != nil {
			log.Printf("Error asking bible: %v", err)
			env.Res.Message = "Sorry, I encountered an error processing your question."
			return env
		}

		var sb strings.Builder
		sb.WriteString(resp.Text)

		if len(resp.References) > 0 {
			sb.WriteString("\n\n*References:*")
			for _, ref := range resp.References {
				sb.WriteString(fmt.Sprintf("\n- [%s](%s)", ref.Verse, ref.URL))
			}
		}

		env.Res.Message = sb.String()
	}
	return env
}
