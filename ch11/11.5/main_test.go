package main

import (
	"strings"
	"testing"
)

/*
	Extend TestSplit to use a table of inputs and expected outputs
*/

func TestSplit(t *testing.T) {
	tests := []struct {
		str  string
		sep  string
		want int
	}{
		{"a:b:c", ":", 3},
		{"ababa", "b", 3},
		{"m''m", "'", 3},
	}
	for _, test := range tests {
		words := strings.Split(test.str, test.sep)
		if got := len(words); got != test.want {
			t.Errorf("Split(%q, %q) returned %d words, want %d",
				test.str, test.sep, got, test.want)
		}
	}
}
