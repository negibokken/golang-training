package main

import "fmt"
import "os"

var prereqs = map[string]map[string]bool{
	"algorithms":     map[string]bool{"data structures": true},
	"calculus":       map[string]bool{"linear algebra": true},
	"linear algebra": map[string]bool{"calculus": true},
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
	cources, err := topoSort(prereqs)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
	}
	for i, course := range cources {
		fmt.Printf("%d: \t%s\n", i+1, course)
	}
}

func topoSort(m map[string]map[string]bool) ([]string, error) {
	var order []string
	seen := make(map[string]bool)
	checkList := make(map[string]bool)

	var visitAll func(items map[string]bool)
	cyclic := false
	visitAll = func(items map[string]bool) {
		if cyclic {
			return
		}
		for item := range items {
			if checkList[item] {
				cyclic = true
				return
			}
			if !seen[item] {
				seen[item] = true

				checkList[item] = true
				visitAll(m[item])
				checkList[item] = false

				order = append(order, item)
			}
		}
	}

	for item := range m {
		visitAll(map[string]bool{item: true})
		if cyclic {
			return nil, fmt.Errorf("the cources are cyclic")
		}
	}

	return order, nil
}
