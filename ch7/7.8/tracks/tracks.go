package tracks

import (
	"sort"
	"time"
)

/*
	Track code
*/

type track struct {
	Artist string
	Title  string
	Album  string
	Year   int
	Length time.Duration
}

type Tracks struct {
	Songs         []*track
	PrimarySort   SortKey
	SecondarySort SortKey
	TertiarySort  SortKey
}

func SomeTracks() Tracks {
	return Tracks{Songs: []*track{
		{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
		{"Go", "Moby", "Moby", 1992, length("3m37s")},
		{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
		{"Grow up", "Na na na", "lazy", 1992, length("2m37s")},
		{"Go", "Moby", "Look another album", 1992, length("5m37s")},
		{"Gooogler", "Oh man", "Album II", 1984, length("5m37s")},
		{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
		{"Ready 2 Go", "Made Up Musician", "Smash", 2001, length("4m24s")},
		{"Made up song", "Real Musician", "You like?", 1962, length("2m41s")},
	}}
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

func (t Tracks) Sort(newPrimarySort SortKey) {
	t.TertiarySort = t.SecondarySort
	t.SecondarySort = t.PrimarySort
	t.PrimarySort = newPrimarySort

	for _, s := range []SortKey{t.PrimarySort, t.SecondarySort, t.TertiarySort} {
		if s != Unset {
			sort.Sort(customSort{t.Songs, sortMethod(s)})
		}
	}
}

/*
	Sort code
*/
type customSort struct {
	t    []*track
	less func(x, y *track) bool
}

// Len is the number of elements in the collection.
func (c customSort) Len() int {
	return len(c.t)
}

// Less reports whether the element with
// index i should sort before the element with index j.
func (c customSort) Less(i int, j int) bool {
	if c.less == nil {
		return false
	}
	return c.less(c.t[i], c.t[j])
}

// Swap swaps the elements with indexes i and j.
func (c customSort) Swap(i int, j int) {
	c.t[i], c.t[j] = c.t[j], c.t[i]
}

type SortKey int

const (
	Unset SortKey = iota
	SortByAlbum
	SortByArtist
	SortByLength
	SortByTitle
	SortByYear
)

func sortMethod(key SortKey) func(x, y *track) bool {
	switch key {
	case SortByArtist:
		return byArtist
	case SortByAlbum:
		return byAlbum
	case SortByLength:
		return byLength
	case SortByTitle:
		return byTitle
	case SortByYear:
		return byYear
	}
	return nil
}

func byAlbum(x, y *track) bool {
	if x.Album != y.Album {
		return x.Album < y.Album
	}
	return false
}

func byArtist(x, y *track) bool {
	if x.Artist != y.Artist {
		return x.Artist < y.Artist
	}
	return false
}

func byLength(x, y *track) bool {
	if x.Length != y.Length {
		return x.Length < y.Length
	}
	return false
}

func byTitle(x, y *track) bool {
	if x.Title != y.Title {
		return x.Title < y.Title
	}
	return false
}

func byYear(x, y *track) bool {
	if x.Year != y.Year {
		return x.Year < y.Year

	}
	return false
}
