package archreader

import "errors"

type ArchReader interface {
	// New takes a filename string as it's sole argument and populates
	// the ArchReader, i.e. opens the file and does package specific setup
	New(string) error
	// Read reads from the current file into the byte slice arg
	// returning number of bytes read and/or an error
	Read([]byte) (int, error)
	// NextFile proceeds to the next file in the archive or returns
	// an error
	NextFile() error
	PrintName()
}

var ErrUnsupportedFormat = errors.New("unsupported format")
var formats = map[string]ArchReader{}

func For(format string) (ArchReader, error) {
	r, ok := formats[format]
	if !ok {
		return nil, ErrUnsupportedFormat
	}
	return r, nil
}

func DefineFormat(format string, f ArchReader) {
	formats[format] = f
}

func Formats() []string {
	output := []string{}
	for key, _ := range formats {
		output = append(output, key)
	}
	return output
}
