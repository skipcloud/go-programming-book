package main

import (
	"bytes"
	"fmt"
	"io"
	"log"

	"github.com/skipcloud/go-programming-book/ch12/12.9/sexpr"
)

/*
	Write a token-based API for decoding S-expressions, following the style
	of xml.Decoder (§7.14). You will need five types of tokens: Symbol, String,
	Int, StartList, and Endlist.
*/

type skip struct {
	Name string
	Age  int
}

func main() {
	s := skip{
		Name: "Alan",
		Age:  102,
	}
	b, _ := sexpr.Marshal(&s)
	d := sexpr.NewDecoder(bytes.NewReader(b))
	for {
		t, err := d.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		switch t := t.(type) {
		case sexpr.StartList:
			fmt.Printf("{")
		case sexpr.Symbol:
			fmt.Printf("Symbol: %q = ", t.Name) // a value should follow
		case sexpr.String:
			fmt.Printf("%q", t.Text)
		case sexpr.Int:
			fmt.Printf("%d", t.Int)
		case sexpr.EndList:
			fmt.Printf("}")
		}
	}
}
