package main

import "fmt"

/*
 * Write a version of rotate that operates in a single pass
 */

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}
	rotate(slice, 10)
	fmt.Println(slice)
}

func rotate(s []int, n int) {
	end := make([]int, len(s)-n)
	copy(end, s[n:])
	copy(s[len(s)-n:], s[:n])
	copy(s, end)
}
