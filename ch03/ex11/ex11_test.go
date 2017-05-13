package main

import "testing"

func TestComma(t *testing.T) {
	tests := []struct {
		name string
		args string
		want string
	}{
		{"長さが0のとき", "", ""},
		{"長さが1のとき", "1", "1"},
		{"長さが2のとき", "11", "11"},
		{"長さが3のとき", "111", "111"},
		{"長さが3Nのとき", "111111", "111,111"},
		{"長さが3N+1のとき", "1111111", "1,111,111"},
		{"長さが3N+2のとき", "11111111", "11,111,111"},
		{"小数かつ長さが2のとき", "1.1", "1.1"},
		{"小数かつ長さが3のとき", "11.1", "11.1"},
		{"小数かつ長さが3Nのとき", "1111.11", "1,111.11"},
		{"小数かつ長さが3N+1のとき", "1111.111", "1,111.111"},
		{"小数かつ長さが3N+2のとき", "11111.111", "11,111.111"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := comma(tt.args); got != tt.want {
				t.Errorf("comma() = %v, want %v", got, tt.want)
			}
		})
	}
}
