package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func TestLimitReader(t *testing.T) {
	type args struct {
		r string
		n int64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"test limit reader", args{"abcdefg", 5}, "abcde"},
		{"test limit reader", args{"abcdefg", 7}, "abcdefg"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := bytes.NewReader([]byte(tt.args.r))
			rr := LimitReader(r, tt.args.n)
			buf, err := ioutil.ReadAll(rr)
			if err != nil {
				fmt.Fprintf(os.Stderr, "%v", err)
			}
			if string(buf) != tt.want {
				t.Errorf("LimitReader() = %v, want %v", string(buf), tt.want)
			}
		})
	}
}
