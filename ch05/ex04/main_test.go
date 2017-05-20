package main

import (
	"os"
	"testing"

	"strings"

	"bytes"

	"golang.org/x/net/html"
)

func Test_visit(t *testing.T) {
	doc, err := html.Parse(strings.NewReader(`
		<html>
			<head>
				<link rel="stylesheet" href="css.css">
				<script src="javascript.js"></script>
			</head>
			<img src="image.jpg">
		</html>
	`))
	if err != nil {
		t.Errorf("%v", err)
		os.Exit(1)
	}
	type args struct {
		n *html.Node
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"link list",
			args{doc},
			"css.css\njavascript.js\nimage.jpg\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out = new(bytes.Buffer)
			visit(tt.args.n)
			got := out.(*bytes.Buffer).String()
			if got != tt.want {
				t.Errorf("%x, want %x", got, tt.want)
			}
		})
	}
}
