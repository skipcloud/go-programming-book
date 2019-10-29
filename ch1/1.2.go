// Echo3 prints its command-line arguments.
package main

// Modify the echo program to print the index
// and value of each of its arguments, one per line

import (
	"fmt"
	"os"
)

func main() {
	for i, arg := range os.Args {
		fmt.Println(i, arg)
	}
}
