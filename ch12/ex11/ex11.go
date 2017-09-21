package main

import (
	"fmt"
	"os"

	"github.com/negibokken/golang-training/ch12/ex11/params"
)

func main() {
	s := struct {
		Name string `http:"name"`
		Age  int    `http:"age"`
	}{"Alice", 20}
	url, err := params.Pack(&s)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println(url)
}
