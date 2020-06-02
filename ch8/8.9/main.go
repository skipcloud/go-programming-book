package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

/*
	Write a version of du that computes and periodically displays separate
	totals for each of the root directories
*/

func main() {
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}
	results := make(chan string, len(roots))
	// create separate channel for each directory
	for _, dir := range roots {
		fileSize := make(chan int64)
		go func(dir string, c chan int64) {
			walkDir(dir, c)
			close(c)
		}(dir, fileSize)
		go listenAndPrint(fileSize, results, dir)
	}

	output := bytes.NewBufferString("\nResults\n")
	for n := len(roots); n > 0; n-- {
		output.WriteString(<-results)
	}
	close(results)
	fmt.Fprintf(os.Stdout, output.String())
}

const outputFormat = "%.32s %d files\t%.1f GB\n"

func listenAndPrint(fileSize <-chan int64, result chan<- string, dir string) {
	ticker := time.Tick(500 * time.Millisecond)
	var size int64
	var count int64
loop:
	for {
		select {
		case s, ok := <-fileSize:
			if !ok {
				break loop
			}
			size += s
			count++
		case <-ticker:
			fmt.Printf(outputFormat, dir, count, float64(size)/1000000000)
		}
	}
	result <- fmt.Sprintf(outputFormat, dir, count, float64(size)/1000000000)
}

func walkDir(dir string, fileSize chan<- int64) {
	for _, f := range dirents(dir) {
		if f.IsDir() {
			walkDir(filepath.Join(dir, f.Name()), fileSize)
		} else {
			fileSize <- f.Size()
		}
	}
}

func dirents(dir string) []os.FileInfo {
	contents, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du: %v\n", err)
		return nil
	}
	return contents
}
