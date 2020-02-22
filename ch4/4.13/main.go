package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

/*
	The JSON-based web service of the Open Movie Database lets you search https://omdapi.com/
	for a movie by name and down its poster image. Write a tool `poster` that downloads the
	poster image for the movie named on the command line
*/

// the website requires the use of API keys, I know having API keys in your
// code is a stupid idea, the right way to go about this is use an environment
// variable but this is just a learning exercise so I don't mind having it here.
const APIKey = "cd793416"

// url paramters:
// apikey = api key
// t = title
const APIurl = "https://omdbapi.com/?apikey=%s&t=%s"

type Movie struct {
	Actors     string
	Awards     string
	BoxOffice  string
	Country    string
	Director   string
	DVD        string
	Genre      string
	ImdbID     string `json:"imdbID"`
	ImdbRating string `json:"imdbRating"`
	ImdbVotes  string `json:"imdbVotes"`
	Language   string
	Metascore  string
	Plot       string
	Poster     string
	Production string
	Rated      string
	Ratings    []Rating
	Released   string
	Response   string
	Runtime    string
	Title      string
	Type       string
	Website    string
	Writer     string
	Year       string
}

type Rating struct {
	Source string
	Value  string
}

func main() {
	if len(os.Args) == 1 {
		fmt.Fprintf(os.Stderr, "missing search query\n")
		os.Exit(1)
	}

	title := strings.Join(os.Args[1:], " ")
	res, err := http.Get(fmt.Sprintf(APIurl, APIKey, url.QueryEscape(title)))
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		fmt.Fprintf(os.Stderr, "Unable to find '%s'\n", title)
		os.Exit(1)
	}
	m, err := parseResponse(res.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}

	fmt.Printf("'%s' poster\n\n%s\n", title, m.Poster)
}

func parseResponse(res io.ReadCloser) (*Movie, error) {
	movie := Movie{}
	if err := json.NewDecoder(res).Decode(&movie); err != nil {
		return nil, err
	}

	return &movie, nil
}
