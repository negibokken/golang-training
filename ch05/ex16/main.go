package main

import "fmt"
import "os"

func main() {
	str, err := join("/", "a", "abc", "abcdefg")
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(1)
	}
	fmt.Printf("%v", str)
}

func join(sep string, strs ...string) (string, error) {
	if len(strs) == 0 {
		return "", fmt.Errorf("strs is nil")
	}
	result := strs[0]

	if len(strs) == 1 {
		return result, nil
	}

	for _, s := range strs[1:] {
		result += sep + s
	}
	return result, nil
}
