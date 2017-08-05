package main

import (
	"fmt"

	"github.com/negibokken/golang-training/ch11/ex01/charcount"
)

func main() {
	counts, utflen := charcount.Charcount()
	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Print("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
}
