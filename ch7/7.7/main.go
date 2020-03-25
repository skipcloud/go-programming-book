package main

import "fmt"

/*
	Explain why the help message contains °C when the default calue of 20.0 does not
*/

/*
	In the flag package the function Var() accepts a value and called value.String to
	set the DefaultValue of a flag. Because .String() is defined on the Celsius type
	which we have embedded Celsius into celsiusFlag we get the String method too.
*/

func main() {
	fmt.Println("hello")
}
