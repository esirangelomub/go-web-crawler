package main

import (
	"flag"
	"fmt"
	"github.com/esirangelomub/go-web-crawler/db"
	"net/http"
	"net/url"
	"time"

	"golang.org/x/net/html"
)

type VisitedLink struct {
	Website     string    `bson:"website"`
	Link        string    `bson:"link"`
	VisitedDate time.Time `bson:"visited_date"`
}

var link string

func init() {
	flag.StringVar(&link, "url", "https://aprendagolang.com.br", "URL to init visits")
}

func main() {
	flag.Parse()

	done := make(chan bool)
	go visitLink(link)
	<-done
}

func visitLink(link string) {

	fmt.Printf("Visiting link: %s\n", link)

	resp, err := http.Get(link)
	if err != nil {
		panic(fmt.Sprintf("[error] get url: %d\n", err))
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf(fmt.Sprintf("[error] status not 200: %d\n", resp.StatusCode))
	}

	doc, err := html.Parse(resp.Body)
	if err != nil {
		panic(fmt.Sprintf("[error] parse doc: %d\n", err))
	}

	extractLinks(doc)
}

func extractLinks(node *html.Node) {

	if node.Type == html.ElementNode && node.Data == "a" {
		for _, attr := range node.Attr {
			if attr.Key != "href" {
				continue
			}

			link, err := url.Parse(attr.Val)
			if err != nil || link.Scheme == "" || link.Scheme == "mailto" {
				continue
			}

			if db.VisitedLink(link.String()) {
				fmt.Printf("link was visited: %s\n", link.String())
				continue
			}

			visitedLink := VisitedLink{
				Website:     link.Host,
				Link:        link.String(),
				VisitedDate: time.Now(),
			}

			err = db.Insert("links", visitedLink)
			if err != nil {
				continue
			}

			go visitLink(link.String())
		}
	}

	for c := node.FirstChild; c != nil; c = c.NextSibling {
		extractLinks(c)
	}
}
