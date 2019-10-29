package main

import "fmt"

/*
 * Modify reverse to reverse the characters of a []byte slice that
 * represents a UTF-8-encoded string, in place. Can you do it without
 * allocating new memory?
 */

func main() {
	b := []byte("reverse me")
	reverse(b)
	fmt.Println(string(b))
}

func reverse(s []byte) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
