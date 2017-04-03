package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

var out io.Writer = os.Stdout

func InefficientEcho(args []string) error {
	s, sep := "", ""
	for _, arg := range args {
		s += sep + arg
		sep = " "
	}
	fmt.Fprintln(out, s)
	return nil
}

func EfficientEcho(args []string) error {
	fmt.Fprintln(out, strings.Join(args, " "))
	return nil
}

func main() {
	InefficientEcho(os.Args)
	EfficientEcho(os.Args)
}
