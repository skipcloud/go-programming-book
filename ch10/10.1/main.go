// The jpeg command reads a PNG image from the standard input
// and writes it as a JPEG image to the standard output.
package main

import (
	"errors"
	"flag"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"os"
	"strings"
)

/*
	Extend the jpeg program so that it converts any supported input format to any
	output format, using image.Decode to detect the input format and a flag to
	select the output format.
*/

type format string

func main() {
	outputFormat := flag.String("format", "jpeg", "the desired output format for the image")
	flag.Parse()

	if err := toFormat(*outputFormat, os.Stdin, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "jpeg: %v\n", err)
		os.Exit(1)
	}
}

func toFormat(format string, in io.Reader, out io.Writer) error {
	format = strings.ToLower(format)
	img, kind, err := image.Decode(in)
	if err != nil {
		return err
	} else if kind == format {
		return errors.New(fmt.Sprintf("image is already %s", format))
	}

	fmt.Fprintln(os.Stderr, "Input format =", kind)

	switch strings.ToLower(format) {
	case "jpeg":
		return jpeg.Encode(out, img, &jpeg.Options{Quality: 95})
	case "png":
		return png.Encode(out, img)
	case "gif":
		return gif.Encode(out, img, &gif.Options{NumColors: 255})
	default:
		return errors.New(fmt.Sprintf("unsupported format %s", format))
	}
}
