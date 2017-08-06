package main

import (
	"fmt"

	"github.com/negibokken/golang-training/ch11/ex07/intset"
	"github.com/negibokken/golang-training/ch11/ex07/intset32"
	"github.com/negibokken/golang-training/ch11/ex07/mapintset"
)

func main() {
	b1 := intset.NewBitIntSet()
	b2 := intset32.NewBitIntSet()
	b3 := mapintset.NewMapIntSet()

	b1.Add(1)
	b2.Add(2)
	b3.Add(3)

	fmt.Println(b1.String())
	fmt.Println(b2.String())
	fmt.Println(b3.String())
	b4 := b3.Copy()
	fmt.Println(b4.String())
}
