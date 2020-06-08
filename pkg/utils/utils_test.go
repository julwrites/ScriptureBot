package utils

import (
	"testing"

	"golang.org/x/net/html"
)

// HTML tests
type HtmlTestStruct struct {
	Root *html.Node
	List []*html.Node
}

func GetTestData() HtmlTestStruct {
	var attr html.Attribute
	attr.Key = "class"
	attr.Val = "positive"

	var test1 html.Node
	test1.Type = html.ElementNode
	test1.Data = "positive"
	test1.Attr = append(test1.Attr, attr)

	var test2 html.Node
	test2.Type = html.CommentNode
	test2.Data = "positive"

	var test3 html.Node
	test3.Type = html.ElementNode
	test3.Data = "positive"
	test3.Attr = append(test3.Attr, attr)

	test2.AppendChild(&test3)

	var root html.Node
	root.AppendChild(&test1)
	root.AppendChild(&test2)

	var list []*html.Node
	list = append(list, &test1)
	list = append(list, &test2)
	list = append(list, &test3)

	return HtmlTestStruct{Root: &root, List: list}
}

func TestGetHtml(t *testing.T) {
	pos := GetHtml("https://www.google.com")

	if pos == nil {
		t.Errorf("Failed GetHtml positive scenario")
	}

	neg := GetHtml("https://www.google.thiscannotbe")

	if neg != nil {
		t.Errorf("Failed GetHtml negative scenario")
	}
}

func TestFindNode(t *testing.T) {
	root := GetTestData().Root

	pos := FindNode(root, func(node *html.Node) bool { return node.Data == "positive" })

	if pos == nil {
		t.Errorf("Failed FindNode positive scenario")
	}

	neg := FindNode(root, func(node *html.Node) bool { return node.Data == "negative" })

	if neg != nil {
		t.Errorf("Failed FindNode negative scenario")
	}
}

func TestFilterTree(t *testing.T) {
	root := GetTestData().Root

	pos := FilterTree(root, func(node *html.Node) bool { return node.Data == "positive" })

	if len(pos) != 3 {
		t.Errorf("Failed TestFilterTree positive scenario")
	}

	neg := FilterTree(root, func(node *html.Node) bool { return node.Data == "negative" })

	if len(neg) != 0 {
		t.Errorf("Failed TestFilterTree negative scenario")
	}
}

func TestFilterNodeList(t *testing.T) {
	list := GetTestData().List

	pos := FilterNodeList(list, func(node *html.Node) bool { return node.Data == "positive" })

	if len(pos) != 3 {
		t.Errorf("Failed FilterNodeList positive scenario")
	}

	neg := FilterNodeList(list, func(node *html.Node) bool { return node.Data == "negative" })

	if len(neg) != 0 {
		t.Errorf("Failed FilterNodeList negative scenario")
	}
}

func TestFilterChildren(t *testing.T) {
	root := GetTestData().Root

	pos := FilterChildren(root, func(node *html.Node) bool { return node.Data == "positive" })

	if len(pos) != 2 {
		t.Errorf("Failed FilterChildren positive scenario")
	}

	neg := FilterChildren(root, func(node *html.Node) bool { return node.Data == "negative" })

	if len(neg) != 0 {
		t.Errorf("Failed FilterChildren negative scenario")
	}
}

func TestMapTreeToString(t *testing.T) {
	root := GetTestData().Root

	output := MapTreeToString(root, func(node *html.Node) string { return node.Data })
	if len(output) != 4 {
		t.Errorf("Failed MapTree")
	}
}

func TestMapNodeListToString(t *testing.T) {
	list := GetTestData().List

	output := MapNodeListToString(list, func(node *html.Node) string { return node.Data })

	if len(output) != 3 {
		t.Errorf("Failed MapNodeList positive scenario")
	}
}

func TestFindByClass(t *testing.T) {
	root := GetTestData().Root

	pos, perror := FindByClass(root, "positive")

	if pos == nil || perror != nil {
		t.Errorf("Failed FindByClass positive scenario")
	}

	neg, nerror := FindByClass(root, "negative")

	if neg != nil || nerror == nil {
		t.Errorf("Failed FindByClass negative scenario")
	}
}

func TestFilterByNodeType(t *testing.T) {
	root := GetTestData().Root

	output := FilterByNodeType(root, html.ElementNode)

	if len(output) != 2 {
		t.Errorf("Failed FilterByNodeType")
	}
}

// URL Query tests
func TestQueryBiblePassage(t *testing.T) {
	doc := QueryBiblePassage("gen 1", "NIV")

	if doc == nil {
		t.Errorf("Could not retrieve bible passage")
	}
}

func TestQueryBibleLexicon(t *testing.T) {
	doc := QueryBibleLexicon("beginning", "NIV")

	if doc == nil {
		t.Errorf("Could not retrieve bible passage")
	}
}
