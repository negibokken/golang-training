package main

import (
	"bytes"
	"fmt"
	"testing"
)

func TestEchoUnit(t *testing.T) {
	programName := "./ex02"
	var tests = []struct {
		args []string
		want []string
	}{
		{
			[]string{programName, "foo", "bar"},
			[]string{fmt.Sprintf("0 %v\n", programName), "1 foo\n", "2 bar\n"},
		},
		{
			[]string{programName, "foo"},
			[]string{fmt.Sprintf("0 %v\n", programName), "1 foo\n"},
		},
		{
			[]string{programName},
			[]string{fmt.Sprintf("0 %v\n", programName)},
		},
	}

	for _, test := range tests {
		for i, arg := range test.args {
			descr := fmt.Sprintf("echo(%d, %q)", i, arg)
			out = new(bytes.Buffer)
			if err := echoUnit(i, arg); err != nil {
				t.Errorf("%s failed: %v", descr, err)
				continue
			}
			got := out.(*bytes.Buffer).String()
			if got != test.want[i] {
				t.Errorf("%s = %q, want %q", descr, got, test.want[i])
			}
		}
	}
}
