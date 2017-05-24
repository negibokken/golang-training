package main

import (
	"testing"

	"strings"

	"golang.org/x/net/html"
)

func TestElementByID(t *testing.T) {
	doc, err := html.Parse(strings.NewReader(`
		<html>
			<head></head>
			<body>
				<div id="foo"></div>
			</body>
		</html>
	`))
	if err != nil {
		t.Errorf("%v", err)
	}
	type args struct {
		doc *html.Node
		id  string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"test get element by id",
			args{
				doc,
				"foo",
			},
			"div",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ElementByID(tt.args.doc, tt.args.id)
			if got.Data != tt.want {
				t.Errorf("ElementByID() = %v, want %v", got.Data, tt.want)
			}
		})
	}
}

func Test_forEachNode(t *testing.T) {
	type args struct {
		n    *html.Node
		pre  func(n *html.Node) bool
		post func(n *html.Node) bool
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := forEachNode(tt.args.n, tt.args.pre, tt.args.post); got != tt.want {
				t.Errorf("forEachNode() = %v, want %v", got, tt.want)
			}
		})
	}
}
