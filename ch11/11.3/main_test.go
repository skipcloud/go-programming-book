package main

import (
	"math/rand"
	"testing"
	"time"

	word "github.com/skipcloud/go-programming-book/ch11/11.3/word1"
)

/*
	TestRandomPalindromes only tests palindromes. Write a randomized test that
	generates and verifies non-palindromes
*/

func randomPalindrome(rng *rand.Rand) string {
	n := rng.Intn(25) // random length up to 24
	runes := make([]rune, n)
	for i := 0; i < (n+1)/2; i++ {
		r := rune(rng.Intn(0x1000)) // random rune up to '\u0999'
		runes[i] = r
		runes[n-1-i] = r
	}
	return string(runes)
}

func TestRandomPalindromes(t *testing.T) {
	// Initialize a pseudo-random number generator
	seed := time.Now().UTC().UnixNano()
	t.Logf("Random seed: %d", seed)
	rng := rand.New(rand.NewSource(seed))
	for i := 0; i < 10; i++ {
		p := randomPalindrome(rng)
		if !word.IsPalindrome(p) {
			t.Errorf("IsPalindrome(%q) = false", p)
		}
	}
}

func TestRandomNonPalindrome(t *testing.T) {
	seed := time.Now().UTC().UnixNano()
	t.Logf("Random seed: %d", seed)
	rng := rand.New(rand.NewSource(seed))
	// ensure we have at least 2 characters
	// because 0 and 1 characters are naturally palindromes
	n := rng.Intn(25) + 2
	for i := 0; i < 1000; i++ {
		runes := make([]rune, n)
		for j := 0; j < (n+2)/2; j++ {
			r := rune(rng.Intn(0x0999)) // up to \u0998 leaving us room to +1
			runes[j] = r
			runes[n-1-j] = r + 1
		}
		if word.IsPalindrome(string(runes)) {
			t.Errorf("IsPalindrome(%q) = true", string(runes))
		}
	}
}
