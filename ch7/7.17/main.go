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

func main() {
	dec := xml.NewDecoder(os.Stdin)
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
			if containsAll(stack, os.Args[1:]) {
				fmt.Printf("%s: %s\n", joinStack(stack, os.Args[1:]), tok)
			}
		}
	}
}

// containsAll reports whether x contains the elements of y, in order.
func containsAll(x []xml.StartElement, y []string) bool {
	for len(y) <= len(x) {
		elementsToSkip := 1
		if len(y) == 0 {
			return true
		}
		// if the element matches, send the element and the input elements
		// after our element to containsAttr. If y[0] is an element then
		// it's possible there are attributes after it in the input.
		// For example: []string{"div", "class=big", "id=post", "div"...}
		if x[0].Name.Local == y[0] {
			if n, ok := containsAttr(x[0], y[1:]); ok {
				// update the num of elements to skip. For example, if the input
				// has []string{element, attr, attr, element} and the xml Element
				// contains the two attrs from the input then we need to skip two
				// more elements in both our slices
				elementsToSkip += n
			}
			// element contains all the
			y = y[elementsToSkip:]
		}
		x = x[elementsToSkip:]
	}
	return false
}

// containsAttr iterates over contiguous attributes in the user input stopping
// at any element it comes across. It will return the number of attributes in
// the input that matched and true if the element contains an attribute
// matching the input otherwise it returns false.
func containsAttr(el xml.StartElement, input []string) (int, bool) {
	// iterate over user input as long as the elements contain "="
	// or in other words as long as the input is for attributes
	var numAttr int
outer:
	for i := 0; i < len(input) && strings.Contains(input[i], "="); i++ {
		// iterate over all the attributes of the element
		for _, a := range el.Attr {
			// check the input equals the attribute
			if input[i] == (a.Name.Local + "=" + a.Value) {
				// if it does then continue, we're good. Jump to
				// the label "outer" to check the next input string
				numAttr++
				continue outer
			}
		}
	}
	if numAttr == 0 {
		return numAttr, false
	}
	return numAttr, true
}

// joinStack takes a stack of xml elements and joins the name of the elements
// in the stack
func joinStack(stack []xml.StartElement, input []string) string {
	newStack := []string{}

	for _, el := range stack {
		newStack = append(newStack, el.Name.Local)
	}
	return strings.Join(newStack, " ")
}
