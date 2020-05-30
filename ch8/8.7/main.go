package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strings"
	"sync"

	"github.com/skipcloud/go-programming-book/ch8/8.7/links"
	"golang.org/x/net/html"
)

/*
	Write a concurrent prgram that creates a local mirror of a wabsite, fetching
	each reachable page and writing it to a directory on the local disk. Only pages
	within the original domain (for instance golang.org) should be fetched. URLs
	within mirrored pages should be altered as needed so that they refer to the
	mirrored page, not the original.

	note from Skip: this is by far a perfect implementation, for example an import
					statement in a CSS file won't be followed so CSS files might be
					missing. That being said I think covered most of the bases.
*/

// The default directory that the website will be saved in
var dir = "/tmp/sites"

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("missing argument: %s url [directory]", os.Args[0])
	}
	// create a channel for the worklist, aka a list of urls to visit
	worklist := make(chan []string)
	// n is the number of items we need to work through
	n := 1
	go func() { worklist <- []string{os.Args[1]} }()

	// Parse the main url, we need to know the details
	// of it so we only visit relevant urls
	base, err := url.Parse(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	// The directory to save the website in
	if len(os.Args) == 3 {
		dir = os.Args[2]
	} else {
		dir = fmt.Sprintf("%s/%s", dir, base.Hostname())
	}

	var wg sync.WaitGroup
	// tokens to help keep concurrency in a sensible range,
	// goroutines will set a token before making any requests
	tokens := make(chan struct{}, 20)
	// seen is a map of urls we have already visited, to stop us
	// get into a loop
	seen := map[string]bool{}
	for ; n > 0; n-- {
		list := <-worklist
		for _, l := range list {
			pl, err := url.Parse(l)
			if err != nil {
				log.Fatal(err)
			}
			if !pl.IsAbs() {
				pl = base.ResolveReference(pl)

			}

			if pl.Hostname() == base.Hostname() && !seen[pl.String()] {
				seen[pl.String()] = true
				wg.Add(1)
				n++
				go func(pl *url.URL) {
					defer wg.Done()

					tokens <- struct{}{}
					body, err := links.UrlToHTMLNode(pl.String())
					if err != nil {
						log.Fatal(err)
					}
					worklist <- links.Extract(body)
					err = updatePageLinks(body, base)
					if err != nil {
						log.Fatal(err)
					}

					err = savePage(body, pl)
					if err != nil {
						log.Fatal(err)
					}

					<-tokens
				}(pl)
			}
		}
	}
	wg.Wait()
	fmt.Println("done")
}

func urlToFilePath(path string) string {
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}
	return fmt.Sprintf("%s%s", dir, path)
}

func savePage(page *html.Node, u *url.URL) error {
	// first we need to create the directories
	dirs := strings.Split(u.Path, "/")
	file := ensureExtension(dirs[len(dirs)-1])
	path := urlToFilePath(strings.Join(dirs[:len(dirs)-1], "/"))
	// if the dir doesn't exist then create the path
	stat, err := os.Stat(path)
	if os.IsNotExist(err) || !stat.IsDir() {
		err = os.MkdirAll(path, 0777)
		if err != nil {
			return err
		}
	}
	if strings.Contains(file, "all") {
		println(fmt.Sprintf("%s/%s", path, file))
	}

	// then we create the file
	f, err := os.Create(fmt.Sprintf("%s/%s", path, file))
	if err != nil {
		return err
	}
	defer f.Close()
	// then we save the file
	if isHTML(file) {
		err = html.Render(f, page)
	} else {
		// if not html then download the actual file,
		// i.e. images, css, javscript etc
		err = downloadAndSaveFile(f, u)
	}
	return err
}

func downloadAndSaveFile(f *os.File, u *url.URL) error {
	resp, err := http.Get(u.String())
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	_, err = io.Copy(f, resp.Body)
	return err
}

func isHTML(file string) bool {
	rgx := regexp.MustCompile("\\.html$")
	return rgx.Match([]byte(file))
}

func ensureExtension(f string) string {
	if f == "" {
		return "index.html"
	}
	rgx := regexp.MustCompile("\\..*$")
	if !rgx.Match([]byte(f)) {
		return fmt.Sprintf("%s.html", f)
	}
	return f
}

func updatePageLinks(page *html.Node, base *url.URL) error {
	var u *url.URL
	var err error
	if links.IsLinkHavingNode(page) {
		for i, a := range page.Attr {
			if a.Key == "href" || a.Key == "src" {
				u, err = url.Parse(a.Val)
				if err != nil {
					return err
				}
				if u.Hostname() == base.Hostname() || !u.IsAbs() {
					if u.Path != "" {
						page.Attr[i].Val = urlToFilePath(u.Path)
					} else {
						println("in here")
						page.Attr[i].Val = urlToFilePath("/index.html")
					}
				}
			}
		}

	}
	for node := page.FirstChild; node != nil; node = node.NextSibling {
		if err = updatePageLinks(node, base); err != nil {
			return err
		}
	}
	return nil
}
