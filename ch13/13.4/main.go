package main

import (
	"bytes"
	"compress/bzip2"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/skipcloud/go-programming-book/ch13/13.4/bzip"
)

/*
	Depending on C libraries has its drawbacks. Provide an alternative pure-Go
	implementation of bzip.NewWriter that uses the os/exec package to run the
	/bin/bzip2 as a subprocess
*/

func main() {
	s := []byte("here is some text")
	b := bytes.NewBuffer([]byte{})
	w := bzip.NewWriter(b)

	// compress
	w.Write(s)
	fmt.Printf("compressed text\n%s\n", b.String())

	// decompress
	r := bzip2.NewReader(strings.NewReader(b.String()))
	bs, _ := ioutil.ReadAll(r)
	fmt.Printf("decompressed text\n%q\n", string(bs))
}
