// Outline prints the outline of an HTML document tree.
package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

/*
	Develop startElement and endElement into a general HTML pretty-printer.
	Print comment nodes, text nodes, and the attributes of each element
	(<a href="...">). Use short forms like <img /> instead of <img></img> when
	an element has no children. Write a test to ensure that the output can
	be parsed successfully. (See chapter 11.)
*/
func main() {
	for _, url := range os.Args[1:] {
		outline(url)
	}
}

func outline(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return err
	}

	forEachNode(doc, startElement, endElement)

	return nil
}

// forEachNode calls the functions pre(x) and post(x) for each node
// x in the tree rooted at n. Both functions are optional.
// pre is called before the children are visited (preorder) and
// post is called after (postorder).
func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		post(n)
	}
}

var depth int

func startElement(n *html.Node) {
	switch n.Type {
	case html.ElementNode:
		switch n.Data {
		case "img", "br":
			fmt.Printf("%*s<%s", depth*2, "", n.Data)
			printAttributes(n)
			fmt.Printf("/")
		default:
			fmt.Printf("%*s<%s", depth*2, "", n.Data)
			printAttributes(n)
			depth++
		}
		fmt.Printf(">\n")
	case html.CommentNode:
		fmt.Printf("%*s<--%s-->\n", depth*2, "", n.Data)
	case html.TextNode:
		if n.Data != "\n" {
			fmt.Printf("%*s %s\n", depth*2, "", n.Data)
		}
	}
}

func endElement(n *html.Node) {
	switch n.Type {
	case html.ElementNode:
		switch n.Data {
		case "img", "br":
		default:
			depth--
			fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
		}
	}
}

func printAttributes(n *html.Node) {
	for _, a := range n.Attr {
		fmt.Printf(" %s=\"%s\"", a.Key, a.Val)
	}
}
