package main

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/skipcloud/go-programming-book/ch7/7.8/tracks"
)

/*
	Many GUIs provide a table widget with a stateful multi-tier sort: the
	primary key is the most recently clicked column head, the secondary sort
	key is the second-most recently clicked column head, and so on. Define
	an implementation of sort.Interface for use by such a table. Compare
	that approach with repeated sorting using sort.Stable
*/

var myTracks = tracks.SomeTracks()

func main() {
	fmt.Printf("No sorting\n\n")
	printTracks(myTracks)

	myTracks.Sort(tracks.SortByAlbum)
	fmt.Printf("Sorted by Album\n\n")
	printTracks(myTracks)

	myTracks.Sort(tracks.SortByArtist)
	fmt.Printf("Sorted by Artist first then Album\n\n")
	printTracks(myTracks)

	myTracks.Sort(tracks.SortByYear)
	fmt.Printf("Sorted by Year first then Artist then Album\n\n")
	printTracks(myTracks)

	myTracks.Sort(tracks.SortByLength)
	fmt.Printf("Sorted by Length first then Year then Artist\n\n")
	printTracks(myTracks)

	myTracks.Sort(tracks.SortByTitle)
	fmt.Printf("Sorted by Title first then Length then Year\n\n")
	printTracks(myTracks)
}

func printTracks(t tracks.Tracks) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")
	for _, s := range t.Songs {
		fmt.Fprintf(tw, format, s.Title, s.Artist, s.Album, s.Year, s.Length)
	}
	fmt.Fprintf(tw, "\n\n")
	tw.Flush()
}
