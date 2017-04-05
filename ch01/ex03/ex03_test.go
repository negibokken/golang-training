package main

import (
	"bytes"
	"fmt"
	"testing"
)

func BenchmarkInefficientEcho(b *testing.B) {
	for i := 0; i < b.N; i++ {
		InefficientEcho([]string{"./ex03", "aa", "bb", "cc"})
	}
}

func BenchmarkEfficientEcho(b *testing.B) {
	for i := 0; i < b.N; i++ {
		EfficientEcho([]string{"./ex03", "aa", "bb", "cc"})
	}
}

func TestInefficientEcho(t *testing.T) {
	programName := "ex03"
	var tests = []struct {
		args []string
		want string
	}{
		{[]string{programName, "foo", "bar"}, fmt.Sprintf("%v foo bar\n", programName)},
		{[]string{programName, "hoge"}, fmt.Sprintf("%v hoge\n", programName)},
		{[]string{programName}, fmt.Sprintf("%v\n", programName)},
	}

	for _, test := range tests {
		descr := fmt.Sprintf("InefficientEcho(%q)", test.args)
		out = new(bytes.Buffer)
		if err := InefficientEcho(test.args); err != nil {
			t.Errorf("%s failed: %v", descr, err)
		}
		got := out.(*bytes.Buffer).String()
		if got != test.want {
			t.Errorf("%s = %q, want %q", descr, got, test.want)
		}
	}
}

func TestEfficientEcho(t *testing.T) {
	programName := "ex03"
	var tests = []struct {
		args []string
		want string
	}{
		{[]string{programName, "foo", "bar"}, fmt.Sprintf("%v foo bar\n", programName)},
		{[]string{programName, "hoge"}, fmt.Sprintf("%v hoge\n", programName)},
		{[]string{programName}, fmt.Sprintf("%v\n", programName)},
	}

	for _, test := range tests {
		descr := fmt.Sprintf("EfficientEcho(%q)", test.args)
		out = new(bytes.Buffer)
		if err := InefficientEcho(test.args); err != nil {
			t.Errorf("%s failed: %v", descr, err)
		}
		got := out.(*bytes.Buffer).String()
		if got != test.want {
			t.Errorf("%s = %q, want %q", descr, got, test.want)
		}
	}
}
