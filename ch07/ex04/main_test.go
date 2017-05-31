package main

import (
	"bytes"
	"testing"

	"golang.org/x/net/html"
)

func TestMyNewReader(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"test my new reader",
			args{"<html><head></head><body></body></html>"},
			"[html]\n[html head]\n[html body]\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out = new(bytes.Buffer)
			doc, err := html.Parse(MyNewReader(tt.args.s))
			if err != nil {
				t.Errorf("%v", err)
			}
			var stack []string
			outline(stack, doc)
			if got := out.(*bytes.Buffer).String(); got != tt.want {
				t.Errorf("MyNewReader() = %v, want %v", got, tt.want)
			}
		})
	}
}
