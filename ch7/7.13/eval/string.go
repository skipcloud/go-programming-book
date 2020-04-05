package eval

import (
	"bytes"
	"fmt"
	"strings"
)

// some global variables to help with printing
var padding = ""

// binary's String method is the heavy lifter in all of this
func (b binary) String() string {
	// The stand padding to add
	var paddingToAdd = "|   "
	left, right := padding+"├── ", padding+"└── "

	if _, ok := b.x.(binary); !ok {
		// if the left node of this tree is not a binary expression
		// then we don't need the downward bar part of the string,
		// we can use the 'right' string.
		left = right
		paddingToAdd = "    "
	}
	// Add string to padding
	padding += paddingToAdd
	//fetch the string of the left node all the way down
	left += b.x.String()
	// return the padding to its previous state
	padding = strings.TrimSuffix(padding, paddingToAdd)

	// right nodes won't lead to yet another node so the
	// padding can just be spaces
	paddingToAdd = "    "
	padding += paddingToAdd
	// fetch the string of the right node all the way down
	right += b.y.String()
	// return the padding to what it was previously
	padding = strings.TrimSuffix(padding, paddingToAdd)

	// put it all together
	str := fmt.Sprintf("%s\n%s\n%s", string(b.op), left, right)
	return str
}

func (u unary) String() string {
	return fmt.Sprintf("%c%s", u.op, u.x)
}

func (c call) String() string {
	// All this does is recreate the method with it's arguments
	// e.g. pow(10)
	var buf bytes.Buffer
	buf.WriteString(c.fn + "(")
	for i, v := range c.args {
		buf.WriteString(v.String())
		if i < len(c.args)-1 {
			buf.WriteString(", ")
		}
	}
	buf.WriteString(")")
	return buf.String()
}

func (v Var) String() string {
	return string(v)
}

func (l literal) String() string {
	return fmt.Sprintf("%g", l)
}
