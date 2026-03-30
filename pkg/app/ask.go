package app

import (
	"fmt"
	stdhtml "html"
	"log"
	"strings"

	"github.com/julwrites/BotPlatform/pkg/def"
	"github.com/julwrites/ScriptureBot/pkg/utils"
)

func GetBibleAsk(env def.SessionData) def.SessionData {
	if !utils.IsAdmin(env) {
		return ProcessNaturalLanguage(env)
	}

	return GetBibleAskWithContext(env, nil)
}

func GetBibleAskWithContext(env def.SessionData, contextVerses []string) def.SessionData {
	if !utils.IsAdmin(env) {
		return env
	}
	if len(env.Msg.Message) > 0 {
		config := utils.DeserializeUserConfig(utils.GetUserConfig(env))

		req := QueryRequest{
			Query: QueryObject{
				Prompt: env.Msg.Message,
				Context: QueryContext{
					Verses: contextVerses,
				},
			},
			User: UserOptions{
				Version: config.Version,
			},
		}

		var resp PromptResponse
		err := SubmitQuery(req, &resp)
		if err != nil {
			log.Printf("Error asking bible: %v", err)
			env.Res.Message = "Sorry, I encountered an error processing your question."
			return env
		}

		var sb strings.Builder
		sb.WriteString(ParseToTelegramHTML(resp.Data.Text))

		if len(resp.Data.References) > 0 {
			sb.WriteString("\n\n<b>References:</b>")
			for _, ref := range resp.Data.References {
				sb.WriteString(fmt.Sprintf("\n• %s", stdhtml.EscapeString(ref.Verse)))
			}
		}

		env.Res.Message = sb.String()
		env.Res.ParseMode = def.TELEGRAM_PARSE_MODE_HTML
	}
	return env
}
