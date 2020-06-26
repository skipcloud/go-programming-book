package intset

import (
	"bytes"
	"fmt"
	"testing"
)

func Example_one() {
	//!+main
	var x, y IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	fmt.Println(x.String()) // "{1 9 144}"

	y.Add(9)
	y.Add(42)
	fmt.Println(y.String()) // "{9 42}"

	x.UnionWith(&y)
	fmt.Println(x.String()) // "{1 9 42 144}"

	fmt.Println(x.Has(9), x.Has(123)) // "true false"
	//!-main

	// Output:
	// {1 9 144}
	// {9 42}
	// {1 9 42 144}
	// true false
}

func Example_two() {
	var x IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	x.Add(42)

	//!+note
	fmt.Println(&x)         // "{1 9 42 144}"
	fmt.Println(x.String()) // "{1 9 42 144}"
	fmt.Println(x)          // "{[4398046511618 0 65536]}"
	//!-note

	// Output:
	// {1 9 42 144}
	// {1 9 42 144}
	// {[4398046511618 0 65536]}
}

// MapIntSet is my built-in map implementation for IntSet, this will be
// used to test against the IntSet implementation.
type MapIntSet map[uint64]struct{}

func (m MapIntSet) Has(x int) bool {
	_, ok := m[uint64(x)]
	return ok
}

func (m MapIntSet) Add(x int) {
	m[uint64(x)] = struct{}{}
}

func (m MapIntSet) UnionWith(mi MapIntSet) {
	for k := range mi {
		if _, ok := m[k]; !ok {
			m[k] = struct{}{}
		}
	}
}

func (m MapIntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for k := range m {
		if buf.Len() > len("{") {
			buf.WriteByte(' ')
		}
		fmt.Fprintf(&buf, "%d", k)
	}
	buf.WriteByte('}')
	return buf.String()
}

func TestHas(t *testing.T) {
	input := []int{1, 2, 3, 4, 65, 66, 67}
	mapSet := MapIntSet{}
	intSet := &IntSet{
		words: []uint64{0, 0},
	}

	// Test for values that should exist
	for _, v := range input {
		mapSet[uint64(v)] = struct{}{}
		idx := 0
		if v >= 64 {
			idx = 1
		}
		// manually set bits
		// brittle as hell
		intSet.words[idx] |= 1 << (v % 64)

		expected := true
		mapValue := mapSet.Has(v)
		got := intSet.Has(v)
		if got != expected || got != mapValue {
			t.Errorf("Has(%d) = %t; expected %t", v, got, expected)
		}
	}

	// Test for values that should not exist
	for _, v := range []int{5, 6, 7, 8} {
		expected := false
		mapValue := mapSet.Has(v)
		got := intSet.Has(v)
		if got != expected || got != mapValue {
			t.Errorf("Has(%d) = %t; expected %t", v, got, expected)
		}
	}
}

func TestAdd(t *testing.T) {
	input := []int{1, 2, 3, 4, 65, 66, 67}
	mapSet := MapIntSet{}
	intSet := &IntSet{}

	// Test for values that should exist
	for _, v := range input {
		mapSet.Add(v)
		intSet.Add(v)
		expected := true
		mapValue := mapSet.Has(v)
		got := intSet.Has(v)
		if got != expected || got != mapValue {
			t.Errorf("Add(%d) did not add value", v)
		}
	}
}

func TestUnionWith(t *testing.T) {
	// Set up
	firstSet := []int{1, 2, 3, 4}
	secondSet := []int{65, 111, 200, 3000, 1013458}

	intSetOne := &IntSet{}
	intSetTwo := &IntSet{}

	mapSetOne := MapIntSet{}
	mapSetTwo := MapIntSet{}

	for _, v := range firstSet {
		intSetOne.Add(v)
		mapSetOne.Add(v)
	}
	for _, v := range secondSet {
		intSetTwo.Add(v)
		mapSetTwo.Add(v)
	}

	// Perform action
	intSetOne.UnionWith(intSetTwo)
	mapSetOne.UnionWith(mapSetTwo)

	for _, v := range append(firstSet, secondSet...) {
		expected := true
		mapValue := mapSetOne.Has(v)
		got := intSetOne.Has(v)
		if got != expected || got != mapValue {
			t.Errorf("UnionWith(%T) didn't add %d", intSetTwo, v)
		}
	}
}
