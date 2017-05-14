package main

import "fmt"

const size = 5

func main() {
	var num = [size]int{1, 2, 3, 4, 5}
	d := 3
	fmt.Println("before: ", num)
	rotate(&num, d)
	fmt.Println("after: ", num)
}

func rotate(num *[size]int, d int) {
	for dd := 0; dd < d; dd++ {
		for i, j := 0, 1; i < len(num)-1; i, j = i+1, j+1 {
			num[i], num[j] = num[j], num[i]
		}
	}
}
