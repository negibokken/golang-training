// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 97.
//!+

package charcount

import (
	"bytes"
	"reflect"
	"strings"
	"testing"
)

func isEqualMap(actual, expect map[rune]int) bool {
	if len(actual) != len(expect) {
		return false
	}
	for k, v := range actual {
		vv, ok := expect[k]
		if !ok {
			return false
		}
		if v != vv {
			return false
		}
	}
	return true
}

func Test_charcount(t *testing.T) {
	tests := []struct {
		name       string
		in         string
		wantCount  map[rune]int
		wantutflen [5]int
	}{
		{
			"test char count",
			"test",
			map[rune]int{
				't': 2,
				'e': 1,
				's': 1,
			},
			[5]int{0, 4, 0, 0, 0},
		},
		{
			"test empty char count",
			"",
			map[rune]int{},
			[5]int{0, 0, 0, 0, 0},
		},
		{
			"test char count",
			"test あいう",
			map[rune]int{
				't': 2,
				'e': 1,
				's': 1,
				' ': 1,
				'あ': 1,
				'い': 1,
				'う': 1,
			},
			[5]int{0, 5, 0, 3, 0},
		},
	}
	for _, tt := range tests {
		inr = new(bytes.Buffer)
		inr = strings.NewReader(tt.in)
		t.Run(tt.name, func(t *testing.T) {
			count, utflen := Charcount()
			if !isEqualMap(count, tt.wantCount) {
				t.Errorf("charcount() = %v, want %v", count, tt.wantCount)
			}
			if !reflect.DeepEqual(utflen, tt.wantutflen) {
				t.Errorf("charcount() = %v, want %v", utflen, tt.wantutflen)
			}
		})
	}
}
