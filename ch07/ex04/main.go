package main

import (
	"fmt"
	"io"
	"os"

	"golang.org/x/net/html"
)

// MyReader type
type MyReader struct {
	s []byte
}

func (r *MyReader) Read(p []byte) (n int, err error) {
	n = copy(p, r.s)
	r.s = r.s[n:]
	if len(r.s) == 0 {
		return 0, io.EOF
	}
	return n, nil
}

// MyNewReader returns pointer of Reader
func MyNewReader(s string) io.Reader {
	return &MyReader{[]byte(s)}
}

var out io.Writer = os.Stdout

func outline(stack []string, n *html.Node) {
	if n.Type == html.ElementNode {
		stack = append(stack, n.Data) // push tag
		fmt.Fprintf(out, "%v\n", stack)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		outline(stack, c)
	}
}

func main() {
	doc, err := html.Parse(MyNewReader("<html><body><p>hello</p></body></html>"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
	}
	var stack []string
	outline(stack, doc)
}
