package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

/*
	Write a function to print the content all text nodes in an HTML document
	tree. Do not descend into <script> or <style> elements, since their contents
	are not visible in a web browser.
*/

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error parsing input\n")
		os.Exit(1)
	}
	printTextNodes(doc)
}

func printTextNodes(n *html.Node) {
	if n.Type == html.TextNode {
		fmt.Printf("%s\n", n.Data)
	}
	for c := n.FirstChild; c != nil && (c.Data != "script" || c.Data != "style"); c = c.NextSibling {
		printTextNodes(c)
	}
}
