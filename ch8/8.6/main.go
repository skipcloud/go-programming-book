package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"gopl.io/ch5/links"
)

/*
	Add depth-limiting to the concurrent crawler. That is, if the user sets -depth=3,
	then only URLs reachable by at most three links will be fetched.
*/

// tokens is a counting semaphore used to
// enforce a limit of 20 concurrent requests.
var tokens = make(chan struct{}, 20)

// Links represents a workload of links to crawl and holds
// URLs  and the current depth for the URLs in relation
// to the root URL
type Links struct {
	urls  []string
	depth int
}

func crawl(url string, depth int) *Links {
	fmt.Printf("%s%s\n", strings.Repeat("\t", depth-1), url)
	tokens <- struct{}{} // acquire a token
	list, err := links.Extract(url)
	<-tokens // release the token

	if err != nil {
		log.Print(err)
	}
	return &Links{
		urls:  list,
		depth: depth,
	}
}

func main() {
	depth := flag.Int("depth", 10, "the depth that the crawler will go")
	flag.Parse()
	// I know there is only ever going to be one option -depth
	// but this loop will move the args along until we get to
	// the URL args.
	for len(os.Args) > 1 && strings.HasPrefix(os.Args[1], "-") {
		os.Args = os.Args[1:]
	}
	worklist := make(chan *Links)
	var n int // number of pending sends to worklist

	// Start with the command-line arguments.
	n++
	go func() {
		worklist <- &Links{
			urls:  os.Args[1:],
			depth: 0,
		}
	}()

	// Crawl the web concurrently.
	seen := make(map[string]bool)
	for ; n > 0; n-- {
		list := <-worklist
		for _, link := range list.urls {
			if !seen[link] {
				seen[link] = true
				if list.depth < *depth {
					n++
					go func(link string) {
						worklist <- crawl(link, list.depth+1)
					}(link)
				}

			}
		}
	}
}

//!-
