package main

import (
	"fmt"
	"log"

	"github.com/skipcloud/go-programming-book/ch12/12.6/sexpr"
)

/*
	Adapt encode so that, as an optimization, it does not encode a field
	whose value is the zero value of it's type.
*/

type SimpleStruct struct {
	A string
}

type MyStruct struct {
	A string
	B int
	C map[string]string
	D map[string]map[string]string
	E []int
	F *SimpleStruct
	G SimpleStruct
	H [10]int
}

func main() {
	m := MyStruct{}
	b, err := sexpr.Marshal(m)
	if err != nil {
		log.Fatalf("error marshaling: %v", err)
	}
	fmt.Printf("All zero values: %v\n", string(b))

	m2 := MyStruct{
		A: "hello",
		B: 42,
		C: map[string]string{
			"key1": "value1",
		},
		D: map[string]map[string]string{
			"key2": {
				"key3": "value2",
			},
		},
		E: []int{1, 2, 3},
		F: &SimpleStruct{
			A: "F *struct",
		},
		G: SimpleStruct{
			A: "G struct",
		},
		H: [10]int{1},
	}
	b, err = sexpr.Marshal(m2)
	if err != nil {
		log.Fatalf("error marshaling: %v", err)
	}
	fmt.Printf("No zero values: %v\n", string(b))
}
