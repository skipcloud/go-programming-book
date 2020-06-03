package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

/*
	Following the approach of mirroredQuery in Section 8.4.4, implement a variant
	of fetch that requests several URLs concurrently. As soon as the first response
	arrives, cancel the other requests.
*/

func main() {
	if len(os.Args) == 1 {
		fmt.Fprintf(os.Stderr, "missing arguments")
		os.Exit(1)
	}
	resp := concurrentFetch(os.Args[1:])
	fmt.Println(resp)
}

func concurrentFetch(urls []string) string {
	cancel := make(chan struct{})
	response := make(chan string)
	for _, url := range urls {
		go func(u string) {
			select {
			case <-cancel:
				break
			default:
				response <- fetch(u, cancel)
			}
		}(url)
	}
	resp := <-response
	close(cancel)
	return resp

}

func fetch(u string, cancel chan struct{}) string {
	url, err := url.Parse(u)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: cannot parse url '%s': %v\n", u, err)
		return ""
	}

	req := http.Request{
		URL:    url,
		Method: "GET",
		Cancel: cancel,
	}
	res, err := http.DefaultClient.Do(&req)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: problem fetching url '%s': %v\n", u, err)
		return ""
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: problem reading response '%s': %v\n", u, err)
		return ""
	}

	return string(body)
}
