package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

/*
	Write a function to populate a mapping from elemant names — p, div, span,
	and so on — to the number of elements with that name in an HTML document tree.
*/

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error parsing input\n")
		os.Exit(1)
	}
	es := Elements{}
	countElements(doc, es)
	for e, count := range es {
		fmt.Printf("\telement: %s\tcount:%d\n", e, count)
	}
}

type Elements map[string]int

func countElements(node *html.Node, elements Elements) {
	if node.Type == html.ElementNode {
		elements[node.Data] += 1
	}
	if node.FirstChild != nil {
		countElements(node.FirstChild, elements)
	}
	if node.NextSibling != nil {
		countElements(node.NextSibling, elements)
	}
}
