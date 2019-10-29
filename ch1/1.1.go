// Echo3 prints its command-line arguments.
package main

// Modify the echo program to also print
// os.Args[0], the name of the command
// that invoked it.

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(strings.Join(os.Args[:], " "))
}
