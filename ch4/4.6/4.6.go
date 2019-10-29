package main

/*
 * Write an in-place function that squashes each run of adjacent
 * Unicode spaces (see unicode.IsSpace) in a UTF-8-encoded []byte
 * slice into a single ASCII space.
 */

import "fmt"
import "unicode"

func main() {
	b := []byte("here is     some text")

	fmt.Printf("normal text:   %s\n", string(b))
	fmt.Printf("squashed text: %s\n", string(squash(b)))
}

func squash(b []byte) []byte {
	i := 0
	foundSpace := false
	for _, v := range b {
		if unicode.IsSpace(rune(v)) {
			if foundSpace == false {
				foundSpace = true
				b[i] = v
				i++
			}
		} else {
			foundSpace = false
			b[i] = v
			i++
		}
	}
	return b[:i]
}
