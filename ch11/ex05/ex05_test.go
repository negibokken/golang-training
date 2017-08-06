package main

import (
	"strings"
	"testing"
)

func isEqualSlice(s, v []string) bool {
	if len(s) != len(v) {
		return false
	}
	for i, ss := range s {
		if ss != v[i] {
			return false
		}
	}
	return true
}

func TestSplit(t *testing.T) {
	var tests = []struct {
		name      string
		s         string
		sep       string
		expect    int
		expectStr []string
	}{
		{
			"test normal split",
			"a:b:c",
			":",
			3,
			[]string{"a", "b", "c"},
		},
		{
			"test normal split",
			"a,b,c,d,e,f,g",
			",",
			7,
			[]string{"a", "b", "c", "d", "e", "f", "g"},
		},
		{
			"test multi byte split",
			"ハあロあーあワあーあルあド",
			"あ",
			7,
			[]string{"ハ", "ロ", "ー", "ワ", "ー", "ル", "ド"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			words := strings.Split(test.s, test.sep)
			if got := len(words); got != test.expect {
				t.Errorf("len(strings.Split(%s,\"%s\")) = %d, want %d", test.s, test.sep, got, test.expect)
			}
			if !isEqualSlice(words, test.expectStr) {
				t.Errorf("strings.Split(%s,\"%s\") = %s, want %s", test.s, test.sep, words, test.expectStr)
			}
		})
	}
}
