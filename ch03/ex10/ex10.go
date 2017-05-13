package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	str := os.Args[1]
	fmt.Println(comma(str))
}

func comma(str string) string {
	var v bytes.Buffer
	i := 3 - len(str)
	for idx, s := range str {
		if i%3 == 0 && idx != 0 {
			v.WriteString(",")
		}
		v.WriteRune(s)
		i++
	}
	return v.String()
}
