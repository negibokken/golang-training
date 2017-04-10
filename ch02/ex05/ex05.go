package main

import (
	"fmt"

	"./popcount"
)

func main() {
	fmt.Println("vim-go")
	for i := 0; i < 8; i++ {
		fmt.Printf("%v has %v bit of 1 in bits\n", i, popcount.BitClearPopCount(uint64(i)))
	}
}
