package main

import (
	"fmt"
	"os"
	"sort"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "mapelment: %v\n", err)
		os.Exit(1)
	}
	elemap := make(map[string]int)
	visit(elemap, doc)
	var str []string
	for k := range elemap {
		str = append(str, k)
	}
	sort.Strings(str)
	for _, s := range str {
		fmt.Printf("%v: %v\n", s, elemap[s])
	}
}

func visit(elemap map[string]int, n *html.Node) {
	if n.Type == html.ElementNode {
		elemap[n.Data]++
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		visit(elemap, c)
	}
}
