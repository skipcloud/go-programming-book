package main

import "fmt"

/*
 * Write an in-place function to eliminate adjacent duplicates in a []string slice
 */

func main() {
	s := []string{"hello", "there", "there", "skip"}
	dedup(&s)
	fmt.Println(s)
}

/*
 * in order to update the slice len so we don't display {"hello", "there", "skip", "skip"}
 * this function accepts a pointer to a slice
 */
func dedup(s *[]string) {
	i := 0

	for j := 0; j < len((*s)); j++ {
		if (j+1 == len((*s))) || (j+1 < len((*s)) && (*s)[j] != (*s)[j+1]) {
			(*s)[i] = (*s)[j]
			i++
		}
	}
	*s = (*s)[:i]
}
