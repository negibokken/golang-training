package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"bufio"

	"strings"

	"golang.org/x/net/html"
)

var out io.Writer = os.Stdout

func main() {
	if os.Args[1] == "" {
		fmt.Printf("%s <url>\n", os.Args[0])
		os.Exit(1)
	}
	url := os.Args[1]
	words, images, err := CountWordsAndImages(url)
	if err != nil {
		fmt.Printf("%v", err)
		os.Exit(1)
	}
	fmt.Printf("%v contains of %d words, %d images\n", url, words, images)
}

// CountWordsAndImages count words and images
func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
		return
	}
	words, images = countWordsAndImages(doc)
	return
}

func countWordsAndImages(n *html.Node) (words, images int) {
	if n.Type == html.ElementNode {
		if n.Data == "script" || n.Data == "style" {
			return
		}
		if n.Data == "img" {
			images++
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		w, i := countWordsAndImages(c)
		words += w
		images += i
	}

	if n.Type == html.TextNode {
		in := bufio.NewScanner(strings.NewReader(n.Data))
		in.Split(bufio.ScanWords)
		for in.Scan() {
			words++
		}
	}
	return words, images
}
