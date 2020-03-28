package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/skipcloud/go-programming-book/ch7/7.8/tracks"
)

/*
	Use the html/template package (section 4.6) to replace printTracks with
	a function that displays the tracks as an HTML table. Use the solution
	to the previous exercise to arrange that each click on a column heads
	makes an HTTP request to sort the table
*/

var myTracks = tracks.SomeTracks()

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		trackTemplate.Execute(w, myTracks)
	})
	http.HandleFunc("/sort", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		sort(r.Form.Get("by"))
		trackTemplate.Execute(w, myTracks)
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}

var trackTemplate = template.Must(template.New("Tracks").
	Parse(`
<!DOCTYPE html>
<html>
<head>
	<meta charset="utf-8" />
</head>
<body>
<table>
	<tr>
		<th><a href="http://localhost:8080/sort?by=title">Title</a></th>
		<th><a href="http://localhost:8080/sort?by=artist">Artist</th>
		<th><a href="http://localhost:8080/sort?by=album">Album</th>
		<th><a href="http://localhost:8080/sort?by=year">Year</th>
		<th><a href="http://localhost:8080/sort?by=length">Length</th>
	</tr>
	{{ range .Songs }}
		<tr>
			<th>{{.Title}}</th>
			<th>{{.Artist}}</th>
			<th>{{.Album}}</th>
			<th>{{.Year}}</th>
			<th>{{.Length}}</th>
		</tr>
	{{ end }}
</table>
</body>
</html>
`))

func sort(s string) {
	switch s {
	case "artist":
		myTracks.Sort(tracks.SortByArtist)
	case "album":
		myTracks.Sort(tracks.SortByAlbum)
	case "title":
		myTracks.Sort(tracks.SortByTitle)
	case "year":
		myTracks.Sort(tracks.SortByYear)
	case "length":
		myTracks.Sort(tracks.SortByLength)
	}
}
