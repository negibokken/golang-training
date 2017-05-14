package main

import "fmt"

const size = 5

func main() {
	var str = []byte("  Hello,  world")
	fmt.Println("before: ", string(str))
	newStr := deleteRedundantSpace(str)
	fmt.Println("after: ", string(newStr))
}

func deleteRedundantSpace(str []byte) []byte {
	var result []byte
	for i := 0; i < len(str); i++ {
		if string(str[i]) == " " && string(str[i+1]) == " " {
			result = append(result, str[i])
			i++
			continue
		}
		result = append(result, str[i])
	}
	return result
}
