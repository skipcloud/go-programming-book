package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/skipcloud/go-programming-book/ch7/7.15/eval"
)

/*
	Write a program that reads a single expression from the standard input,
	prompts the user to provide values for any variables, then evaluates
	the expression in the resulting environment. Handle all errors gracefully.
*/

func main() {
	s := bufio.NewScanner(os.Stdin)
	if !s.Scan() {
		fmt.Fprintf(os.Stderr, "error: reading input\n")
		os.Exit(1)
	}
	input := s.Text()
	e, err := eval.Parse(input)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: parsing input\n")
		os.Exit(1)
	}

	vars := eval.FindVariables(e, []eval.Var{})
	env := eval.Env{}
	var val float64
	if len(vars) > 0 {
		for _, v := range vars {
			fmt.Printf("What is the value of %s: ", v)
			if !s.Scan() {
				fmt.Fprintf(os.Stderr, "error: reading input\n")
				os.Exit(1)
			}
			if val, err = strconv.ParseFloat(s.Text(), 64); err != nil {
				fmt.Fprintf(os.Stderr, "error: parsing input")
				os.Exit(1)
			}
			env[v] = val
		}
	}
	fmt.Printf("answer: %g\n", e.Eval(env))
}
