package main

import (
	"fmt"
	"log"
	"os"

	"github.com/skipcloud/go-programming-book/ch8/8.10/links"
)

/*
	HTTP requests may be cancelled by closing the optional Cancel channel in
	the http.Request struct. Modify the web crawler of Section 8.6 to support
	cancellation.

	Hint: the http.Get convenience function does not give you an opportunity
	to customize a Request. Instead, create the request using http.NewRequest,
	set its Cancel field, then perform the request by calling
	http.DefaultClient.Do(req)
*/

func crawl(url string, cancel chan struct{}) []string {
	fmt.Println(url)
	list, err := links.Extract(url, cancel)
	if err != nil {
		log.Print(err)
	}
	return list
}

//!+
func main() {
	worklist := make(chan []string)  // lists of URLs, may have duplicates
	unseenLinks := make(chan string) // de-duplicated URLs
	cancel := make(chan struct{})

	// Add command-line arguments to worklist.
	go func() { worklist <- os.Args[1:] }()
	// create option to cancel
	go func() {
		b := make([]byte, 1)
		os.Stdin.Read(b)
		close(cancel)
		close(worklist)
		close(unseenLinks)
	}()

	// Create 20 crawler goroutines to fetch each unseen link.
	for i := 0; i < 20; i++ {
		go func() {
			for link := range unseenLinks {
				go func(link string) {
					select {
					case <-cancel:
						break
					case worklist <- crawl(link, cancel):
					}
				}(link)
			}
		}()
	}

	// The main goroutine de-duplicates worklist items
	// and sends the unseen ones to the crawlers.
	seen := make(map[string]bool)
	for list := range worklist {
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				unseenLinks <- link
			}
		}
	}
}

//!-
