package main

import (
	"os"

	"github.com/negibokken/golang-training/ch12/ex02/display"
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

	type Cycle struct {
		Value int
		Tail  *Cycle
	}
	var c Cycle
	c = Cycle{42, &c}
	display.Display("cyclic", c)
}
