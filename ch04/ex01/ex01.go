package main

import (
	"crypto/sha256"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println(len(os.Args))
		fmt.Println("./ex01 <str1> <str2>")
		os.Exit(0)
	}
	c1 := sha256.Sum256([]byte(os.Args[1]))
	c2 := sha256.Sum256([]byte(os.Args[2]))
	f, num := compareSHA256Bits(c1, c2)
	fmt.Printf("arg1: %v\narg2: %v\n", os.Args[1], os.Args[2])
	fmt.Printf("equality: %v num: %v\n", f, num)
}

func compareSHA256Bits(c1, c2 [32]byte) (bool, int) {
	cnt := 0
	for idx, c := range c1 {
		if c != c2[idx] {
			cnt++
		}
	}
	return cnt == 0, cnt
}
