package main

import (
	"fmt"
	. "gobook/ch2/popcount"
)

/*
 * The expression x&(x-1) clears the rightmost non-zero bit of x.
 * Write a version of PopCount that counts bits by using this
 * fact, and assess its performance.
 */

func main() {
	// Original function
	fmt.Println(PopCount(255))

	// New funtion
	fmt.Printf("%.64b\n", 255)
	fmt.Println(PopCount4(255))
}
