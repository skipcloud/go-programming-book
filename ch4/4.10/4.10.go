package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"gobook/ch4/github"
)

/*
 * Modify issues to report the results in age categories, say less
 * than a month old, less than a year old, and more than a year old.
 */

func main() {
	const (
		lessThanMonth = "less than month old"
		lessThanYear  = "less than year old"
		moreThanYear  = "more than year"
	)
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)

	categories := make(map[string][]*github.Issue)

	now := time.Now()
	for _, item := range result.Items {
		switch {
		case item.CreatedAt.Before(now.AddDate(0, 1, 0)):
			categories[lessThanMonth] = append(categories[lessThanMonth], item)
		case item.CreatedAt.Before(now.AddDate(1, 0, 0)):
			categories[lessThanYear] = append(categories[lessThanYear], item)
		case item.CreatedAt.After(now.AddDate(1, 0, 0)):
			categories[moreThanYear] = append(categories[moreThanYear], item)
		}
	}

	fmt.Printf("\nLess than a month\n")
	for _, item := range categories[lessThanMonth] {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}

	fmt.Printf("\nLess than a year\n")
	for _, item := range categories[lessThanYear] {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}

	fmt.Printf("\nMore than a year\n")
	for _, item := range categories[moreThanYear] {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}
}
