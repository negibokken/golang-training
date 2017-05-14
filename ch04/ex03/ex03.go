package main

import "fmt"

const size = 5

func main() {
	var num = [size]int{1, 2, 3, 4, 5}
	fmt.Println("before: ", num)
	reverse(&num)
	fmt.Println("after: ", num)
}

func reverse(num *[size]int) {
	for i, j := 0, len(num)-1; i < j; i, j = i+1, j-1 {
		num[i], num[j] = num[j], num[i]
	}
}
