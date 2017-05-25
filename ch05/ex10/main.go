package main

import "fmt"

var prereqs = map[string]map[string]bool{
	"algorithms": map[string]bool{"data structures": true},
	"calculus":   map[string]bool{"linear algebra": true},

	"compilers": map[string]bool{
		"data structures":       true,
		"formal languages":      true,
		"computer organization": true,
	},

	"data structures":       map[string]bool{"discrete math": true},
	"databases":             map[string]bool{"data structures": true},
	"discrete math":         map[string]bool{"intro to programming": true},
	"formal languages":      map[string]bool{"discrete math": true},
	"networks":              map[string]bool{"operating systems": true},
	"operating systems":     map[string]bool{"data structures": true, "computer organization": true},
	"programming languages": map[string]bool{"data structures": true, "computer organization": true},
}

func main() {
	for i, course := range topoSort(prereqs) {
		fmt.Printf("%d: \t%s\n", i+1, course)
	}
}

func topoSort(m map[string]map[string]bool) []string {
	var order []string
	seen := make(map[string]bool)

	var visitAll func(items map[string]bool)
	visitAll = func(items map[string]bool) {
		for item := range items {
			if !seen[item] {
				seen[item] = true
				visitAll(m[item])
				order = append(order, item)
			}
		}
	}

	for item := range m {
		visitAll(map[string]bool{item: true})
	}
	return order
}
