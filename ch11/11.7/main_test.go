package main

import (
	"math"
	"math/rand"
	"testing"
	"time"

	"github.com/skipcloud/go-programming-book/ch11/11.2/intset"
	"github.com/skipcloud/go-programming-book/ch11/11.2/mapset"
)

/*
	Write benchmarks for Add, UnionWith and other methods of *IntSet (§6.5)
	using large pseudo-random inputs. How fast can you make these methods run?
	How does the choice of word size affect performance? How fast is IntSet
	compared to a set implementation based on the built-in map type?

	How fast?
			   That all depends on the input really.

	Word size?
			   Bigger word sizes really affect the performance of the IntSet
			   methods, I originally used an int64 for the random number but
			   the computer froze and the bench marks fails after 280 seconds.
			   Mainly because the program ran out of memory.

	Speed against map type?
			   The maptype is infinitely faster, it seems there's very few (if any)
			   memory allocations for the map type.
*/

func BenchmarkIntSetAdd(b *testing.B) {
	input := generateInput()
	ints := intset.IntSet{}
	b.Logf("Generated input: %d", input)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ints.Add(input)
	}
}

func BenchmarkIntSetHas(b *testing.B) {
	input := generateInput()
	ints := intset.IntSet{}
	ints.Add(input)
	b.Logf("Generated input: %d", input)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ints.Has(input)
	}
}

func BenchmarkIntSetUnionWith(b *testing.B) {
	input := generateInput()
	ints2 := intset.IntSet{}
	ints2.Add(input)
	ints1 := intset.IntSet{}
	b.Logf("Generated input: %d", input)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ints1.UnionWith(&ints2)
	}
}

func BenchmarkMapSetAdd(b *testing.B) {
	input := generateInput()
	maps := mapset.MapIntSet{}
	b.Logf("Generated input: %d", input)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		maps.Add(input)
	}
}

func BenchmarkMapSetHas(b *testing.B) {
	input := generateInput()
	maps := mapset.MapIntSet{}
	maps.Add(input)
	b.Logf("Generated input: %d", input)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		maps.Has(input)
	}
}

func BenchmarkMapSetUnionWith(b *testing.B) {
	input := generateInput()
	maps2 := mapset.MapIntSet{}
	maps2.Add(input)
	maps1 := mapset.MapIntSet{}
	b.Logf("Generated input: %d", input)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		maps1.UnionWith(maps2)
	}
}

func generateInput() int {
	seed := time.Now().UTC().UnixNano()
	rng := rand.New(rand.NewSource(seed))
	return rng.Intn(math.MaxInt32)
}
