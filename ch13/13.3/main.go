package main

import (
	"bytes"
	"compress/bzip2"
	"fmt"
	"sync"

	"github.com/skipcloud/go-programming-book/ch13/13.3/bzip"
)

/*
	Use sync.Mutex to make bzip2.Writer safe for concurrent use by
	multiple goroutines
*/

func main() {
	b := bytes.NewBuffer([]byte{})
	w := bzip.NewWriter(b)

	var wg sync.WaitGroup
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func(i int) {
			s := fmt.Sprintf("here is some text from goroutine number %d\n", i)
			w.Write([]byte(s))
			wg.Done()
		}(i)
	}
	wg.Wait()
	w.Close()
	fmt.Printf("compressed output:\n%s\n\n", b.String())

	r := bzip2.NewReader(b)
	output := make([]byte, 32*1024)
	r.Read(output)
	fmt.Printf("uncompressed output:\n%s", string(output))
}
