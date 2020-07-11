package main

import (
	"fmt"
	"os"

	"github.com/skipcloud/go-programming-book/ch12/12.8/sexpr"
)

/*
	The sexpr.Unmarshal function, like json.Unmarshal, requires the complete
	input in a byte slice before it can begin decoding. Define a sexpr.Decoder
	type that, like json.Decoder,  allows a sequence of values to be decoded
	from an io.Reader. Change sexpr.Unmarshal to use this new type.
*/

type skip struct {
	Name string
}

func main() {
	s := skip{"alan"}
	b, _ := sexpr.Marshal(s)
	fmt.Println(string(b)) // to use as input
	var s2 skip

	fmt.Printf("enter S-expression: ")
	d := sexpr.NewDecoder(os.Stdin)
	d.Decode(&s2)
	fmt.Printf("output: %#v\n", s2)
}
