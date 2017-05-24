package main

import (
	"testing"
)

func TestShouldSkipped(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"should skipped",
			args{"  "},
			true,
		},
		{
			"should skipped",
			args{"\n "},
			true,
		},
		{
			"should skipped",
			args{"\t "},
			true,
		},
		{
			"should skipped",
			args{" a"},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := shouldSkipped(tt.args.str); got != tt.want {
				t.Errorf("shouldSkipped() = %v, want %v", got, tt.want)
			}
		})
	}
}
