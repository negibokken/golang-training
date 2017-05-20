package main

import (
	"fmt"
	"io"
	"os"

	"golang.org/x/net/html"
)

var out io.Writer = os.Stdout

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks: %v\n", err)
		os.Exit(1)
	}
	visit(doc)
}

func visit(n *html.Node) {
	if n.Type == html.ElementNode {
		if n.Data == "script" || n.Data == "link" || n.Data == "img" {
			for _, a := range n.Attr {
				if a.Key == "href" || a.Key == "src" {
					fmt.Fprintf(out, "%v\n", a.Val)
				}
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		visit(c)
	}
}
