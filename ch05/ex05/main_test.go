package main

import (
	"os"
	"strings"
	"testing"

	"fmt"

	"golang.org/x/net/html"
)

func Test_countWordsAndImages(t *testing.T) {
	doc, err := html.Parse(strings.NewReader(`
		<html>
			<head></head>
			<body>
				<img src="image.png" />
				<p>hello my</p>
				<p>counts words and images.<p>
				<img src="image.png" />
			</body>
		</html>
	`))
	if err != nil {
		fmt.Printf("%v", err)
		os.Exit(1)
	}
	type args struct {
		n *html.Node
	}
	tests := []struct {
		name       string
		args       args
		wantWords  int
		wantImages int
	}{
		{
			"Test Words Count",
			args{doc},
			6,
			2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotWords, gotImages := countWordsAndImages(tt.args.n)
			if gotWords != tt.wantWords {
				t.Errorf("countWordsAndImages() gotWords = %v, want %v", gotWords, tt.wantWords)
			}
			if gotImages != tt.wantImages {
				t.Errorf("countWordsAndImages() gotImages = %v, want %v", gotImages, tt.wantImages)
			}
		})
	}
}
