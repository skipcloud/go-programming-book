package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/skipcloud/go-programming-book/ch4/4.12/xkcd"
)

/*
	The popular web comic xkcd has a JSON interface. For example, a request to
	https://xkcd.com/571/info.0.json produces a detailed description of comic
	571, one of many favourites. Download each URL (once!) and build an offline
	index. Write a tool xkcd that, using this index, prints the URL and transcript
	of each comic that matches a search term provided on the command line.
*/

func main() {
	if len(os.Args) == 1 {
		log.Fatal("search term missing")
		os.Exit(1)
	}
	client, err := xkcd.NewClient()
	if err != nil {
		log.Fatal(err)
	}

	client.Index.Update()
	r := client.Search(strings.Join(os.Args[1:], " "))

	fmt.Printf("\n%d comics found\n\n", len(r.Comics))
	for _, comic := range r.Comics {
		fmt.Println(comic.Title)
		fmt.Println(comic.Link)
		fmt.Printf("%s\n\n--------------\n\n", comic.Transcript)
	}
}
