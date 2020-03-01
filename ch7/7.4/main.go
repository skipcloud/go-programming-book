package main

import (
	"fmt"
	"io"
	"os"

	"golang.org/x/net/html"
)

/*
	The strings.NewReader function in the io package returns a value that satisfies
	the io.Reader interface (and others) by reading from its argument, a string.
	Implement a simple version of NewReader yourself, and use it to make the HTML
	Parser (secton 5.2) take input from a string.
*/

const myHtml = `
<html>
	<head>
		<title>My website</title>
	</head>
	<body>
		<div>
			<h1>Hi</h1>
		</div>
	</body>
<html>
`

func main() {
	r := NewReader(myHtml)
	doc, err := html.Parse(r)
	if err != nil {
		fmt.Fprintf(os.Stderr, "outline: %v\n", err)
		os.Exit(1)
	}
	outline(nil, doc)
}

func outline(stack []string, n *html.Node) {
	if n.Type == html.ElementNode {
		stack = append(stack, n.Data) // push tag
		fmt.Println(stack)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		outline(stack, c)
	}
}

type myReader struct {
	bb []byte
	i  int
}

func (m *myReader) Read(p []byte) (n int, err error) {
	if m.i == len(m.bb) {
		return 0, io.EOF
	}
	n = copy(p, m.bb[m.i:])
	m.i += n
	return n, nil
}

func NewReader(s string) io.Reader {
	return &myReader{bb: []byte(s)}
}
