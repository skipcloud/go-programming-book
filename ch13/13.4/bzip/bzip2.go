package bzip

import (
	"bytes"
	"io"
	"os/exec"
)

const bzipPath = "/bin/bzip2"

type writer struct {
	w   io.Writer
	cmd string
}

func NewWriter(out io.Writer) io.Writer {
	w := &writer{w: out, cmd: bzipPath}
	return w
}

func (w *writer) Write(p []byte) (n int, err error) {
	c := exec.Command("bzip2", "--stdout")
	c.Stdin = bytes.NewReader(p)
	c.Stdout = w.w

	err = c.Run()
	n = len(p)
	return
}
