package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) == 2 {
		fmt.Println("./ex12 <str1> <str2>")
		os.Exit(0)
	}
	str1 := os.Args[1]
	str2 := os.Args[2]
	fmt.Printf("isAnagram(%v, %v) = %v\n", str1, str2, isAnagram(str1, str2))
}

func isAnagram(str1, str2 string) bool {
	if len(str1) != len(str2) {
		return false
	}
	cnt := 0
	for _, s := range str1 {
		if strings.Contains(str2, string(s)) {
			cnt++
		}
	}
	return cnt == len(str2)
}
