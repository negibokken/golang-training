package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

var out io.Writer = os.Stdout

func responseOut(resp *http.Response) {
	_, err := io.Copy(out, resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: reading: %v\n", err)
		os.Exit(1)
	}
}

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		responseOut(resp)
		resp.Body.Close()
	}
}
