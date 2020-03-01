package main

import (
	"fmt"
	"io"
	"strings"
)

/*
	The LimitReader function in the io package accepts an io.Reader r and a
	number of bytes n, and returns another Reader that reads from r but reports
	an end-of-file condition after n bytes. Implement it.

		func LimitReader(r io.Reader, n int64) io.Reader
*/

func main() {
	l := LimitReader(strings.NewReader("Hi there"), 4)
	buf := make([]byte, 10)

	l.Read(buf)
	fmt.Println(buf)

	buf = make([]byte, 10)
	l.Read(buf)
	fmt.Println(buf)
}

type limitReader struct {
	reader io.Reader
	limit  int64
	i      int64
}

func (l *limitReader) Read(p []byte) (n int, err error) {
	if l.i >= l.limit {
		return 0, io.EOF
	}
	buf := make([]byte, l.limit-l.i)
	_, err = l.reader.Read(buf)
	n = copy(p, buf)
	l.i += int64(n)
	if err != nil {
		return n, err
	}
	return n, nil
}

func LimitReader(r io.Reader, n int64) io.Reader {
	return &limitReader{
		reader: r,
		limit:  n,
	}
}
