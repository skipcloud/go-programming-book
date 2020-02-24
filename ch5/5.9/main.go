package main

import (
	"fmt"
	"os"
	"strings"
)

/*
	Write a function expand(s string, f func(string) string) string that replaces
	each substring "$foo" within s by the text returned by f("foo").
*/

func main() {
	if len(os.Args) == 1 {
		fmt.Fprintf(os.Stderr, "missing arguments")
		os.Exit(1)
	}
	// let's assume a quoted string is passed in
	input := os.Args[1]

	// I'm too lazy to think of more mappings
	replacements := map[string]string{
		"foo": "bar",
	}
	fmt.Printf("%s\n", expand(input, func(s string) string {
		if val, ok := replacements[s]; ok {
			return val
		}
		return ""
	}))

}

func expand(s string, f func(string) string) string {
	words := strings.Split(s, " ")
	output := make([]string, len(words))
	for i, word := range words {
		if strings.HasPrefix(word, "$") {
			output[i] = f(strings.TrimPrefix(word, "$"))
			continue
		}
		output[i] = word
	}
	return strings.Join(output, " ")
}
