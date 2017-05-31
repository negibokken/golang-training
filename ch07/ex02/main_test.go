package main

import (
	"bytes"
	"io"
	"testing"
)

func TestCountingWriter(t *testing.T) {
	tests := []struct {
		name  string
		buf   []byte
		want1 int64
		want2 int
		wantW string
	}{
		{"test counting writer", []byte("hello world"), int64(11), 11, "hello world"},
		{"test counting writer", []byte("good golang!"), int64(12), 12, "good golang!"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var w bytes.Buffer
			got, got1 := CountingWriter(&w)
			num, err := got.Write([]byte("hello world"))
			if err != nil {
				t.Errorf("%v", err)
			}
			if num != tt.want2 {
				t.Errorf("got: %v, want: %v", num, tt.want2)
			}
			if *got1 != tt.want1 {
				t.Errorf("CountingWriter() got1 = %v, want %v", got1, tt.want1)
			}
			if gotW := w.String(); gotW != tt.wantW {
				t.Errorf("CountingWriter() gotW = %v, want %v", gotW, tt.wantW)
			}
		})
	}
}
