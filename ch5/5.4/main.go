package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

/*
	Extend the visit function so that is extracts other kinds of links from the
	document, such as images, scripts, and style sheets.
*/
func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	l := visit(nil, doc)
	fmt.Printf("count: %d\n", len(l))
}

func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && elementHasLink(n) {
		for _, a := range n.Attr {
			if a.Key == "href" || a.Key == "src" {
				links = append(links, a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}

func elementHasLink(n *html.Node) bool {
	return n.Data == "a" || n.Data == "img" || n.Data == "script" || n.Data == "link"
}

/*
//!+html
package html

type Node struct {
	Type                    NodeType
	Data                    string
	Attr                    []Attribute
	FirstChild, NextSibling *Node
}

type NodeType int32

const (
	ErrorNode NodeType = iota
	TextNode
	DocumentNode
	ElementNode
	CommentNode
	DoctypeNode
)

type Attribute struct {
	Key, Val string
}

func Parse(r io.Reader) (*Node, error)
//!-html
*/
