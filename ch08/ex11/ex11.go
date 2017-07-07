package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
)

var cancel = make(chan struct{})

// Result struct
type Result struct {
	filename string
	n        int64
}

func main() {
	if len(os.Args) <= 1 {
		fmt.Fprintf(os.Stderr, "./ex11 <url>...")
		return
	}
	urls := os.Args[1:]
	fmt.Println(urls)
	mirroredFetch(urls)
}

func mirroredFetch(urls []string) {
	result := make(chan Result, len(urls))
	for _, url := range urls {
		go func(url string) {
			fileName, n, err := fetch(url)
			if err != nil {
				fmt.Fprintf(os.Stderr, "error: %v", err)
				return
			}
			result <- Result{fileName, n}
		}(url)
	}
	res := <-result
	fmt.Printf("created: %v (%v byte)", res.filename, res.n)
}

func fetch(url string) (filename string, n int64, err error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", 0, err
	}
	req.Cancel = cancel

	resp, err := http.DefaultClient.Do(req)
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
	n, err = io.Copy(f, resp.Body)
	if closeErr := f.Close(); err == nil {
		err = closeErr
	}
	return local, n, err
}
