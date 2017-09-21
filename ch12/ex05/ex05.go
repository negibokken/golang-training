package main

import (
	"fmt"
	"os"

	"github.com/negibokken/golang-training/ch12/ex05/json"
)

func main() {
	byt, err := json.Marshal(`{"name": "Alice"}`)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(string(byt))
}
