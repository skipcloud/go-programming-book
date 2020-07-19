package main

import (
	"fmt"
	"log"

	"github.com/skipcloud/go-programming-book/ch12/12.13/sexpr"
)

/*
	Modify the S-Expression encoder (§12.4) and decoder (§12.6) so that
	they honor the 'sexpr:"..."' field tag in a similar manner to
	encoding/json (§4.5)

	edit: not sure what they mean by "..." so I'm going to assume they just mean
	      allowing you to define names for fields and omit empty fields etc
*/

type skip struct {
	Name    string            `sexpr:"-"`
	Age     int               `sexpr:"a"`
	Hobbies []string          `sexpr:"b"`
	Secret  string            `sexpr:",omitempty"`
	Things  map[string]string `sexpr:"c"`
}

func main() {
	s := skip{
		Name:    "gary",
		Age:     123,
		Hobbies: []string{"fishing", "cards"},
		Things: map[string]string{
			"key": "value",
		},
	}
	fmt.Printf("before marshalling: %#v\n", s)

	b, err := sexpr.Marshal(&s)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("marshalled: %s\n", b)

	var s2 skip
	err = sexpr.Unmarshal(b, &s2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("unmarshalled: %#v\n", s2)
}
