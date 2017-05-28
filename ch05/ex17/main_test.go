package main

import (
	"strings"
	"testing"

	"golang.org/x/net/html"
)

func TestElementByTagName(t *testing.T) {
	doc, err := html.Parse(strings.NewReader(`
		<html>
			<head></head>
			<div></div>
		</html>
	`))
	if err != nil {
		t.Errorf("%v", err)
	}
	type args struct {
		doc  *html.Node
		strs []string
	}
	tests := []struct {
		name    string
		args    args
		wantTag string
	}{
		{
			"test element by tag name",
			args{doc, []string{"div"}},
			"div",
		},
		{
			"test element by tag name",
			args{doc, []string{"head"}},
			"head",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ElementByTagName(tt.args.doc, tt.args.strs...)
			if got[0].Data != tt.wantTag {
				t.Errorf("ElementByTagName() = %v, want %v", got, tt.wantTag)
			}
		})
	}
}

func Test_isTargetTag(t *testing.T) {
	type args struct {
		tag  string
		strs []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"test is target tag",
			args{"div", []string{"div", "link"}},
			true,
		},
		{
			"test is target tag",
			args{"img", []string{"div", "link"}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isTargetTag(tt.args.tag, tt.args.strs); got != tt.want {
				t.Errorf("isTargetTag() = %v, want %v", got, tt.want)
			}
		})
	}
}
