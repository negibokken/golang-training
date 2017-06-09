package main

import (
	"fmt"
	"os"

	"net/http"

	"golang.org/x/net/html"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "./ex08 <url> <id>")
	}
	url, id := os.Args[1], os.Args[2]

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

	node := ElementByID(doc, id)
	if node == nil {
		fmt.Fprintf(os.Stderr, "%v is not found %v", id, url)
		os.Exit(1)
	}
	fmt.Printf("%v", node)
}

func ElementByID(doc *html.Node, id string) *html.Node {
	var result *html.Node

	pre := func(n *html.Node) bool {
		if n.Type != html.ElementNode {
			return true
		}

		for _, a := range n.Attr {
			if a.Key == "id" && a.Val == id {
				result = n
				return false
			}
		}
		return true
	}

	// post := func(n *html.Node) bool { return true }
	// forEachNode(doc, pre, post)
	forEachNode(doc, pre, nil)
	return result
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
