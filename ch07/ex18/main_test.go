package main

import (
	"encoding/xml"
	"reflect"
	"strings"
	"testing"
)

func Test_makeTree(t *testing.T) {
	dec := xml.NewDecoder(strings.NewReader(`
		<html>
			<head></head>
			<body><p>hello</p></body>
		</html>
	`))
	tests := []struct {
		name    string
		args    *xml.Decoder
		want    Node
		wantErr bool
	}{
		{
			"test makeTree",
			dec,
			[]Node{
				Element{
					xml.Name{"", "html"},
					[]xml.Attr{},
					[]Node{},
				},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := makeTree(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("makeTree() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !equal(got, tt.want) {
				t.Errorf("makeTree() = %v, want %v", got, tt.want)
			}
		})
	}
}

func equal(got, want Node) bool {
	gotE := got.(*Element)
	wantE := got.(*Element)
	if !reflect.DeepEqual(gotE.Attr, wantE.Attr) {
		return false
	}
	if !reflect.DeepEqual(gotE.Type, wantE.Type) {
		return false
	}
	return true
}
