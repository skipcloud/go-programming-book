package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strings"
)

/*
	Extend xmlselect so that elements may be selected not just by name, but by
	their attributes too, in the manner of CSS, so that, for instance, an
	element like <div id="page" class="wide"> could be selected by a matching
	id or class as well as its name.

	Note from Skip: attributes are provided on the command line with this format:
					attribute=value, e.g. class=container
*/

type userInput struct {
	userElements []*element
}

type element struct {
	name  string
	attrs []string
}

func main() {
	dec := xml.NewDecoder(os.Stdin)
	input := parseInput()
	var stack []xml.StartElement // stack of element names
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "xmlselect: %v\n", err)
			os.Exit(1)
		}
		switch tok := tok.(type) {
		case xml.StartElement:
			stack = append(stack, tok) // push Element
		case xml.EndElement:
			stack = stack[:len(stack)-1] // pop
		case xml.CharData:
			if containsAll(stack, input.userElements) {
				fmt.Printf("%s: %s\n", joinStack(stack, input.userElements), tok)
			}
		}
	}
}

func parseInput() userInput {
	input := userInput{}
	var el *element

	for _, v := range os.Args[1:] {
		if strings.Contains(v, "=") {
			el.attrs = append(el.attrs, v)
			continue
		}
		el = &element{
			name: v,
		}
		input.userElements = append(input.userElements, el)
	}
	return input
}

// containsAll reports whether x contains the elements of y, in order.
func containsAll(x []xml.StartElement, y []*element) bool {
	for len(y) <= len(x) {
		if len(y) == 0 {
			return true
		}
		// if the element matches, send the xml element and the input element
		// to containsAttrs to check if the StartElement contains all the attributes
		// in the element y[0]
		if x[0].Name.Local == y[0].name && containsAttrs(x[0], y[0]) {
			y = y[1:]
		}
		x = x[1:]
	}
	return false
}

// containsAttrs takes an xml.StartElement and a user element
// and checks all of the attributes in the user element against
// the attributes of the StartElement. The function returns true
// if the StartElement contains all of the attributes in the user
// element
func containsAttrs(el xml.StartElement, input *element) bool {
outer:
	for _, v := range input.attrs {
		// iterate over all the attributes of the element
		for _, a := range el.Attr {
			// check the input equals the attribute
			if v == (a.Name.Local + "=" + a.Value) {
				// if it does then continue, we're good. Jump to
				// the label "outer" to check the next user attribute
				continue outer
			}
		}
		return false
	}
	return true
}

// joinStack takes a slice of xml.StartElements and a slice of user elements
// and creates a "stack trace" inserting attibutes if the StartElement contains
// all of the specified attributes, e.g. div <id=container class=big> img p
func joinStack(stack []xml.StartElement, input []*element) string {
	newStack := []string{}

	for _, el := range stack {
		newStack = append(newStack, el.Name.Local)
		for _, v := range input {
			if el.Name.Local == v.name && containsAttrs(el, v) {
				newStack = append(newStack, fmt.Sprintf("<%s>", strings.Join(v.attrs, " ")))
			}

		}
	}
	return strings.Join(newStack, " ")
}
