package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

/*
	Write a variadic function ElementsByTagName that, given an HTML node
	tree and zero or more names, returns all the elements that match one
	of those names. Here are two example calls:

		func ElementsByTagName(doc *html.Node, name ...string) []*html.Node

		images := ElementsByTagName(doc, "img")
		headings := ElementsByTagName(doc, "h1", "h2", "h3", "h4")

*/

func main() {
	const url = "http://www.google.com"
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	r := ElementsByTagName(doc, "div", "a")
	fmt.Printf("There are %d div and anchor elements at %s\n", len(r), url)

	r = ElementsByTagName(doc, "h1", "h2", "h3", "h4")
	fmt.Printf("There are %d header elements at %s\n", len(r), url)

	r = ElementsByTagName(doc, "html")
	fmt.Printf("There are %d html elements at %s\n", len(r), url)
}

func ElementsByTagName(doc *html.Node, names ...string) []*html.Node {
	var results []*html.Node

	if len(names) == 0 {
		return results
	}
	if doc.Type == html.ElementNode {
		for _, name := range names {
			if strings.ToLower(doc.Data) == strings.ToLower(name) {
				results = append(results, doc)
			}
		}
	}
	for c := doc.FirstChild; c != nil; c = c.NextSibling {
		results = append(results, ElementsByTagName(c, names...)...)
	}
	return results
}
