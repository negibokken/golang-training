package main

import "testing"
import "bytes"

func Test_deleteRedundantSpace(t *testing.T) {
	type args struct {
		str []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{"逆順", args{[]byte("Hello, World")}, []byte("dlroW ,olleH")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			deleteRedundantSpace(tt.args.str)
			if !bytes.Equal(tt.args.str, tt.want) {
				t.Errorf("got: %v, expecte: %v", tt.args.str, tt.want)
			}
		})
	}
}
