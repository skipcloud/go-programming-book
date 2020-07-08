package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/skipcloud/go-programming-book/ch12/12.5/sexpr"
)

/*
	Adapt encode to emit JSON instead of S-expressions. Test your
	encoder using the standard decoder, json.Unmarshal.
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
}

func main() {
	s := MyStruct{
		A: "hi",
		B: 138,
		C: map[string]string{
			"oh":      "a value",
			"another": "eep",
		},
		D: map[string]map[string]string{
			"mymaps": {
				"hello": "dude",
			},
		},
		E: []int{1, 2, 3, 4},
		F: &SimpleStruct{
			"the only field",
		},
	}
	b, err := sexpr.Marshal(s)
	if err != nil {
		log.Fatalf("error marshaling: %v", err)
	}
	fmt.Println(string(b))
	var s2 MyStruct
	err = json.Unmarshal(b, &s2)
	if err != nil {
		log.Fatalf("error unmarshaling: %v", err)
	}
	fmt.Printf("%#v", s2)
}
