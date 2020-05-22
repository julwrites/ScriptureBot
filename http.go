// Brief: HTTP Wrapper
// Primary responsibility: Logging and abstraction around HTTP calls

package scripturebot

import (
	"io"
	"log"
	"net/http"

	"golang.org/x/net/html"
)

func Get(url string) *html.Node {
	res, getErr := http.Get(url)
	if getErr != nil {
		log.Fatalf("Error in GET call: %v", getErr)
		return nil
	}

	doc, parseErr := html.Parse(res.Body)
	if parseErr != nil {
		log.Fatalf("Error parsing html: %v", parseErr)
	}

	return doc
}

func Post(url string, body io.Reader) {
	_, postErr := http.Post(url, "markdown", body)
	if postErr != nil {
		log.Fatalf("Error in POST call: %v", postErr)
	}
}
