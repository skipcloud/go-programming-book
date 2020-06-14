package tar

import (
	"archive/tar"
	"os"

	"github.com/skipcloud/go-programming-book/ch10/10.2/archreader"
)

func init() {
	r := TarReader{}
	archreader.DefineFormat("tar", &r)
}

type TarReader struct {
	r           *tar.Reader
	currentFile *tar.Header
}

// New takes a filename string as it's sole argument and populates
// the ArchReader, i.e. opens the file and does package specific setup
func (t *TarReader) New(file string) error {
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	t.r = tar.NewReader(f)
	return nil
}

// Read reads from the current file into the byte slice arg
// returning number of bytes read and/or an error
func (t *TarReader) Read(p []byte) (int, error) {
	return t.r.Read(p)
}

// NextFile proceeds to the next file in the archive or returns
// an error
func (t *TarReader) NextFile() error {
	h, err := t.r.Next()
	if err != nil {
		return err
	}
	t.currentFile = h
	return nil
}

func (t *TarReader) PrintName() {
	println(t.currentFile.Name)
}
