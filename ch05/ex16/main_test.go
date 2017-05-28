package main

import (
	"testing"
)

func Test_join(t *testing.T) {
	type args struct {
		sep  string
		strs []string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			"test join",
			args{"/", []string{"a", "b", "c"}},
			"a/b/c",
			false,
		},
		{
			"test join",
			args{"z", []string{"a", "b", "c"}},
			"azbzc",
			false,
		},
		{
			"test join",
			args{"", []string{"a", "b", "c"}},
			"abc",
			false,
		},
		{
			"test join",
			args{"", []string{}},
			"",
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := join(tt.args.sep, tt.args.strs...)
			if (err != nil) != tt.wantErr {
				t.Errorf("join() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("join() = %v, want %v", got, tt.want)
			}
		})
	}
}
