package main

import "github.com/skipcloud/go-programming-book/ch12/12.11/params"

/*
	Write the corresponding Pack function. Given a struct value, Pack should
	return a URL incorporating the parameter values from the struct
*/

type skip struct {
	Age      int      `http:"a"`
	Kids     float64  `http:"c"`
	Name     string   `http:"n"`
	Employee bool     `http:"e"`
	Hobbies  []string `http:"h"`
}

func main() {
	baseurl := "http://localhost:1234"
	s := skip{
		Age:      4,
		Kids:     1.3,
		Name:     "alan",
		Employee: false,
		Hobbies:  []string{"plate spinning", "keyboards"},
	}
	u, _ := params.Pack(baseurl, &s)
	println(u)
}
