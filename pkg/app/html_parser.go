package app

import (
	"fmt"
	stdhtml "html"
	"strings"

	"golang.org/x/net/html"
)

// ParseToTelegramHTML converts generic HTML to Telegram-supported HTML.
// It converts block elements like <p> to newlines, handles lists, and preserves
// inline formatting like <b>, <i>, <a> while stripping unsupported tags.
func ParseToTelegramHTML(htmlStr string) string {
	doc, err := html.Parse(strings.NewReader(htmlStr))
	if err != nil {
		// Fallback to original string if parsing fails
		return htmlStr
	}

	return strings.TrimSpace(parseNodesForTelegram(doc))
}

func parseNodesForTelegram(node *html.Node) string {
	var parts []string

	for child := node.FirstChild; child != nil; child = child.NextSibling {
		switch tag := child.Data; tag {
		case "b", "strong":
			parts = append(parts, fmt.Sprintf("<b>%s</b>", parseNodesForTelegram(child)))
		case "i", "em":
			parts = append(parts, fmt.Sprintf("<i>%s</i>", parseNodesForTelegram(child)))
		case "u", "ins":
			parts = append(parts, fmt.Sprintf("<u>%s</u>", parseNodesForTelegram(child)))
		case "s", "strike", "del":
			parts = append(parts, fmt.Sprintf("<s>%s</s>", parseNodesForTelegram(child)))
		case "code":
			parts = append(parts, fmt.Sprintf("<code>%s</code>", parseNodesForTelegram(child)))
		case "pre":
			parts = append(parts, fmt.Sprintf("<pre>%s</pre>", parseNodesForTelegram(child)))
		case "a":
			href := ""
			for _, attr := range child.Attr {
				if attr.Key == "href" {
					href = attr.Val
					break
				}
			}
			if href != "" {
				parts = append(parts, fmt.Sprintf(`<a href="%s">%s</a>`, href, parseNodesForTelegram(child)))
			} else {
				parts = append(parts, parseNodesForTelegram(child))
			}
		case "p":
			parts = append(parts, parseNodesForTelegram(child))
			parts = append(parts, "\n\n")
		case "br":
			parts = append(parts, "\n")
		case "ul", "ol":
			parts = append(parts, parseNodesForTelegram(child))
		case "li":
			parts = append(parts, fmt.Sprintf("• %s\n", strings.TrimSpace(parseNodesForTelegram(child))))
		case "h1", "h2", "h3", "h4", "h5", "h6":
			parts = append(parts, fmt.Sprintf("<b>%s</b>\n", strings.TrimSpace(parseNodesForTelegram(child))))
		default:
			if child.Type == html.TextNode {
				parts = append(parts, stdhtml.EscapeString(child.Data))
			} else if child.Type == html.ElementNode {
				// Recurse for unknown elements (like div, span) to preserve content
				parts = append(parts, parseNodesForTelegram(child))
			}
		}
	}

	return strings.Join(parts, "")
}
