package main

import (
	"fmt"
)

/*
 * Write a funtion that reports whether two strings are anagrams
 * of each other, that is, that contain the same letters in a
 * different order.
 */

func main() {
	fmt.Println(anagram("funeral", "realfun"))
}

func anagram(s1 string, s2 string) bool {

	if len(s1) != len(s2) {
		return false
	}

	s1Letters := make(map[rune]int)
	s2Letters := make(map[rune]int)

	for _, r := range s1 {
		s1Letters[rune(r)]++
	}
	for _, r := range s2 {
		s2Letters[rune(r)]++
	}

	for char, s1Amount := range s1Letters {
		if s2Amount, ok := s2Letters[char]; !ok || s2Amount != s1Amount {
			return false
		}

	}
	return true
}
