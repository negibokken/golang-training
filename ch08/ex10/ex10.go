package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {

	ctx, cancelFunc := context.WithCancel(context.Background())

	req, err := http.NewRequest("GET", "http://example.com", nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(1)
	}
	req = req.WithContext(ctx)

	fmt.Println("Press any key to cancel")
	result := make(chan string)
	go func() {
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v", err)
			os.Exit(1)
		}
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v", err)
			os.Exit(1)
		}
		result <- string(body)
	}()

	go func() {
		os.Stdin.Read(make([]byte, 1))
		cancelFunc()
	}()

	fmt.Println(<-result)
}
