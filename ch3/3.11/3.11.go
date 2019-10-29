package main

import (
	"fmt"
	"strings"
)

/*
 * Enhance comma so that it deals correctly with floating-point
 * numbers and an optional sign.
 */

func main() {
	fmt.Println(comma("-1234567.1234567"))
}

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	if s[0] == '-' || s[0] == '+' {
		return string(s[0]) + comma(s[1:])
	}
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		return comma(s[:dot]) + s[dot:]
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}
