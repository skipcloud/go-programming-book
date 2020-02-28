package main

import "fmt"

/*
	Write a variadic version of strings.Join
*/
func main() {

	fmt.Printf("'%s'\n", join("hi", "there", "this", "is", "a", "variadic", "function", " "))
}

func join(args ...string) string {
	var str string
	if len(args) <= 1 {
		return str
	}

	sep := args[len(args)-1]
	args = args[:len(args)-1]
	for i, a := range args {
		str += a
		if i != len(args)-1 {
			str += sep
		}
	}
	return str
}
