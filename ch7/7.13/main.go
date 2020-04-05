package main

import (
	"fmt"
	"os"

	"github.com/skipcloud/go-programming-book/ch7/7.13/eval"
)

/*
	Add a String method to eval.Expr to pretty-print the syntax tree. Check
	that the results, when parsed again, yield an equivalent tree
*/

func main() {
	var input string
	if len(os.Args) > 1 {
		input = os.Args[1]
	}
	if input == "" {
		input = "(x + 1) * 10 - 100"
	}

	fmt.Printf("Syntax tree for %q\n\n", input)
	e, _ := eval.Parse(input)
	fmt.Println(e)
}
