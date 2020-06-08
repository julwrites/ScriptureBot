package app

import (
	"log"
	"strings"

	"github.com/julwrites/BotMultiplexer/pkg/def"
	"github.com/julwrites/ScriptureBot/pkg/utils"
	"golang.org/x/net/html"
)

// TODO: How to retrieve a Javascript triggered change in site?
func GetBibleWord(env def.SessionData) def.SessionData {
	if len(env.Msg.Message) > 0 {
		doc := utils.QueryBibleLexicon(env.Msg.Message, utils.DeserializeUserConfig(env.User.Config).Version)

		if doc != nil {
			filtNodes := utils.FilterTree(doc, func(node *html.Node) bool {
				for _, attr := range node.Attr {
					if attr.Key == "class" && attr.Val == "lexiconcData" {
						return true
					}
				}
				return false
			})

			log.Printf("Filtered to %d nodes", len(filtNodes))

			var blocks []string

			for _, node := range filtNodes {
				blocks = append(blocks, strings.Join(utils.MapTreeToString(node, func(node *html.Node) string {
					var parts []string

					switch tag := node.Data; tag {
					case "td":
						parts = append(parts, strings.Join(utils.MapTreeToString(node, func(node *html.Node) string {
							return node.Data
						}), " "))
						break
					default:
						break
					}

					output := strings.Join(parts, " ")

					return output
				}), " "))
			}

			log.Printf("Output: %s", strings.Join(blocks, "\n"))
		} else {
			log.Printf("Failed to get anything from bible word query")
		}
	}

	return env
}
