package app

import (
	"bytes"
	"log"
	"net/http"
	"strings"

	"github.com/julwrites/BotPlatform/pkg/def"
	"github.com/julwrites/ScriptureBot/pkg/utils"
	"golang.org/x/net/html"
)

// getTextFromNode extracts the text content from an HTML node.
// fetchHTMLPage fetches an HTML page from the given URL and parses it into an *html.Node.
func fetchHTMLPage(url string) *html.Node {
	res, err := http.Get(url)
	if err != nil {
		log.Printf("Error fetching HTML from %s: %v", url, err)
		return nil
	}
	defer res.Body.Close()

	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(res.Body)
	if err != nil {
		log.Printf("Error reading response body from %s: %v", url, err)
		return nil
	}
	// fmt.Printf(buf.String()) // Removed debug print

	doc, err := html.Parse(buf)
	if err != nil {
		log.Printf("Error parsing HTML from %s: %v", url, err)
		return nil
	}
	return doc
}

// parseArticlesFromHTML parses an HTML document and extracts articles based on common patterns.
func parseArticlesFromHTML(doc *html.Node) []def.Option {
	var options []def.Option

	itemNodes := utils.FilterTree(doc, func(node *html.Node) bool {
		return node.Data == "item"
	})

	for _, node := range itemNodes {
		titleNode := utils.FindNode(node, func(node *html.Node) bool {
			return node.Data == "title"
		})
		linkNode := utils.FindNode(node, func(node *html.Node) bool {
			return node.Data == "link"
		})

		label := titleNode.FirstChild.Data
		link := linkNode.Data
		if linkNode.FirstChild != nil {
			link = linkNode.FirstChild.Data
		} else if linkNode.NextSibling != nil {
			link = linkNode.NextSibling.Data
		}
		link = strings.TrimSpace(link)

		log.Printf("Label: %s, Link: %s", label, link)

		if len(label) > 0 && len(link) > 0 {
			options = append(options, def.Option{Text: label, Link: link})
		}
	}

	return options
}

func GetDesiringGodArticles() []def.Option {
	doc := fetchHTMLPage("http://rss.desiringgod.org")
	if doc == nil {
		return []def.Option{}
	}
	return parseArticlesFromHTML(doc)
}

func GetUtmostForHisHighestArticles() []def.Option {
	var options []def.Option
	options = append(options, def.Option{Text: "Modern Classic", Link: "http://utmost.org/modern-classic/today"})
	options = append(options, def.Option{Text: "Classic", Link: "http://utmost.org/classic/today"})
	options = append(options, def.Option{Text: "Updated", Link: "http://utmost.org/updated/today"})
	return options
}
