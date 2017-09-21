package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/negibokken/golang-training/ch12/ex09/decode"
)

func main() {
	dec := decode.NewDecoder(strings.NewReader(`(3 "a" (b))`))
	var tokens []decode.Token
	for {
		token, err := dec.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		tokens = append(tokens, token)
	}
	fmt.Println(tokens)
}
