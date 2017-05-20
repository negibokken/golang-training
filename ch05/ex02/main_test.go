package main

import (
	"testing"

	"strings"

	"golang.org/x/net/html"
)

func TestVisit(t *testing.T) {
	var elemap = make(map[string]int)

	doc, err := html.Parse(strings.NewReader(`
				<html>
					<head></head>
					<body>
						<a href="https://example.com"></a>
						<a href="https://example.co.jp"></a>
						<a href="https://example.org"></a>
					</body>
				</html>
			`))
	if err != nil {
		t.Errorf("%v", err)
	}
	type args struct {
		elemap map[string]int
		n      *html.Node
	}
	tests := []struct {
		name     string
		args     args
		expected map[string]int
	}{
		{
			"Test elemenet map",
			args{elemap, doc},
			map[string]int{"html": 1, "head": 1, "body": 1, "a": 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			visit(tt.args.elemap, tt.args.n)
			if !isEqualMap(tt.args.elemap, tt.expected) {
				t.Errorf("%v, want %v", tt.args.elemap, tt.expected)
			}
		})
	}
}

func isEqualMap(a, b map[string]int) bool {
	if len(a) != len(b) {
		return false
	}
	for k, av := range a {
		if bv, ok := b[k]; !ok || bv != av {
			return false
		}
	}
	return true
}
