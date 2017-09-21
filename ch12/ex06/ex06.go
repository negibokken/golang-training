package main

import (
	"fmt"
	"os"

	"github.com/negibokken/golang-training/ch12/ex06/sexpr"
)

type Profile struct {
	Name string
	Age  int
}

func main() {
	var structTest = Profile{Name: "Alice", Age: 0}
	by, err := sexpr.Marshal(structTest)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("%v", by)
	var prof Profile
	err = sexpr.Unmarshal(by, &prof)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("%v", prof)
}
