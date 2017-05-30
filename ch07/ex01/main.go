package main

import (
	"bufio"
	"bytes"
	"fmt"
)

// WordCounter that counts number of word
type WordCounter int

// Writer that write and count word number
func (w *WordCounter) Writer(p []byte) (int, error) {
	in := bufio.NewScanner((bytes.NewReader(p)))
	in.Split(bufio.ScanWords)
	for in.Scan() {
		*w++
	}
	return len(p), nil
}

// LineCounter that counts number of new line
type LineCounter int

// Writer that write and count new line number
func (l *LineCounter) Writer(p []byte) (int, error) {
	in := bufio.NewScanner((bytes.NewReader(p)))
	in.Split(bufio.ScanLines)
	for in.Scan() {
		*l++
	}
	return len(p), nil
}

func main() {
	var w WordCounter

	w.Writer([]byte("hello my world"))
	fmt.Println(w)

	var l LineCounter
	l.Writer([]byte("hello\nmy\nworld\n\n"))
	fmt.Println(l)
}
