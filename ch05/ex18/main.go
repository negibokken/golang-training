package main

import (
	"fmt"
	"io"
	"os"
	"path"

	"net/http"
)

func main() {
	if os.Args[1] == "" {
		fmt.Fprintf(os.Stderr, "./ex18 <url>")
		os.Exit(1)
	}
	filename, n, err := fetch(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(1)
	}
	fmt.Printf("%s(%d byte) is generated\n", filename, n)
}

func fetch(url string) (filename string, n int64, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()

	local := path.Base(resp.Request.URL.Path)
	if local == "/" {
		local = "index.html"
	}
	f, err := os.Create(local)
	if err != nil {
		return "", 0, err
	}
	defer f.Close()
	n, err = io.Copy(f, resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
	}
	return local, n, err
}
