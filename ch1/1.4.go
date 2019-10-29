package main

// Modify dup2 to print the names of all the files
// in which each duplicated line occurs.

import (
	"bufio"
	"fmt"
	"os"
)

type Dups struct {
	Count int
	Files map[string]int
}

func main() {
	counts := make(map[string]Dups)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for line, dup := range counts {
		if dup.Count > 1 {
			fmt.Printf("%d\t%s\n", dup.Count, line)
			for file, count := range dup.Files {
				fmt.Printf("%s: %d\n", file, count)
			}
		}
	}
}

func countLines(f *os.File, counts map[string]Dups) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		line := counts[input.Text()]
		line.Count++
		if line.Files == nil {
			line.Files = make(map[string]int)
		}
		line.Files[f.Name()]++
		counts[input.Text()] = line
	}
}
