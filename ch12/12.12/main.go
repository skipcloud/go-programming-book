package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/skipcloud/go-programming-book/ch12/12.12/params"
)

/*
	Extend the field tag notation to express parameter validity requirements.
	For example, a string might need to be a valid email address or credit-card
	number, and an integer might need to be a valid US ZIP code. Modify Unpack
	to check these requirements.
*/

type skip struct {
	Age      int      `http:"a"`
	Name     string   `http:"n"`
	Email    string   `http:"m" validate:"email"`
	Employee bool     `http:"e"`
	Hobbies  []string `http:"h"`
	ZIP      int      `http:"z" validate:"zip"`
}

func main() {
	// create an HTTP request (can't be bothered building a server)
	input := "http://localhost:1234?a=400&n=Kevin&e=t&h=meep&h=moop"
	// add "valid" email
	input += "&m=valid@email.com"
	// add valid ZIP
	input += "&z=12345"
	u, _ := url.Parse(input)
	r := http.Request{
		URL: u,
	}

	var s skip
	err := params.Unpack(&r, &s)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v\n", s)
}
