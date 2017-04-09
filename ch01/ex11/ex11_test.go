package main

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func TestFetch(t *testing.T) {
	tests := []struct {
		url string
	}{
		{"http://gopl.io"},
		{"http://www.naist.jp"},
	}
	ch := make(chan string)
	for _, test := range tests {
		go fetch(test.url, ch)
	}
	for range tests {
		message := <-ch
		if message == "" {
			t.Errorf("Could not receive message")
		}
	}
}

func TestSaveToFile(t *testing.T) {
	tests := []struct {
		messages []string
	}{
		{
			[]string{"aaa", "bbb", "ccc"},
		},
		{
			[]string{"ccc", "ddd", "eee"},
		},
	}
	for _, test := range tests {
		err := saveToFile(test.messages)
		if err != nil {
			t.Errorf("%v", err)
		}
	}
	fileInfos, err := ioutil.ReadDir(".")
	if err != nil {
		t.Errorf("Could not get files: %v", err)
	}
	for _, fileInfo := range fileInfos {
		fileName := fileInfo.Name()
		if strings.HasSuffix(fileName, ".txt") {
			os.Remove(fileName)
		}
	}
}
