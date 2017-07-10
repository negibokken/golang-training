package main

import (
	"testing"
)

func BenchmarkPipeLine10(b *testing.B)    { benchmarkMyPipeLine(b, 10) }
func BenchmarkPipeLine100(b *testing.B)   { benchmarkMyPipeLine(b, 100) }
func BenchmarkPipeLine1000(b *testing.B)  { benchmarkMyPipeLine(b, 1000) }
func BenchmarkPipeLine10000(b *testing.B) { benchmarkMyPipeLine(b, 10000) }

func benchmarkMyPipeLine(b *testing.B, n int) {
	in, out := MyPipeLine(n)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		go func() { in <- struct{}{} }()
		<-out
	}
}
