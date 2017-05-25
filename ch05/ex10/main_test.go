package main

import (
	"testing"
)

func TestTopoSort(t *testing.T) {
	type args struct {
		m map[string]map[string]bool
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			"test topological sort",
			args{
				map[string]map[string]bool{
					"algorithm":      map[string]bool{"data structure": true},
					"calculus":       map[string]bool{"data structure": true},
					"data structure": map[string]bool{"intro to programming": true},
				},
			},
			[][]string{
				{"intro to programming", "data structure", "algorithm", "calculus"},
				{"intro to programming", "data structure", "calculus", "algorithm"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := topoSort(tt.args.m)
			if isEqualSomeWant(got, tt.want) {
				t.Errorf("topoSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func isEqualSomeWant(got []string, want [][]string) bool {
	var flag bool
	for _, ww := range want {
		if len(got) != len(ww) {
			return false
		}
		flag = true
		// Check each element in slice
		for i, w := range ww {
			if got[i] != w {
				flag = false
				break
			}
		}
	}
	return flag
}
