package main

import (
	"testing"
)

func TestDeleteNeighborRedundant(t *testing.T) {
	type args struct {
		str []string
	}
	tests := []struct {
		name     string
		args     args
		expected []string
	}{
		{"delete redundant neighbor", args{[]string{"a", "a", "b"}}, []string{"a", "b"}},
		{"delete redundant neighbor", args{[]string{"a", "a", "b", "b", "c", "c", "a"}}, []string{"a", "b", "c", "a"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := deleteNeighborRedundant(tt.args.str)
			if !isEqual(got, tt.expected) {
				t.Errorf("got: %v expected: %v", got, tt.expected)
			}
		})
	}
}

func isEqual(str1, str2 []string) bool {
	if len(str1) != len(str2) {
		return false
	}
	for i := 0; i < len(str1); i++ {
		if str1[i] != str2[i] {
			return false
		}
	}
	return true
}
