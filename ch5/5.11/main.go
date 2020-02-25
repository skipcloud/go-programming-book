// The toposort program prints the nodes of a DAG in topological order.
package main

import (
	"fmt"
	"sort"
)

/*
	The instructor of the linear algebra course decides that calculus is now a
	prerequisite. Extend the topoSort function to report cycles
*/

// prereqs maps computer science courses to their prerequisites.
var prereqs = map[string][]string{
	"linear algebra": {"calculus"},
	"algorithms":     {"data structures"},
	"calculus":       {"linear algebra"},

	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},

	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

func main() {
	for i, course := range topoSort(prereqs) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

func topoSort(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)
	var visitAll func(items []string)

	visitAll = func(items []string) {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				for _, req := range m[item] {
					if contains(m[req], item) {
						fmt.Printf("Cycle found for class %s\n", item)
					}
				}
				visitAll(m[item])
				order = append(order, item)
			}
		}
	}

	var keys []string
	for key := range m {
		keys = append(keys, key)
	}

	sort.Strings(keys)
	visitAll(keys)
	return order
}

func contains(reqs []string, class string) bool {
	for _, v := range reqs {
		if v == class {
			return true
		}
	}
	return false
}
