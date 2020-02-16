package xkcd

import "strings"

type xkcd struct {
	Index *index
}

func NewClient() (*xkcd, error) {
	i, err := loadIndex()
	if err != nil {
		return nil, err
	}

	return &xkcd{Index: i}, nil
}

type searchResult struct {
	Comics []*comic
}

func (c *xkcd) Search(term string) *searchResult {
	results := searchResult{}
	for _, comic := range c.Index.Comics {
		if strings.Contains(strings.ToLower(comic.Transcript), strings.ToLower(term)) {
			results.Comics = append(results.Comics, comic)
		}
	}
	return &results
}
