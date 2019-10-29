package main

import (
	"bytes"
	"fmt"
)

/*
 * Write a non-recursive version of comma, using bytes.Buffer instead
 * of string concatenation.
 */

func main() {
	s := "0000001234567890"
	fmt.Printf("%s\n", comma(s))
}

func comma(s string) string {
	var buf bytes.Buffer
	var newS string

	n := len(s)
	for {
		newS = buf.String()
		buf.Reset()
		if n <= 3 {
			buf.WriteString(s[:n] + newS)
			break
		}
		buf.WriteString("," + s[n-3:n] + newS)
		n -= 3
	}
	n = buf.Len()
	return buf.String()
}
