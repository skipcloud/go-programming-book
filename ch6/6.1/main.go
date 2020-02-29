package main

import (
	"bytes"
	"fmt"
)

/*
	Implement these additional methods

		func (*IntSet) Len() int      // return the number of elements
		func (*IntSet) Remove(x int)  // remove x from the set
		func (*IntSet) Clear()        // remove all elements from the set
		func (*IntSet) Copy() *IntSet // return a copy of the set
*/

func main() {
	s := IntSet{}
	s.Add(1)
	s.Add(10)
	s.Add(1000)
	s.Add(623)
	fmt.Printf("There are %d items in set\n", s.Len())

	fmt.Printf("Set has 623? %t\n", s.Has(623))
	fmt.Println("Removing 623")
	s.Remove(623)
	fmt.Printf("There are %d items in set now\n", s.Len())
	fmt.Printf("Set has 623? %t\n", s.Has(623))

	fmt.Println("copying set")
	c := s.Copy()
	fmt.Printf("There are %d items in copied set\n", c.Len())
	fmt.Println("adding to copy")
	c.Add(54321)
	fmt.Printf("There are %d items in copied set\n", c.Len())
	fmt.Printf("There are %d items in original set\n", s.Len())

	fmt.Println("Clearing original set")
	s.Clear()
	fmt.Printf("There are %d items in original set now\n", s.Len())

	fmt.Printf("There are %d items in copied set\n", c.Len())
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

func (s *IntSet) Len() int {
	leng := 0
	for _, word := range s.words {
		for j := 0; j < 64; j++ {
			if (word>>j)&1 != 0 {
				leng += 1
			}
		}
	}
	return leng
}

func (s *IntSet) Remove(x int) {
	word, bit := x/64, uint(x%64)
	if word > len(s.words) {
		// doesn't exist in set
		return
	}
	s.words[word] &= ^(1 << bit)
}

func (s *IntSet) Clear() {
	s.words = []uint64{}
}

func (s *IntSet) Copy() *IntSet {
	newSet := IntSet{
		words: make([]uint64, len(s.words)),
	}
	for i, word := range s.words {
		newSet.words[i] = word
	}
	return &newSet
}
