package main

import (
	"log"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

/*
	Modify forEachNode so that the pre and post functions return a boolean result
	indicating whether to continue the traversal. Use it to write a function
	ElementByID with the following signature that finds the first HTML element with
	the specified id attribute. The function should stop traversal as soon as a
	match is found.

		func ElementByID(n *html.Node, id string) *html.Node

	edit: I never finished this one because, I'll be honest, the task didn't make
	much sense to me. I understood ElementByID needed to find the node with a certain
	ID but why would the post function need to return a bool? Is ElementByID supposed
	to take the place of `outline` from a few exercises ago? Do I print out these nodes?
	I got myself tangled up and confused so I moved on. I might come back and try to
	figure out exactly what they want from me here, because right now I haven't a clue.
*/

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			log.Fatal()
		}
		defer resp.Body.Close()

		doc, err := html.Parse(resp.Body)
		if err != nil {
			log.Fatal()
		}

		id := "about"
		ElementByID(doc, id)
	}
}

func ElementByID(n *html.Node, id string) *html.Node {

}

func forEachNode(n *html.Node, pre, post func(n *html.Node) bool) {
	var found bool
	if pre != nil {
		found = pre(n)
		if found {
			return
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		found = post(n)
		if found {
			return
		}
	}
}

func checkElement(n *html.Node) bool {
	for _, a := range n.Attr {
		if a.Key == "id" && a.Val == id {
			return true
		}
	}

	return false
}
