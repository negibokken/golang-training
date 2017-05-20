package main

import (
	"bytes"
	"testing"

	"strings"

	"io/ioutil"

	"golang.org/x/net/html"
)

func TestVisit(t *testing.T) {
	doc, err := html.Parse(strings.NewReader(`
				<html>
					<a href="https://example.com"></a>
					<a href="https://example.co.jp"></a>
					<a href="https://example.org"></a>
				</html>
			`))
	if err != nil {
		t.Errorf("%v", err)
	}

	bs, err := ioutil.ReadFile("./golang.org.html")
	if err != nil {
		t.Errorf("%v", err)
	}
	doc2, err := html.Parse(strings.NewReader(string(bs)))
	if err != nil {
		t.Errorf("%v", err)
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
			"Test get links",
			args{doc},
			`https://example.com
https://example.co.jp
https://example.org
`,
		},
		{
			"Test get links",
			args{doc2},
			`/
/
#
/doc/
/pkg/
/project/
/help/
/blog/
http://play.golang.org/
#
#
//tour.golang.org/
https://golang.org/dl/
//blog.golang.org/
https://developers.google.com/site-policies#restrictions
/LICENSE
/doc/tos.html
http://www.google.com/intl/en/policies/privacy/
`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out = new(bytes.Buffer)
			visit(tt.args.n)
			got := out.(*bytes.Buffer).String()
			if got != tt.want {
				t.Errorf("--%s--, want --%s--", got, tt.want)
			}
		})
	}
}
