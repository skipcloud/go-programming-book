package main

import (
	"fmt"
	"sort"
)

/*
	The sort.Interface type can be adapted to other uses. Write a function IsPalindrome(s sort.Interface) bool
	that reports whether the sequence s is a palindrome, in other words, reversing the sequence would not
	change it. Assume that the elements at indices i and j are equal if !s.Less(i, j) && !s.Less(j, i)
*/

func main() {
	words := Words([]string{"hello", "hello"})
	fmt.Println(IsPalindrome(words))

	words = Words([]string{"this", "isn't", "a", "palindrome"})
	fmt.Println(IsPalindrome(words))

	// even number of elements
	words = Words([]string{"this", "is", "a", "palindrome", "palindrome", "a", "is", "this"})
	fmt.Println(IsPalindrome(words))

	// odd number of elements
	words = Words([]string{"this", "is", "a", "palindrome", "holla", "palindrome", "a", "is", "this"})
	fmt.Println(IsPalindrome(words))
}

func IsPalindrome(s sort.Interface) bool {
	for i, j := 0, s.Len()-1; i < s.Len()/2; i, j = i+1, j-1 {
		if !(!s.Less(i, j) && !s.Less(j, i)) {
			return false
		}
	}
	return true
}

type Words []string

// Len is the number of elements in the collection.
func (w Words) Len() int {
	return len(w)
}

// Less reports whether the element with
// index i should sort before the element with index j.
func (w Words) Less(i int, j int) bool {
	return w[i] < w[j]
}

// Swap swaps the elements with indexes i and j.
func (w Words) Swap(i int, j int) {
	w[i], w[j] = w[j], w[i]
}
