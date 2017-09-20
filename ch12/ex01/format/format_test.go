package format

import (
	"reflect"
	"strings"
	"testing"
)

func TestFormatAtom(t *testing.T) {
	type args struct {
		v interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"invalid", args{}, "invalid"},
		{"int", args{1}, "1"},
		{"uint", args{uint(1)}, "1"},
		{"bool", args{bool(true)}, "true"},
		{"string", args{string("string")}, "\"string\""},
		{"reference", args{map[string]string{"key": "value"}}, "map[string]string"},
		{"reference", args{make(chan string)}, "chan string"},
		{"reference", args{struct{ int }{1}}, "struct { int } value"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.name == "reference" {
				got := FormatAtom(reflect.ValueOf(tt.args.v))
				if !strings.HasPrefix(got, tt.want) {
					t.Errorf("FormatAtom() = %v, want %v", got, tt.want)
				}
			} else {
				if got := FormatAtom(reflect.ValueOf(tt.args.v)); got != tt.want {
					t.Errorf("FormatAtom() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func TestFormatMap(t *testing.T) {
	type args struct {
		v interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"reference", args{map[struct{ x int }]string{{1}: "value"}}, "map[struct { x int }]string"},
		{"reference", args{map[[3]string]string{[3]string{"a", "b", "c"}: "value"}}, "map[[3]string]string"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.name == "reference" {
				got := FormatAtom(reflect.ValueOf(tt.args.v))
				if !strings.HasPrefix(got, tt.want) {
					t.Errorf("FormatAtom() = %v, want %v", got, tt.want)
				}
			} else {
				if got := FormatAtom(reflect.ValueOf(tt.args.v)); got != tt.want {
					t.Errorf("FormatAtom() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}
