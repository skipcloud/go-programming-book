package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

/*
 * Modify charcount to count letters, digits, and so on in their
 * unicode categories, using functions like unicode.IsLetter
 */

func main() {
	var utflen [utf8.UTFMax + 1]int
	counts := make(map[rune]int)
	invalid := 0
	letters := 0
	digits := 0
	punctuation := 0

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune() // returns rune, nbytes, error
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		if unicode.IsLetter(r) {
			letters++
		}
		if unicode.IsNumber(r) {
			digits++
		}
		if unicode.IsPunct(r) {
			punctuation++
		}
		counts[r]++
		utflen[n]++
	}
	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Print("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	if letters > 0 {
		fmt.Printf("\n%d letters\n", letters)
	}
	if digits > 0 {
		fmt.Printf("\n%d numbers\n", digits)
	}
	if punctuation > 0 {
		fmt.Printf("\n%d punctuation characters\n", punctuation)
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}
