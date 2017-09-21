package main

import (
	"fmt"
	"os"
	"reflect"

	"github.com/negibokken/golang-training/ch12/ex13/sexpr"
)

func main() {
	type Interface interface{}
	type Record struct {
		B    bool
		F32  float32
		F64  float64
		C64  complex64
		C128 complex128
		I    Interface `sexpr:"face"`
	}
	sexpr.Interfaces["sexpr.Interface"] = reflect.TypeOf(int(0))
	input := `((B t) (F32 2.5) (F64 0) (I ("sexpr.Interface" 5)))`

	var r Record
	err := sexpr.Unmarshal([]byte(input), &r)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println(r)
}
