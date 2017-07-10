package main

import (
	"flag"
	"fmt"
	"time"
)

func MyPipeLine(size int) (in chan<- struct{}, out <-chan struct{}) {
	for i := 0; i < size; i++ {
		c := make(chan struct{})
		go func(in chan<- struct{}, out <-chan struct{}) {
			for {
				in <- <-out
			}
		}(c, out)
		if in == nil {
			in = c
		}
		out = c
	}
	return
}

func main() {
	size := 10
	flag.IntVar(&size, "size", 10000, "pipeline size")
	flag.Parse()

	fmt.Printf("pipeline size: %d\n", size)
	in, out := MyPipeLine(size)
	start := time.Now()
	go func() { in <- struct{}{} }()
	<-out
	fmt.Printf("%v", time.Since(start))
}
