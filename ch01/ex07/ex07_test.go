package main

import (
	"bytes"
	"net/http"
	"testing"
)

func TestResponseOut(t *testing.T) {
	args := []string{
		"http://gopl.io",
	}
	for _, url := range args {
		resp, err := http.Get(url)
		if err != nil {
			t.Errorf("fetch: %v\n", err)
		}
		out = new(bytes.Buffer)
		responseOut(resp)
		got := out.(*bytes.Buffer).String()
		if got == "" {
			t.Errorf("response is nil. Something is wrong")
		}
		resp.Body.Close()
	}
}
