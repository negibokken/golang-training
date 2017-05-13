package main

import "fmt"

func main() {
	const (
		KB = 1000
		MB = KB * 1000
		GB = MB * 1000
		TB = GB * 1000
		PB = TB * 1000
		EB = PB * 1000
		ZB = EB * 1000
		YB = ZB * 1000
	)
	fmt.Printf("KB = %v B\n", KB)
	fmt.Printf("MB = %v B\n", MB)
	fmt.Printf("GB = %v B\n", GB)
	fmt.Printf("TB = %v B\n", TB)
	fmt.Printf("PB = %v B\n", PB)
	fmt.Printf("EB = %v B\n", EB)
	// fmt.Printf("ZB = %v B\n", ZB) // cannot display
	// fmt.Printf("YB = %v B\n", YB) // cannot display
}
