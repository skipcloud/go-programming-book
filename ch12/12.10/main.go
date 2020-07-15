package main

import (
	"fmt"
	"log"

	"github.com/skipcloud/go-programming-book/ch12/12.10/sexpr"
	encodeSexpr "github.com/skipcloud/go-programming-book/ch12/12.3/sexpr"
)

/*
	Extend sexpr.Unmarshal to handle the booleans, floating-point numbers, and interfaces
	encoded by your solution to exercise 12.3. (Hint: to decode interfaces, you will need
	a mapping from the name of each supported type to its reflect.Type)

	edit: didn't implement interface decoding
*/

type skip struct {
	Name     string
	Age      float64
	Employee bool
}

func main() {
	s := skip{
		Name:     "alan",
		Age:      1.23,
		Employee: true,
	}
	b, err := encodeSexpr.Marshal(s)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("encoded data: %s\n", string(b))
	var s2 skip
	err = sexpr.Unmarshal(b, &s2)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("output: %#v\n", s2)
}
