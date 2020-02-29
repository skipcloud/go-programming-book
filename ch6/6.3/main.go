package main

import (
	"bytes"
	"fmt"
)

/*
	(*IntSet).UnionWith computes the union of two sets using |, the word-
	parallel bitwise OR operator. Implement methods for IntersectWith,
	DifferenceWith, and SymmetricDifference for the corresponding set
	operations. (The symmetric difference of two sets contains the elements
	present in one set or the other but not both)
*/

func main() {

	fmt.Printf("%b\n", 1^3)
	fmt.Println("IntersectionWith")
	s := &IntSet{}
	s.Add(1)
	s.Add(2)
	s.Add(3)

	t := &IntSet{}
	t.Add(1)
	t.Add(4)
	t.Add(5)

	s.IntersectWith(t)
	fmt.Printf("s has 1? %t\n", s.Has(1))
	fmt.Printf("s has 2? %t\n", s.Has(2))
	fmt.Printf("s has 3? %t\n", s.Has(3))

	fmt.Println("DifferenceWith")
	s = &IntSet{}
	s.Add(1)
	s.Add(2)
	s.Add(3)

	t = &IntSet{}
	t.Add(1)
	t.Add(4)
	t.Add(5)

	s.DifferenceWith(t)
	fmt.Printf("s has 1? %t\n", s.Has(1))
	fmt.Printf("s has 2? %t\n", s.Has(2))
	fmt.Printf("s has 3? %t\n", s.Has(3))

	fmt.Println("SymmetricDifference")
	s = &IntSet{}
	s.Add(1)
	s.Add(2)
	s.Add(3)

	t = &IntSet{}
	t.Add(1)
	t.Add(4)
	t.Add(5)

	s.SymmetricDifference(t)
	fmt.Printf("s has 1? %t\n", s.Has(1))
	fmt.Printf("s has 2? %t\n", s.Has(2))
	fmt.Printf("s has 3? %t\n", s.Has(3))
	fmt.Printf("s has 4? %t\n", s.Has(4))
	fmt.Printf("s has 5? %t\n", s.Has(5))
}

// An IntSet is a set of small non-negative integers.
// Its zero value represents the empty set.
type IntSet struct {
	words []uint64
}

// Has reports whether the set contains the non-negative value x.
func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add adds the non-negative value x to the set.
func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// UnionWith sets s to the union of s and t.
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// String returns the set as a string of the form "{1 2 3}".
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

func (s *IntSet) IntersectWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] &= tword
		}
		// if word not in s.words just ignore it,
		// no point ANDing on zero
	}
}

func (s *IntSet) DifferenceWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] &= ^tword
		}
		// if word not in s.words just ignore it,
		// no point getting difference with zero
	}
}

func (s *IntSet) SymmetricDifference(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] = s.words[i] ^ tword
		} else {
			// if word not in s.words then add it
			// as we have none of those bits
			s.words = append(s.words, tword)
		}
	}
}
