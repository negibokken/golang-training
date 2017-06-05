package main

import (
	"io/ioutil"
	"os"

	"fmt"

	"./eval"
)

func main() {
	buf, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(1)
	}
	expr, err := eval.Parse(string(buf))
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(1)
	}
	fmt.Println(expr.Eval(eval.Env{"x": 1, "y": 2}))
}
