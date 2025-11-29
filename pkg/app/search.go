package app

import (
	"fmt"
	"log"
	"strings"

	"github.com/julwrites/BotPlatform/pkg/def"
	"github.com/julwrites/ScriptureBot/pkg/secrets"
	"github.com/julwrites/ScriptureBot/pkg/utils"
)

func GetBibleSearch(env def.SessionData) def.SessionData {
	if len(env.Msg.Message) > 0 {
		config := utils.DeserializeUserConfig(env.User.Config)

		// Parse message into words?
		// The API expects a list of words.
		// If the user types "Grace of God", we can just pass "Grace", "of", "God" or the whole string as one item?
		// The API spec says `words` is `[]string`.
		// Usually "Word Search" implies searching for the phrase or keywords.
		// Let's split by space for now, or just pass the whole phrase if the API supports phrase search.
		// The API example shows "Grace".
		// If I assume it's a list of keywords to AND/OR together.
		// Let's treat the entire message as the search query. The API takes a list, so maybe it allows multiple separate search terms.
		// But usually a user types a sentence or phrase.
		// I will split by spaces to be safe and provide them as individual words.
		words := strings.Fields(env.Msg.Message)

		req := QueryRequest{
			Query: QueryObject{
				Words: words,
			},
			Context: QueryContext{
				User: UserContext{
					Version: config.Version,
				},
			},
		}

		var resp WordSearchResponse
		projectID, _ := secrets.Get("GCLOUD_PROJECT_ID")
		err := SubmitQuery(req, &resp, projectID)
		if err != nil {
			log.Printf("Error searching bible: %v", err)
			env.Res.Message = "Sorry, I encountered an error while searching."
			return env
		}

		if len(resp) > 0 {
			var sb strings.Builder
			sb.WriteString(fmt.Sprintf("Found %d results for '%s':\n", len(resp), env.Msg.Message))
			for _, res := range resp {
				// Format: - Verse (URL)
				// Markdown link: [Verse](URL)
				sb.WriteString(fmt.Sprintf("- [%s](%s)\n", res.Verse, res.URL))
			}
			env.Res.Message = sb.String()
		} else {
			env.Res.Message = fmt.Sprintf("No results found for '%s'.", env.Msg.Message)
		}
	}
	return env
}
