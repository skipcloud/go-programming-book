package main

import (
	"fmt"

	"github.com/skipcloud/go-programming-book/ch12/12.2/display"
)

/*
	Make display safe to use on cyclical data structures by bounding the number
	of steps it takes before abandoning the recursion. (In section 13.3, we'll see
	another way to detect cycles)
*/

func main() {
	// a struct that references another struct of the same type
	type Me struct {
		name   string
		friend *Me
	}
	var m Me
	m = Me{
		name: "Skip",
		// cyclical reference
		friend: &m,
	}
	display.Display("me", m)
	fmt.Println()
}
