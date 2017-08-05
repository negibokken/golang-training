package main

import (
	"bytes"
	"reflect"
	"testing"
)

func Test_execGoList(t *testing.T) {
	type args struct {
		arg string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{"Test go list", args{"github.com/negibokken/golang-training/ch10/ex02"}, []string{
			"archive/tar", "archive/zip", "bufio", "bytes", "compress/flate", "encoding/binary",
			"errors", "fmt", "github.com/negibokken/golang-training/ch10/ex02/archive",
			"github.com/negibokken/golang-training/ch10/ex02/archive/tar",
			"github.com/negibokken/golang-training/ch10/ex02/archive/zip",
			"hash", "hash/crc32", "internal/race", "io", "io/ioutil", "log", "math", "os",
			"path", "path/filepath", "reflect", "runtime", "runtime/internal/atomic",
			"runtime/internal/sys", "sort", "strconv", "strings", "sync", "sync/atomic",
			"syscall", "time", "unicode", "unicode/utf8", "unsafe",
		}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := execGoList(tt.args.arg)
			if (err != nil) != tt.wantErr {
				t.Errorf("execGoList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("execGoList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_printPackages(t *testing.T) {
	type args struct {
		ma map[string]bool
	}
	tests := []struct {
		name     string
		args     args
		expected string
	}{
		{
			"Test Print Packages",
			args{map[string]bool{
				"test":  true,
				"test2": true,
				"test3": true,
			}},
			"\t- test\n\t- test2\n\t- test3\n",
		},
		{
			"Test Print Packages2",
			args{map[string]bool{
				"test":   true,
				"test2":  true,
				"test3":  true,
				"test4":  true,
				"test5":  true,
				"test6":  true,
				"test7":  true,
				"test8":  true,
				"test9":  true,
				"test10": true,
			}},
			"\t- test\n\t- test10\n\t- test2\n\t- test3\n\t- test4\n\t- test5\n\t- test6\n\t- test7\n\t- test8\n\t- test9\n",
		},
		{
			"If packages is empty",
			args{map[string]bool{}},
			"",
		},
	}
	// TODO: Add test cases.
	for _, tt := range tests {
		out = new(bytes.Buffer)
		t.Run(tt.name, func(t *testing.T) {
			printPackages(tt.args.ma)
			if got := out.(*bytes.Buffer).String(); got != tt.expected {
				t.Errorf("printPackages() = %q, want %q", got, tt.expected)
			}
		})
	}
}
