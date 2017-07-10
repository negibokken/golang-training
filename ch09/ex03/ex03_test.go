package memo

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
	"time"
)

func httpGetBody(url string, cancel Cancel) (interface{}, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func incomingURLs() <-chan string {
	ch := make(chan string)
	urls := []string{
		"http://gopl.io",
		"http://gopl.io",
		"https://google.com",
		"http://gopl.io",
		"https://google.com",
	}
	go func() {
		for _, url := range urls {
			ch <- url
		}
		close(ch)
	}()
	return ch
}

func Test(t *testing.T) {
	done := make(chan struct{})
	m := New(httpGetBody)
	for url := range incomingURLs() {
		start := time.Now()
		value, err := m.Get(url, done)
		if err != nil {
			t.Errorf("%v", err)
		}
		fmt.Printf("%s, %s, %d bytes\n",
			url, time.Since(start), len(value.([]byte)))
	}
}
