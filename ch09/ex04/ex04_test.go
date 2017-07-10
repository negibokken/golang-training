package main

import (
	"fmt"
	"runtime"
	"testing"
)

func BenchmarkPipeLine10(b *testing.B)      { benchmarkMyPipeLine(b, 10) }
func BenchmarkPipeLine100(b *testing.B)     { benchmarkMyPipeLine(b, 100) }
func BenchmarkPipeLine1000(b *testing.B)    { benchmarkMyPipeLine(b, 1000) }
func BenchmarkPipeLine10000(b *testing.B)   { benchmarkMyPipeLine(b, 10000) }
func BenchmarkPipeLine100000(b *testing.B)  { benchmarkMyPipeLine(b, 100000) }
func BenchmarkPipeLine1000000(b *testing.B) { benchmarkMyPipeLine(b, 1000000) }

func benchmarkMyPipeLine(b *testing.B, n int) {
	fmt.Printf("\n---- Number of Goroutine: %d ----", n)
	var startMemory runtime.MemStats
	runtime.ReadMemStats(&startMemory)
	in, out := MyPipeLine(n)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		go func() { in <- struct{}{} }()
		<-out
	}

	var endMemory runtime.MemStats
	runtime.ReadMemStats(&endMemory)

	fmt.Printf("\nmemory all: %f MB\n",
		float64(endMemory.Alloc-startMemory.Alloc)/float64(1024*1024))
}
