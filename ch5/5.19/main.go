package main

import "fmt"

/*
	Use panic and recover to write a function that contains no return statement
	yet returns a non-zero value
*/

func main() {

	fmt.Println(myFunc())
}

func myFunc() (value string) {
	defer func() {
		value = recover().(string)
	}()
	panic("hi")
}
