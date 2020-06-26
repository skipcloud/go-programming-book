package main

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestMain(t *testing.T) {
	type test struct {
		input string
		want  string
	}
	// the charcount program uses a map to store rune counts so the output
	// is nondeterministic, which makes the output impossible to test. The
	// following test will pass sometimes and fail others.
	tests := []test{
		{"hi", "rune\tcount\n'h'\t1\n'i'\t1\n\nlen\tcount\n1\t2\n2\t0\n3\t0\n4\t0\n"},
	}
	for _, tt := range tests {
		// create new files to use for Stdin and Stdout, this way we control
		// the input and can capture the output
		r, w, err := newFiles()
		if err != nil {
			t.Fatalf("error creating test read/write files: %v", err)
		}
		defer func() {
			w.Close()
			r.Close()
			os.Remove(w.Name())
			os.Remove(r.Name())
		}()
		// make the swap
		oldStdin := os.Stdin
		oldStdout := os.Stdout
		os.Stdin = w
		os.Stdout = r

		// write our test input
		n, err := w.WriteString(tt.input)
		if err != nil {
			t.Fatalf("error writing to stdin tmp file: %v", err)
		}
		if n != len(tt.input) {
			t.Fatal("error writing to stdin tmp file")
		}

		// ensure stdin file is read from start of file
		w.Seek(0, 0)
		main()
		// ensure stdout is read from start of file
		r.Seek(0, 0)
		got, err := ioutil.ReadAll(r)
		if err != nil {
			t.Fatalf("error reading output: %v", err)
		}
		// check against what we want
		if string(got) != tt.want {
			t.Errorf("got\n%s\n\nwant\n%s", got, tt.want)
		}

		// replace the Stdout and Stdin
		os.Stdin = oldStdin
		os.Stdout = oldStdout
	}
}

// newFiles creates two new files
func newFiles() (r *os.File, w *os.File, err error) {
	r, err = ioutil.TempFile(os.TempDir(), "read")
	if err != nil {
		return
	}

	w, err = ioutil.TempFile(os.TempDir(), "write")
	if err != nil {
		return
	}
	return
}
