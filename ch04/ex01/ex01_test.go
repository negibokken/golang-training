package main

import (
	"crypto/sha256"
	"testing"
)

func TestCompareSHA256Bits(t *testing.T) {
	type args struct {
		c1 [32]byte
		c2 [32]byte
	}
	tests := []struct {
		name  string
		args  args
		want  bool
		want1 int
	}{
		{
			"一致しているとき",
			args{sha256.Sum256([]byte("x")),
				sha256.Sum256([]byte("x"))},
			true,
			0,
		},
		{
			"一致していないとき",
			args{sha256.Sum256([]byte("x")),
				sha256.Sum256([]byte("X"))},
			false,
			31,
		},
		{
			"一致していないとき",
			args{sha256.Sum256([]byte("abcdefg")),
				sha256.Sum256([]byte("vwxyz"))},
			false,
			32,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := compareSHA256Bits(tt.args.c1, tt.args.c2)
			if got != tt.want {
				t.Errorf("compareSHA256Bits() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("compareSHA256Bits() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
