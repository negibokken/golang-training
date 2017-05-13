package main

import "testing"

func TestIsAnagram(t *testing.T) {
	type args struct {
		str1 string
		str2 string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"アナグラムのとき", args{"dormitory", "dirtyroom"}, true},
		{"アナグラムのとき", args{"abc", "acb"}, true},
		{"アナグラムじゃないとき", args{"abd", "acb"}, false},
		{"アナグラムじゃないとき", args{"a", "ab"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isAnagram(tt.args.str1, tt.args.str2); got != tt.want {
				t.Errorf("isAnagram() = %v, want %v", got, tt.want)
			}
		})
	}
}
