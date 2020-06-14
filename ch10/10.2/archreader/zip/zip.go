package zip

import (
	"archive/zip"
	"errors"
	"io"

	"github.com/skipcloud/go-programming-book/ch10/10.2/archreader"
)

func init() {
	r := ZipReader{}
	archreader.DefineFormat("zip", &r)
}

type ZipReader struct {
	r                *zip.ReadCloser
	currentFile      io.ReadCloser
	currentFileIndex int
}

func (z *ZipReader) New(file string) error {
	r, err := zip.OpenReader(file)
	if err != nil {
		return err
	}
	z.r = r
	f, err := z.r.File[0].Open()
	z.currentFile = f
	return nil
}

func (z *ZipReader) PrintName() {
	println(z.r.File[z.currentFileIndex].Name)
}

func (z *ZipReader) Read(p []byte) (int, error) {
	return z.currentFile.Read(p)
}

func (z *ZipReader) NextFile() error {
	// check there's still a file to read
	if z.currentFileIndex+1 >= len(z.r.File) {
		return errors.New("end of files")
	}
	// update index and set current file
	z.currentFileIndex += 1
	f := z.r.File[z.currentFileIndex]
	rc, err := f.Open()
	if err != nil {
		return err
	}
	z.currentFile = rc
	return nil
}
