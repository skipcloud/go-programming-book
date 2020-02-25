package main

import (
	"fmt"
)

/*
	Rewrite topoSort to use maps instead of slices and eliminate the initial sort.
	Verify that the results, though nondeterministic, are valid topological orderings.
*/

// prereqs maps computer science courses to their prerequisites.
var prereqs = map[string]map[string]string{
	"algorithms": {"data structures": ""},
	"calculus":   {"linear algebra": ""},

	"compilers": {
		"data structures":       "",
		"formal languages":      "",
		"computer organization": "",
	},

	"data structures":       {"discrete math": ""},
	"databases":             {"data structures": ""},
	"discrete math":         {"intro to programming": ""},
	"formal languages":      {"discrete math": ""},
	"networks":              {"operating systems": ""},
	"operating systems":     {"data structures": "", "computer organization": ""},
	"programming languages": {"data structures": "", "computer organization": ""},
}

func main() {
	for i, course := range topoSort(prereqs) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

func topoSort(m map[string]map[string]string) []string {
	var order []string
	seen := make(map[string]bool)
	var visitAll func(items map[string]string)

	visitAll = func(items map[string]string) {
		for k, _ := range items {
			if !seen[k] {
				seen[k] = true
				visitAll(m[k])
				order = append(order, k)
			}
		}
	}

	keys := map[string]string{}
	for key := range m {
		keys[key] = ""
	}

	// sort.Strings(keys)
	visitAll(keys)
	return order
}
