package main

import (
	"fmt"
	"os"

	"net/http"

	"golang.org/x/net/html"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Fprintf(os.Stderr, "./ex17 <url> <tag> <tag>...")
		os.Exit(1)
	}
	url, tags := os.Args[1], os.Args[2:len(os.Args)]

	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "url:%v %v", url, err)
		os.Exit(1)
	}

	doc, err := html.Parse(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(1)
	}

	node := ElementByTagName(doc, tags...)
	if node == nil {
		fmt.Fprintf(os.Stderr, "%v is not found %v", tags, url)
		os.Exit(1)
	}
	for _, n := range node {
		fmt.Printf("%#v\n", n)
	}
}

func ElementByTagName(doc *html.Node, strs ...string) []*html.Node {
	var results []*html.Node

	if len(strs) == 0 {
		return nil
	}
	pre := func(n *html.Node) bool {
		if n.Type != html.ElementNode {
			return true
		}

		if isTargetTag(n.Data, strs) {
			node := n
			results = append(results, node)
		}
		return true
	}

	post := func(n *html.Node) bool { return true }
	forEachNode(doc, pre, post)
	return results
}

func isTargetTag(tag string, strs []string) bool {
	for _, str := range strs {
		if tag == str {
			return true
		}
	}
	return false
}

func forEachNode(n *html.Node, pre, post func(n *html.Node) bool) bool {
	if pre != nil {
		if !pre(n) {
			return false
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if !forEachNode(c, pre, post) {
			return false
		}
	}
	if post != nil {
		if !post(n) {
			return false
		}
	}
	return true
}
