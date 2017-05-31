package main

import (
	"bytes"
	"fmt"
)

type tree struct {
	value       int
	left, right *tree
}

func (t *tree) String() string {
	var buf bytes.Buffer
	if t.left != nil {
		buf.WriteString(fmt.Sprintf("%d", t.left.value))
	}
	buf.WriteString(fmt.Sprintf("%d", t.value))
	if t.right != nil {
		buf.WriteString(fmt.Sprintf("%d", t.right.value))
	}
	return buf.String()
}

func main() {
	t2 := tree{2, nil, nil}
	t3 := tree{3, nil, nil}
	t := tree{1, &t2, &t3}
	fmt.Println(t.String())
}
