package links

import (
	"net/http"

	"golang.org/x/net/html"
)

// Extract extracts links from a parsed HTML document
func Extract(body *html.Node) []string {
	links := []string{}
	links = walkChildren(body, links)
	return links
}

// UrlToHTMLNode changes a URL string into an html.Node
func UrlToHTMLNode(url string) (*html.Node, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	body, err := html.Parse(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

// walkChildren walks the tree depth first pulling out the links in each node
func walkChildren(node *html.Node, links []string) []string {
	for node := node.FirstChild; node != nil; node = node.NextSibling {
		if node.FirstChild != nil {
			links = walkChildren(node, links)
		}
		links = append(links, extractLinks(node)...)
	}
	return links
}

func extractLinks(node *html.Node) []string {
	links := []string{}
	if IsLinkHavingNode(node) {
		for _, v := range node.Attr {
			if v.Key == "href" || v.Key == "src" {
				links = append(links, v.Val)
			}
		}
	}
	return links
}

func IsLinkHavingNode(node *html.Node) bool {
	return node.Type == html.ElementNode && (node.Data == "a" || node.Data == "link" || node.Data == "img")
}
