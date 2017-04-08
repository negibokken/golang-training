package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

var out io.Writer = os.Stdout

func responseOut(resp *http.Response) {
	_, err := io.Copy(out, resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: reading: %v\n", err)
		os.Exit(1)
	}
}

func cordinatePrefix(url string) string {
	if !strings.HasPrefix(url, "http://") {
		url = "http://" + url
	}
	return url
}

func main() {
	for _, url := range os.Args[1:] {
		url = cordinatePrefix(url)
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		responseOut(resp)
		resp.Body.Close()
	}
}
