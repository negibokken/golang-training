package main

import "fmt"

const size = 5

func main() {
	var str = []string{"a", "a", "b", "b", "c", "c", "d", "d"}
	fmt.Println("before: ", str)
	newStr := deleteNeighborRedundant(str)
	fmt.Println("after: ", newStr)
}

func deleteNeighborRedundant(str []string) []string {
	var result = str[:1]
	for i := 1; i < len(str); i++ {
		if str[i-1] == str[i] {
			continue
		}
		result = append(result, str[i])
	}
	return result
}
