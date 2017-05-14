package main

import "fmt"

const size = 5

func main() {
	var str = []byte("Hello,  world")
	fmt.Println("before: ", string(str))
	deleteRedundantSpace(str)
	fmt.Println("after: ", string(str))
}

func deleteRedundantSpace(str []byte) {
	for i, j := 0, len(str)-1; i < j; i, j = i+1, j-1 {
		str[i], str[j] = str[j], str[i]
	}
}
