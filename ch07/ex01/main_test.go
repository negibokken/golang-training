package main

import (
	"fmt"
	"testing"
)

func TestWordCounter_Writer(t *testing.T) {
	type args struct {
		p []byte
	}
	tests := []struct {
		name    string
		args    args
		want    int
		count   string
		wantErr bool
	}{
		{"test word counte writer", args{[]byte("hello world")}, 11, "2", false},
		{"test word counte writer", args{[]byte("hello my world")}, 14, "3", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var w WordCounter
			got, err := w.Writer(tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("WordCounter.Writer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("WordCounter.Writer() = %v, want %v", got, tt.want)
			}
			if s := fmt.Sprintf("%v", w); s != tt.count {
				t.Errorf("got %v, want %v", s, tt.count)
			}
		})
	}
}

func TestLineCounter_Writer(t *testing.T) {
	type args struct {
		p []byte
	}
	tests := []struct {
		name    string
		args    args
		want    int
		count   string
		wantErr bool
	}{
		{"test line counte writer", args{[]byte("hello\nworld\n")}, 12, "2", false},
		{"test line counte writer", args{[]byte("\nhello\n\nworld\n")}, 14, "4", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var l LineCounter
			got, err := l.Writer(tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("LineCounter.Writer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("LineCounter.Writer() = %v, want %v", got, tt.want)
			}
		})
	}
}
