package main

import (
	"fmt"

	"github.com/skipcloud/go-programming-book/ch12/12.4/sexpr"
)

/*
	Modify encode to pretty-print the S-expression in the style shown above
*/

func main() {
	m := MyStruct{
		a: 10,
		b: true,
		f: struct {
			a2 int
			b2 string
		}{
			a2: 1,
			b2: "hiya",
		},
		c: 1 + 2i,
		d: 1.2345,
		e: []int{1, 2, 3},
		g: nil,
		h: []string{"one", "two", "three"},
		i: map[string]int{
			"key1": 88,
			"key2": 99,
			"key3": 10000,
		},
		j: false,
	}
	b, _ := sexpr.Marshal(m)
	fmt.Println(string(b))
}

type MyStruct struct {
	a int
	b bool
	c complex64
	d float64
	e Skip
	f struct {
		a2 int
		b2 string
	}
	g interface{}
	h []string
	i map[string]int
	j bool
}

type Skip interface {
	/* TODO: add methods */
}
