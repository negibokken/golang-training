package main

import (
	"testing"

	"strings"

	"os"

	"bytes"

	"golang.org/x/net/html"
)

func TestVisit(t *testing.T) {
	doc, err := html.Parse(strings.NewReader(`
		<html>
			<head>
				<script>alert(1);</script>
			</head>
			<body>
				<p>hello<p>
				<p>world</p>
			</body>
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
			"print web except for script and style",
			args{doc},
			"hello\nworld\n",
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
