package main

import (
	"fmt"
	"os"

	"github.com/negibokken/golang-training/ch12/ex03/sexpr"
)

type Profile struct {
	name string
	age  int
}

func main() {
	// var structTest = []Profile{
	// 	{name: "Alice", age: 20},
	// 	{name: "Bob", age: 21},
	// 	{name: "Charlie", age: 21},
	// }
	var structTest = Profile{name: "Alice", age: 20}
	var out Profile
	data, err := sexpr.Marshal(structTest)
	if err != nil {
		fmt.Printf("%v", err)
		os.Exit(1)
	}
	fmt.Printf("Encode result: %v\n", string(data))
	sexpr.Unmarshal(data, &out)
	fmt.Printf("Decode result: %v\n", out)
}
