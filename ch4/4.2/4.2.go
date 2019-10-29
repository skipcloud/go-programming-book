package main

/*
 * Write a program that prints SHA256 hash of its standard input by default
 * but supports a command-line flag to print the SHA384 or SHA-512 hash instead
 */

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	hashFuncType := flag.Int("hash", 256, "SHA hash function type: 256, 384, or 512")
	flag.Parse()

	var input []byte
	var err error

	for {
		input, err = ioutil.ReadAll(os.Stdin)
		if err != nil {
			return
		}

		switch *hashFuncType {
		case 384:
			fmt.Printf("%x\n", sha512.Sum384(input))
		case 512:
			fmt.Printf("%x\n", sha512.Sum512(input))
		default:
			fmt.Printf("%x\n", sha256.Sum256(input))
		}
	}

}
