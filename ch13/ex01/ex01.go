package main

import (
	"fmt"

	"github.com/negibokken/golang-training/ch13/ex01/equalish"
)

func main() {
	samples := []struct {
		x interface{}
		y interface{}
	}{
		{1, 1},
		{1, 2},
		{"hello", "hello"},
	}
	for _, sample := range samples {
		fmt.Printf("Equalish %v == %v ? = %v\n", sample.x, sample.y, equalish.Equalish(sample.x, sample.y))
	}
}
