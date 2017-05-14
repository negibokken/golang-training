package main

import "os"
import "fmt"
import "bufio"

func main() {
	counts := make(map[string]int)

	if os.Args[1] == "" {
		fmt.Println("./ex09 <filename>")
		os.Exit(0)
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Printf("%v", err)
		os.Exit(1)
	}
	in := bufio.NewScanner(file)
	in.Split(bufio.ScanWords)
	for in.Scan() {
		counts[in.Text()]++
	}
	for k, v := range counts {
		fmt.Printf("%v: %v\n", k, v)
	}
}
