package main

import (
	"fmt"
	"io"
	"os"
)

/*
	Write a function CountingWriter with the signature below that, given an
	io.Writer, returns a new Writer that wraps the original, and a pointer to
	an int64 variable that at any moment contains the numbers of bytes written
	to the new Writer.

		func CountingWriter(w io.Writer) (io.Writer, *int64)
*/

func main() {

	w, c := CountingWriter(os.Stdout)
	fmt.Fprintf(w, "here are some words to print\n")
	fmt.Printf("There were %d bytes printed\n", *c)
}

type countingWriter struct {
	count  int64
	writer io.Writer
}

func (c *countingWriter) Write(p []byte) (int, error) {
	b, err := c.writer.Write(p)
	c.count += int64(b)
	return b, err
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	c := countingWriter{writer: w}
	return &c, &c.count
}
