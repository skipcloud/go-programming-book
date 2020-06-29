package main

import (
	"testing"

	"gopl.io/ch2/popcount"

	mypopcount "github.com/skipcloud/go-programming-book/ch2/popcount"
)

/*
	Write benchmarks to compare the PopCount implementation in Section 2.6.2
	with your solutions to Exercise 2.4 and Exercise 2.5. At what point does
	the table-based approach break even?

	Answer: at around a million iterations the times seem to level off
*/

const popInput = 1010101010

func BenchmarkPopCount(b *testing.B) {
	// bench mark for original PopCount
	for i := 0; i < 1000000; i++ {
		popcount.PopCount(popInput)
	}
}

func BenchmarkPopCount3(b *testing.B) {
	// bench mark for Exercise 2.4
	for i := 0; i < 1000000; i++ {
		mypopcount.PopCount3(popInput)
	}
}

func BenchmarkPopCount4(b *testing.B) {
	// bench mark for Exercise 2.4
	for i := 0; i < 1000000; i++ {
		mypopcount.PopCount4(popInput)
	}
}
