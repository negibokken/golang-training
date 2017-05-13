package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	str := os.Args[1]
	fmt.Println(comma(str))
}

func comma(str string) string {
	var v bytes.Buffer
	splited := strings.Split(str, ".")
	bef := splited[0]
	var aft string
	if len(splited) == 2 {
		aft = "." + splited[1]
	}
	i := 3 - len(bef)
	for idx, s := range bef {
		if i%3 == 0 && idx != 0 {
			v.WriteString(",")
		}
		v.WriteRune(s)
		i++
	}
	v.WriteString(aft)
	return v.String()
}
