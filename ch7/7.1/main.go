package main

import (
	"bufio"
	"bytes"
	"fmt"
)

/*
	Using the ideas from ByteCounter, implement counters for words and for lines.
	You will find bufio.ScanWords useful.
*/

const text = ` Lorem ipsum dolor sit amet consetetur sadipscing elitr, sed
diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed
diam voluptua. At vero eos et accusam et justo duo dolores et ea rebum. Stet
clita kasd gubergren, no sea takimata sanctus est Lorem ipsum dolor sit amet.

Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod
tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At
vero eos et accusam et justo duo dolores et ea rebum. Stet clita kasd
gubergren, no sea takimata sanctus est Lorem ipsum dolor sit amet.`

type WordCounter int
type LineCounter int

func (w *WordCounter) Write(p []byte) (int, error) {
	scanner := bufio.NewScanner(bytes.NewReader(p))
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		*w += 1
	}
	return len(p), nil
}

func (l *LineCounter) Write(p []byte) (int, error) {
	scanner := bufio.NewScanner(bytes.NewReader(p))
	for scanner.Scan() {
		*l += 1
	}
	return len(p), nil
}

func main() {
	var w WordCounter
	var l LineCounter

	fmt.Fprintf(&w, text)
	fmt.Fprintf(&l, text)
	fmt.Println(w)
	fmt.Println(l)
}
