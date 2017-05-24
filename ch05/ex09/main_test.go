package main

import "testing"

func TestExpand(t *testing.T) {
	type args struct {
		s string
		f func(string) string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"test expand",
			args{
				"string$123",
				func(s string) string { return s + s },
			},
			"string123123",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := expand(tt.args.s, tt.args.f); got != tt.want {
				t.Errorf("expand() = %v, want %v", got, tt.want)
			}
		})
	}
}
