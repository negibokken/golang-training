package main

import (
	"os"

	"fmt"

	"./eval"
)

func main() {
	expr, err := eval.Parse(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
	}
	fmt.Println(expr.Eval(eval.Env{"x": 1, "y": 2}))
}
