package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

func echoUnit(num int, arg string) error {
	fmt.Fprintln(out, fmt.Sprintf("%v %v", num, arg))
	return nil
}

func main() {
	for i, arg := range os.Args {
		echoUnit(i, arg)
	}
}
