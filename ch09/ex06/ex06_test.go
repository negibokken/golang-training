package main

import (
	"bytes"
	"fmt"
	"testing"
	"time"
)

func Test_main(t *testing.T) {
	out = new(bytes.Buffer)
	start := time.Now()
	main()
	fmt.Printf("time: %v\n", time.Since(start))
}
