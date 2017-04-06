package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCountLines(t *testing.T) {
	var tests = []struct {
		args   []string
		counts map[string]Info
	}{
		{
			[]string{"aaa.txt"},
			map[string]Info{
				"bar": Info{
					filenames: []string{"aaa.txt"},
					count:     2,
				},
			},
		},
		{
			[]string{"bbb.txt"},
			map[string]Info{
				"hoge": Info{
					filenames: []string{"bbb.txt"},
					count:     2,
				},
			},
		},
		{
			[]string{"aaa.txt", "bbb.txt"},
			map[string]Info{
				"bar": Info{
					filenames: []string{"aaa.txt"},
					count:     2,
				},
				"hoge": Info{
					filenames: []string{"bbb.txt"},
					count:     2,
				},
				"foo": Info{
					filenames: []string{"aaa.txt", "bbb.txt"},
					count:     2,
				},
			},
		},
	}

	for _, test := range tests {
		descr := fmt.Sprintf("countLines(%q)", test.args)
		counts := make(map[string]Info)
		countLines(test.args, counts)
		if !reflect.DeepEqual(counts, test.counts) {
			t.Errorf("%s failed", descr)
		}
	}
}

func TestDeleteDuplicated(t *testing.T) {
	var tests = []struct {
		files    []string
		expected []string
	}{
		{
			files:    []string{"aaa.txt", "bbb.txt", "ccc.txt", "aaa.txt"},
			expected: []string{"aaa.txt", "bbb.txt", "ccc.txt"},
		},
	}

	for _, test := range tests {
		actual := deleteDuplicated(test.files)
		if !reflect.DeepEqual(actual, test.expected) {
			t.Errorf("%s failed: Expected:%v", test.files, test.expected)
		}
	}
}
