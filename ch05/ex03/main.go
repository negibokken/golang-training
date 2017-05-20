package main

import (
	"fmt"
	"io"
	"os"

	"strings"

	"golang.org/x/net/html"
)

var out io.Writer = os.Stdout

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "mapelment: %v\n", err)
		os.Exit(1)
	}
	visit(doc)
}

func visit(n *html.Node) {
	// Return if script or style
	if n.Type == html.ElementNode {
		if n.Data == "script" || n.Data == "style" {
			return
		}
	}
	// Search reucursively
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		visit(c)
	}
	// Print content of TextNode
	if n.Type == html.TextNode {
		str := strings.Trim(n.Data, "\t\n ")
		if str != "" {
			fmt.Fprintf(out, "%s\n", str)
		}
	}
}
