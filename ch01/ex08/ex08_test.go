package main

import (
	"bytes"
	"net/http"
	"testing"
)

func TestCordinatePrefix(t *testing.T) {
	var tests = []struct {
		url      string
		expected string
	}{
		{
			"http://gopl.io",
			"http://gopl.io",
		},
		{
			"",
			"http://",
		},
		{
			"gopl.io",
			"http://gopl.io",
		},
	}
	for _, test := range tests {
		if actual := cordinatePrefix(test.url); test.expected != actual {
			t.Errorf("Expected: %v is not equal to Actual: %v", test.expected, actual)
		}
	}
}

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
