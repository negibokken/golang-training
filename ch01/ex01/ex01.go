package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

var out io.Writer = os.Stdout

func echo(args []string) error {
	fmt.Fprintln(out, strings.Join(args, " "))
	return nil
}

func main() {
	echo(os.Args)
}
