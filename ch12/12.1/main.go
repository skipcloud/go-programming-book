package main

import (
	"fmt"

	"github.com/skipcloud/go-programming-book/ch12/12.1/display"
)

/*
	Extend Display so that it can display maps whose keys are structs or arrays
*/

func main() {
	// A map which uses a struct as the key value
	type Key struct {
		key string
	}
	type Me struct {
		name    string
		friends map[Key]string
	}
	m := Me{
		name: "Skip",
		friends: map[Key]string{
			Key{
				key: "kevin",
			}: "hi",
			Key{
				key: "malone",
			}: "there",
		},
	}
	display.Display("me", m)
	fmt.Println()

	// A map which uses an array as a map key
	type Key2 [10]int
	type Person map[Key2]string
	k := Key2{1, 2, 3, 4, 5, 6}
	p := Person{
		k: "hi",
	}
	display.Display("p", p)
}
