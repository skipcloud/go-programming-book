package main

import (
	"os"

	"github.com/skipcloud/go-programming-book/ch12/12.7/sexpr"
)

/*
	Create a streaming API for the S-expression encoder, following
	the style of json.Encoder
*/

type skip struct {
	Name string
}

func main() {
	s := skip{Name: "alan"}
	e := sexpr.NewEncoder(os.Stdout)
	e.Encode(&s)
}
