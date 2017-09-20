package main

import (
	"os"

	"github.com/negibokken/golang-training/ch12/ex01/display"
)

func main() {
	display.Display("os.Stdin", os.Stdin)
	display.Display("os.Stdout", os.Stdout)
	display.Display("os.Stderr", os.Stderr)

	var i int = 1
	display.Display("i", i)

	var m = map[string]string{
		"key": "value",
	}
	display.Display("map", m)

	var structMap = map[struct {
		x string
		y string
	}]string{
		{x: "a", y: "b"}: "hello",
	}
	display.Display("map", structMap)
}
