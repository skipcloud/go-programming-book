package github2

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

const (
	IssuesURL      = "https://api.github.com/search/issues"      // GET
	CreateIssueURL = "https://api.github.com/repos/%s/%s/issues" // POST
)

type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string    // in markdown format
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

func SearchIssues(terms []string) (*IssuesSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	resp, err := http.Get(IssuesURL + "?q=" + q)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	var result IssuesSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

type CreateBody struct {
	Title []byte
	Body  []byte
}

func CreateIssue(owner, repo string) string {
	url := fmt.Sprintf(CreateIssueURL, owner, repo)
	requestBody := buildCreateBody()

	marshaledBody, err := json.Marshal(requestBody)
	if err != nil {
		log.Fatalf("create issue: error marshaling json - %v\n", err)
	}
	resp, err := http.Post(url, "application/vnd.github.symmetra-preview+json", bytes.NewBuffer(marshaledBody))
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("create issue: response from github API %d\n", resp.StatusCode)
	}
	fmt.Println(resp.Body)
	return url
}

func buildCreateBody() *CreateBody {
	var title, body []byte

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Title of your issue: ")
	scanner.Scan()
	title = scanner.Bytes()

	fmt.Printf("Do you require a text editor to enter the body of your issue? [y/n] ")
	scanner.Scan()
	if scanner.Text() == "y" {

	} else {
		fmt.Printf("Enter the body of your issue: ")
		scanner.Scan()
		body = scanner.Bytes()
	}
	return &CreateBody{title, body}
}
