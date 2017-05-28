package main

import "fmt"

func main() {
	fmt.Printf(foo())
}

func foo() (result string) {
	defer func() { recover() }()
	result = "bar"
	panic(nil)
}
