package main

import (
	"fmt"

	. "gobook/ch2/popcount"
)

/*
 * Rewrite popcount to use a loop instead of a single expression. Compare
 * the performance of the two versions. (Section 11.4 shows how to compare
 * the performance of different implementations systematically)
 */

func main() {
	// Original function
	fmt.Println(PopCount(255))

	// New funtion
	fmt.Println(PopCount2(255))

}
