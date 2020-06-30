package main

import (
	"fmt"

	"github.com/skipcloud/go-programming-book/ch12/12.3/sexpr"
)

/*
	Implement the missing cases of the encode function. Encode booleans as t and nil,
	floating-point numbers using Go's notation, and complex numbers like 1+2i as
	#C(1.0 2.0). Interfaces can be encoded as a pair of a type name and a value, for
	instance ("[]int" (1 2 3)), but beware that this notation is ambiguous: the
	reflect.Type.String method may return the same string for different types.
*/

func main() {
	m := MyStruct{
		b: true,
		c: 1 + 2i,
		d: 1.2345,
		e: []int{1, 2, 3},
	}
	b, _ := sexpr.Marshal(m)
	fmt.Println(string(b))
}

type MyStruct struct {
	b bool
	c complex64
	d float64
	e Skip
}

type Skip interface {
	/* TODO: add methods */
}
