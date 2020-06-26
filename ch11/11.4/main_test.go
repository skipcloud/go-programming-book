package main

import (
	"math/rand"
	"testing"
	"time"

	word "github.com/skipcloud/go-programming-book/ch11/11.3/word1"
)

/*
	Modify randomPalindrome to exercise IsPalindrome's handling
	of punctuation and spaces.
*/

func randomPalindrome(rng *rand.Rand, addPunct bool) string {
	n := rng.Intn(25) // random length up to 24
	runes := make([]rune, n)
	for i := 0; i < (n+1)/2; i++ {
		r := rune(rng.Intn(0x1000)) // random rune up to '\u0999'
		runes[i] = r
		runes[n-1-i] = r
	}
	if addPunct {
		r := make([]rune, len(runes)+1)
		copy(r, runes)
		r[len(r)-1] = rune('.')
		runes = r
	}
	return string(runes)
}

func TestRandomPalindromes(t *testing.T) {
	// Initialize a pseudo-random number generator
	seed := time.Now().UTC().UnixNano()
	t.Logf("Random seed: %d", seed)
	rng := rand.New(rand.NewSource(seed))
	for i := 0; i < 10; i++ {
		addPunct := i%2 == 0
		p := randomPalindrome(rng, addPunct)
		if addPunct {
			if word.IsPalindrome(p) {
				t.Errorf("IsPalindrome(%q) = true", p)
			}
		} else {
			if !word.IsPalindrome(p) {
				t.Errorf("IsPalindrome(%q) = false", p)
			}
		}
	}
}
