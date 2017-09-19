package main

import (
	"fmt"

	"github.com/negibokken/golang-training/ch12/ex01/display"
)

func main() {
	fmt.Println("hello")
	display.Display("aaa", struct {
		args string
		str  []string
	}{args: "bbb", str: []string{"ccc", "ddd", "eee"}})
}
