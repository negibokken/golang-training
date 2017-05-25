package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(1)
	}
	var depth int

	forEachNode(doc,
		func(n *html.Node) {
			shouldPutNewline := false
			if n.Type == html.ElementNode {
				fmt.Printf("%s", formatAttribute(n, depth, n.FirstChild == nil))
				depth++
				shouldPutNewline = true
			} else if n.Type == html.TextNode {
				if !shouldSkipped(n.Data) {
					str := strings.Replace(n.Data, "\n", "", -1)
					str = strings.Replace(str, "\t", " ", -1)
					if str == "  " {
						fmt.Printf("%x", n.Data)
					}
					fmt.Printf("%*s%s", depth*2, "", str)
					shouldPutNewline = true
				}
			}
			if shouldPutNewline {
				fmt.Printf("\n")
			}
		},
		func(n *html.Node) {
			if n.Type == html.ElementNode && n.FirstChild == nil {
				depth--
				return
			}
			if n.Type == html.ElementNode {
				depth--
				fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
			}
		})
}

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}

func formatAttribute(n *html.Node, depth int, hasNochild bool) string {
	str := fmt.Sprintf("%*s<%s", depth*2, "", n.Data)
	for _, a := range n.Attr {
		str += fmt.Sprintf(" %s='%s'", a.Key, a.Val)
	}
	if hasNochild {
		str += "/"
	}
	str += ">"
	return str
}

func shouldSkipped(str string) bool {
	isEmpty := str == ""
	isMixed := regexp.MustCompile("^[\t \n]+$").Match([]byte(str))
	return isMixed || isEmpty
}
