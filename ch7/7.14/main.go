package main

import (
	"fmt"
	"os"

	"github.com/skipcloud/go-programming-book/ch7/7.14/eval"
)

/*
	Define a new concrete type that satisfies the Expr interface and provides
	a new operation such as computing the minimum value of its operands. Since
	the Parse function does not create instances of this new type, to use it
	you will need to constuct a syntax tree directly (or extend the Parser)
*/

func main() {
	var input string
	if len(os.Args) > 1 {
		input = os.Args[1]
	}
	if input == "" {
		input = "min(1,2,3,-10)"
	}

	e, _ := eval.Parse(input)
	fmt.Println(e.Eval(eval.Env{}))
}
