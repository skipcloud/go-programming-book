package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"os"
)

/*
	Using the token-based decoder API, write a program that will read an
	arbitrary XML document and construct a tree of generic nodes that
	represent it. Nodes are of two kinds: CharData nodes represent text
	strings, and Element nodes represent named elements and their
	attributes. Each element node has a slice of child nodes.

	You may find the following declarations helpful:

	import "encoding/xml"

	type Node interface{} // CharData or *Element

	type CharData string

	type Element struct {
		Type     xml.Name
		Attr     []xml.Attr
		Children []Node
	}
*/

type Node interface{} // CharData or *Element

type CharData string

type Element struct {
	Type     xml.Name
	Attr     []xml.Attr
	Children []Node
}

type Tree []Node

func main() {
	// get a new xml decoder
	dec := xml.NewDecoder(os.Stdin)

	// the main tree
	var tree Tree

	// top level loop
	for {

		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
		}
		// switch on the dynamic type
		switch tok := tok.(type) {
		// if it is a StartElement then call parseElement
		case xml.StartElement:
			el := newElement(tok)
			parseElement(tok, el, dec)
			// after parsing add the node (with all of its children)
			// to the tree
			tree = append(tree, el)
		// if it is CharData then just add it to the tree
		// as a toplevel node
		case xml.CharData:
			tree = append(tree, CharData(tok))
		}
	}
	// Because I defined a String() method on Tree just
	// pass it to Printf and let String() do all the work.
	fmt.Printf("%s", tree)

}

// newElement takes a xml.StartElement and returns an *Element
// with the data prefilled
func newElement(el xml.StartElement) *Element {
	return &Element{
		Type: el.Name,
		Attr: el.Attr,
	}
}

// parseElement takes the token to be parsed, the parentNode, and the xml decoder
// and recursively adds new nodes to the parentNode's children slice
func parseElement(token xml.Token, parentNode *Element, dec *xml.Decoder) {
	for {
		// find the next token
		tok := nextToken(dec)
		// if it's nil we have probably hit EOF so return
		if tok == nil {
			return
		}
		// switch on dynamic type
		switch tok := tok.(type) {
		case xml.StartElement:
			// if StartElement then start a new node and recursively
			// find all child nodes, then add the new node to the parentNode
			e := newElement(tok)
			parseElement(tok, e, dec)
			parentNode.Children = append(parentNode.Children, e)
		case xml.EndElement:
			// if EndElement then return, it's the end of the current node
			return
		case xml.CharData:
			// if CharData add it as a child to the parent node
			parentNode.Children = append(parentNode.Children, CharData(tok))
		}
	}
}

// nextToken is a helper function to reduce duplication of code when
// getting the next token from the xml decoder
func nextToken(dec *xml.Decoder) xml.Token {
	tok, err := dec.Token()
	if err == io.EOF {
		return nil
	} else if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	return tok
}

func (e *Element) String() string {
	buf := bytes.NewBuffer([]byte{})
	buf.WriteString(e.Type.Local + "\n")
	for _, v := range e.Children {
		switch v := v.(type) {
		case *Element:
			buf.WriteString("  " + v.String())
		case CharData:
			buf.WriteString(string(v))
		}

	}
	return buf.String()
}

func (t Tree) String() string {
	buf := bytes.NewBuffer([]byte{})
	for _, v := range t {
		switch v := v.(type) {
		case *Element:
			buf.WriteString(v.String())
		case CharData:
			buf.WriteString(string(v))
		}

	}
	return buf.String()
}
