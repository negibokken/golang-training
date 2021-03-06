package main

import (
	"fmt"
	"os"

	"github.com/negibokken/golang-training/ch12/ex10/sexpr"
)

func main() {
	type Interface interface{}
	type Record struct {
		B    bool
		F32  float32
		F64  float64
		C64  complex64
		C128 complex128
		I    Interface
	}
	var r Record
	err := sexpr.Unmarshal([]byte(`((B t) (F32 2.5) (F64 0) (I nil))`), &r)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println(r)
}
