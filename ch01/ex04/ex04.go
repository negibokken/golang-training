package main

import (
	"bufio"
	"fmt"
	"os"
)

type Info struct {
	filenames []string
	count     int
}

func main() {
	counts := make(map[string]Info)
	files := os.Args[1:]
	countLines(files, counts)
	for line, inf := range counts {
		fmt.Printf("%d\t%s\t%v\n", inf.count, line, inf.filenames)
	}
}

func deleteDuplicated(files []string) []string {
	var arr []string
	for _, file := range files {
		duplicated := false
		for _, a := range arr {
			if file == a {
				duplicated = true
				break
			}
		}
		if duplicated == false {
			arr = append(arr, file)
		}
	}
	return arr
}

func setToMap(input *bufio.Scanner, _counts map[string]Info, arg string) {
	for input.Scan() {
		c := _counts[input.Text()]
		c.count++
		c.filenames = append(c.filenames, arg)
		_counts[input.Text()] = c
	}
}

func countLines(args []string, _counts map[string]Info) {
	var counts = make(map[string]Info)
	if len(args) == 0 {
		input := bufio.NewScanner(os.Stdin)
		setToMap(input, counts, "")
	} else {
		for _, arg := range args {
			f, err := os.Open(arg)
			input := bufio.NewScanner(f)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			setToMap(input, counts, arg)
			f.Close()
		}
	}
	for key := range counts {
		c := counts[key]
		if c.count > 1 {
			c.filenames = deleteDuplicated(c.filenames)
			_counts[key] = c
		}
	}
}
