package main

import (
	"fmt"

	"github.com/negibokken/golang-training/ch13/ex02/cyclic"
)

func main() {
	type Cyclic struct {
		v int
		c *Cyclic
	}
	var c Cyclic
	c = Cyclic{1, &c}

	samples := []struct {
		x interface{}
	}{
		{c},
		{1},
		{"hello"},
	}
	for _, sample := range samples {
		fmt.Printf("Cyclic(%v) = %v\n", sample.x, cyclic.Cyclic(sample.x))
	}
}
