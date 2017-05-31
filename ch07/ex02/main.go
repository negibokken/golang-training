package main

import (
	"bytes"
	"fmt"
	"io"
)

// Counter wrap and num
type Counter struct {
	writer io.Writer
	count  int64
}

// Write redeclare write
func (c *Counter) Write(p []byte) (int, error) {
	n, err := c.writer.Write(p)
	c.count += int64(n)
	return int(n), err
}

// CountingWriter  wrap io.Writer
func CountingWriter(w io.Writer) (io.Writer, *int64) {
	c := &Counter{writer: w}
	return c, &c.count
}

func main() {
	var b bytes.Buffer
	c, num := CountingWriter(&b)
	c.Write([]byte("hello world"))
	fmt.Println(c, num)
}
