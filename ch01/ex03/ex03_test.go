package main

import "testing"

func BenchmarkInefficientEcho(b *testing.B) {
	for i := 0; i < b.N; i++ {
		InefficientEcho([]string{"./ex03", "aa", "bb", "cc"})
	}
}

func BenchmarkEfficientEcho(b *testing.B) {
	for i := 0; i < b.N; i++ {
		EfficientEcho([]string{"./ex03", "aa", "bb", "cc"})
	}
}
