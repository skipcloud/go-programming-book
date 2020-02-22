package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Fprintf(os.Stderr, "missing url\n")
		os.Exit(1)
	}
	url := strings.Join(os.Args[1:], " ")
	w, i, err := CountWordsAndImages(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("words: %d - images: %d\n", w, i)
}

// CountWordsAndImages does an HTTP GET request for the HTML
// document url and returns the number of words and images in it.
func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}

	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("parse HTML: %s", err)
		return
	}

	words, images = countWordsAndImages(doc)
	return
}

func countWordsAndImages(n *html.Node) (words, images int) {
	switch n.Type {
	case html.ElementNode:
		if n.Data == "img" {
			images += 1
		}
	case html.TextNode:
		in := bufio.NewScanner(strings.NewReader(n.Data))
		in.Split(bufio.ScanWords)

		for in.Scan() {
			words += 1
		}
	}

	var w, i int
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		w, i = countWordsAndImages(c)
		words += w
		images += i
	}

	return
}
