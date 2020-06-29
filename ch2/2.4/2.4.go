package main

import (
	"fmt"

	. "github.com/skipcloud/go-programming-book/ch2/popcount"
)

/*
 * Write a version of PopCount that counts bits by shifting its argument through 64
 * bit positions,  testing the right most bit each time. Compare its performance to
 * the table-lookup version.
 */

func main() {
	// Original function
	fmt.Println(PopCount(255))

	// New funtion
	fmt.Println(PopCount3(255))
}
