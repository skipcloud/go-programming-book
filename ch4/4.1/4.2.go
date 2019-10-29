package main

/*
 * Write a program that prints SHA256 hash of its standard input by default
 * but supports a command-line flag to print the SHA384 or SHA-512 hash instead
 */

import (
	"flag"
	"fmt"
)

func main() {
	hashFuncType := flag.Int("hash", 256, "SHA hash function type: 256, 384, or 512")
	flag.Parse()
	fmt.Println(hashFuncType)
}
