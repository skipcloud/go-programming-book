package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/skipcloud/go-programming-book/ch10/10.2/archreader"
	_ "github.com/skipcloud/go-programming-book/ch10/10.2/archreader/tar"
	_ "github.com/skipcloud/go-programming-book/ch10/10.2/archreader/zip"
)

/*
	Define a generic archive file-reading function capable of reading ZIP file
	(archive/zip) and POSIX tar files (archive/tar). Use a registration mechanism
	similar to the one described above so that support for each file format can
	be plugged in using blank imports

	edit by Skip: not wanting to spend too much time on this exercise I have done
	the bare minimum required to make this work. Errors aren't handled correctly,
	but it will print the file names and the contents to standard out.
*/

func main() {
	format := flag.String("format", "", "the format of the file to be read")
	flag.Parse()
	var validFormat bool

	if *format == "" {
		log.Fatal("missing format option")
	}

	for _, f := range archreader.Formats() {
		if f == *format {
			validFormat = true
		}
	}
	if !validFormat {
		log.Fatalf("format '%s' is not supported", *format)
	}

	if len(flag.Args()) != 1 {
		log.Fatal("one file path required")
	}
	r, err := archreader.For(*format)
	if err != nil {
		log.Fatalf("%v", err)
	}
	r.New(flag.Args()[0])
	if err = printFileContents(r); err != nil {
		fmt.Printf("%v", err)
	}
}

func printFileContents(r archreader.ArchReader) error {
	buf := make([]byte, 1024)
	for {
		n, err := r.Read(buf)
		if err != nil {
			if n > 0 {
				io.Copy(os.Stdout, bytes.NewReader(buf))
			}
			if err == io.EOF {
				if err = r.NextFile(); err != nil {
					return err
				}
				r.PrintName()
			}
		}
	}
}
