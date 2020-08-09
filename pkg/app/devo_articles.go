package app

import (
	"log"

	"github.com/julwrites/BotPlatform/pkg/def"
	"github.com/julwrites/ScriptureBot/pkg/utils"
	"golang.org/x/net/html"
)

func GetDesiringGodHtml() *html.Node {
	query := "http://rss.desiringgod.org"

	return utils.QueryHtml(query)
}

func GetDesiringGodArticles() []def.Option {
	var options []def.Option

	doc := GetDesiringGodHtml()

	itemNodes := utils.FilterTree(doc, func(node *html.Node) bool {
		return node.Data == "item"
	})

	for _, node := range itemNodes {
		log.Printf("DesiringGod RSS itemtitle nodes %v", node)

		titleNode := utils.FindNode(node, func(node *html.Node) bool {
			return node.Data == "title"
		})
		linkNode := utils.FindNode(node, func(node *html.Node) bool {
			return node.Data == "link"
		})

		label := titleNode.FirstChild.Data
		link := linkNode.FirstChild.Data

		log.Printf("Label: %s, Link: %s", label, link)

		if len(label) > 0 && len(link) > 0 {
			options = append(options, def.Option{Text: label, Link: link})
		}
	}

	return options
}
