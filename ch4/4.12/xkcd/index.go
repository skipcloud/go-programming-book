package xkcd

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os/user"
)

const xkcdUrl = "https://xkcd.com/%d/info.0.json"
const indexFileName = ".xkcdindex"

var IndexFilePath string

type index struct {
	Comics         []*comic `json:"comics"`
	indexLoaded    bool
	filePath       string
	existingComics map[int]struct{}
}

type comic struct {
	Month      string `json:"month"`
	ID         int    `json:"num"`
	Link       string `json:"link"`
	Year       string `json:"year"`
	News       string `json:"news"`
	SafeTitle  string `json:"safe_title"`
	Transcript string `json:"transcript"`
	Alt        string `json:"alt"`
	Image      string `json:"img"`
	Title      string `json:"title"`
	Day        string `json:"day"`
}

func loadIndex() (*index, error) {
	index := index{}

	// set up file path
	user, _ := user.Current()
	if user == nil {
		index.filePath = "./" + indexFileName
	} else {
		index.filePath = user.HomeDir + "/" + indexFileName
	}

	// try to read index file if it exists
	if err := index.readFile(); err != nil {
		return nil, err
	}

	// populate already fetched comics
	index.alreadyFetchedComics()
	return &index, nil
}

func (i *index) readFile() error {
	// read the file
	b, err := ioutil.ReadFile(i.filePath)
	if err != nil {
		return err
	}

	// Unmarshall index file into a struct if there
	// is content in the file
	if len(b) > 0 {
		err = json.Unmarshal(b, i)
		if err != nil {
			return err
		}
	}
	return nil
}

func (i *index) writeFile() error {
	b, err := json.Marshal(i)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(i.filePath, b, 0644)
	if err != nil {
		return err
	}
	return nil
}

func (i *index) Update() error {
	// popular index with new comics
	err := i.fetchComics()
	if err != nil {
		return err
	}
	i.writeFile()
	return nil
}

func (i *index) fetchComics() error {
	// no need to create these during each loop so
	// define them here and reuse.
	var res *http.Response
	var comic *comic
	var err error

	id := 1
	fmt.Printf("There are %d comics in the index\n", len(i.Comics))
	fmt.Println("Attempting to fetch new comics.")

	for {
		if id == 404 {
			// of course comic 404 doesn't exist
			id += 1
			continue
		}
		if _, ok := i.existingComics[id]; ok {
			// comic already in index so skip it
			id += 1
			continue
		}

		fmt.Printf("Fetching comic %d...", id)
		res, err = http.Get(fmt.Sprintf(xkcdUrl, id))
		if err != nil {
			return err
		}
		if res.StatusCode == http.StatusNotFound {
			fmt.Printf("not found\n")
			break
		}

		// unmarshall into a struct
		comic, err = parseComicReponse(res.Body)
		if err != nil {
			return err
		}

		// add comic to the index
		i.Comics = append(i.Comics, comic)
		fmt.Printf("done!\n")
		id += 1
	}
	fmt.Println("Finished")
	return nil
}

func (i *index) alreadyFetchedComics() {
	i.existingComics = make(map[int]struct{})
	for _, c := range i.Comics {
		i.existingComics[c.ID] = struct{}{}
	}
}

func parseComicReponse(body io.ReadCloser) (*comic, error) {
	var comic comic
	if err := json.NewDecoder(body).Decode(&comic); err != nil {
		return nil, err
	}
	return &comic, nil
}
