package main

import (
	"crypto/sha256"
	"fmt"
)

/*
 * Write a function that counts the number of bits that are different in
 * two SHA256 hashes. (See popcount from section 2.6.2)
 */

func main() {
	hash1 := sha256.Sum256([]byte("x"))
	hash2 := sha256.Sum256([]byte("X"))
	popCount(hash1, hash2)
}

func popCount(hash1 [32]byte, hash2 [32]byte) {
	const byteLen = 8
	var count int

	for i, v := range hash1 {
		for j := uint(0); j < byteLen; j++ {
			if ((v >> j) & 1) != ((hash2[i] >> j) & 1) {
				count++
			}
		}

	}
	fmt.Println(count)
