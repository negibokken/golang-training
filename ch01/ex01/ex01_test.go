package main

import (
	"bytes"
	"fmt"
	"testing"
)

func TestEcho(t *testing.T) {
	programName := "./ex01"
	var tests = []struct {
		args []string
		want string
	}{
		{[]string{programName, "foo", "bar"}, fmt.Sprintf("%v foo bar\n", programName)},
		{[]string{programName, "hoge"}, fmt.Sprintf("%v hoge\n", programName)},
		{[]string{programName}, fmt.Sprintf("%v\n", programName)},
	}

	for _, test := range tests {
		descr := fmt.Sprintf("echo(%q)", test.args)

		out = new(bytes.Buffer)
		if err := echo(test.args); err != nil {
			t.Errorf("%s failed: %v", descr, err)
			continue
		}
		got := out.(*bytes.Buffer).String()
		if got != test.want {
			t.Errorf("%s = %q, want %q", descr, got, test.want)
		}
	}
}
