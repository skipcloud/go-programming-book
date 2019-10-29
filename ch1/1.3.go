// Echo3 prints its command-line arguments.
package main

// experiment to measure the difference in running time
// between our potentially ineddicient versiona dn the one
// that uses strings.Join.

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {

	start := time.Now()
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
	secs := time.Since(start).Seconds()
	fmt.Printf("echo1 - for: %fs\n\n", secs)

	start = time.Now()
	s, sep = "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
	secs = time.Since(start).Seconds()
	fmt.Printf("echo2 - range: %fs\n\n", secs)

	start = time.Now()
	fmt.Println(strings.Join(os.Args[1:], " "))
	secs = time.Since(start).Seconds()
	fmt.Printf("echo3 - strings.Join: %fs\n", secs)
}
