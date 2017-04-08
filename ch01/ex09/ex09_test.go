package main

import (
	"bytes"
	"net/http"
	"testing"
)

func TestHeaderOut(t *testing.T) {
	var tests = []struct {
		header   http.Header
		expected string
	}{
		{
			http.Header{
				"Date":    []string{"Sat, 08 Apr 2017 14:49:21 GMT"},
				"X-Timer": []string{"S1491662961.184274,VS0,VE0"},
			},
			"Date: [Sat, 08 Apr 2017 14:49:21 GMT]\nX-Timer: [S1491662961.184274,VS0,VE0]\n",
		},
		{
			http.Header{
				"Vary": []string{"Accept-Encoding"},
			},
			"Vary: [Accept-Encoding]\n",
		},
		{
			http.Header{},
			"",
		},
	}
	for _, test := range tests {
		out = new(bytes.Buffer)
		headerOut(test.header)
		got := out.(*bytes.Buffer).String()
		if got != test.expected {
			t.Errorf("Actual: %v Expected: %v is not much", got, test.expected)
		}
	}
}

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
