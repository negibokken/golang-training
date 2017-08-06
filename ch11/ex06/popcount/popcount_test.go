package popcount

import "testing"

func Benchmark10(b *testing.B)      { benchmarkPopCount(b, 10) }
func Benchmark100(b *testing.B)     { benchmarkPopCount(b, 100) }
func Benchmark1000(b *testing.B)    { benchmarkPopCount(b, 1000) }
func Benchmark10000(b *testing.B)   { benchmarkPopCount(b, 10000) }
func Benchmark100000(b *testing.B)  { benchmarkPopCount(b, 100000) }
func Benchmark1000000(b *testing.B) { benchmarkPopCount(b, 1000000) }

func BenchmarkMyPopCount10(b *testing.B)      { benchmarkMyPopCount(b, 10) }
func BenchmarkMyPopCount100(b *testing.B)     { benchmarkMyPopCount(b, 100) }
func BenchmarkMyPopCount1000(b *testing.B)    { benchmarkMyPopCount(b, 1000) }
func BenchmarkMyPopCount10000(b *testing.B)   { benchmarkMyPopCount(b, 10000) }
func BenchmarkMyPopCount100000(b *testing.B)  { benchmarkMyPopCount(b, 100000) }
func BenchmarkMyPopCount1000000(b *testing.B) { benchmarkMyPopCount(b, 1000000) }

func BenchmarkDirtyPopCount10(b *testing.B)      { benchmarkDirtyPopCount(b, 10) }
func BenchmarkDirtyPopCount100(b *testing.B)     { benchmarkDirtyPopCount(b, 100) }
func BenchmarkDirtyPopCount1000(b *testing.B)    { benchmarkDirtyPopCount(b, 1000) }
func BenchmarkDirtyPopCount10000(b *testing.B)   { benchmarkDirtyPopCount(b, 10000) }
func BenchmarkDirtyPopCount100000(b *testing.B)  { benchmarkDirtyPopCount(b, 100000) }
func BenchmarkDirtyPopCount1000000(b *testing.B) { benchmarkDirtyPopCount(b, 1000000) }

func BenchmarkBitClearPopCount10(b *testing.B)      { benchmarkBitClearPopCount(b, 10) }
func BenchmarkBitClearPopCount100(b *testing.B)     { benchmarkBitClearPopCount(b, 100) }
func BenchmarkBitClearPopCount1000(b *testing.B)    { benchmarkBitClearPopCount(b, 1000) }
func BenchmarkBitClearPopCount10000(b *testing.B)   { benchmarkBitClearPopCount(b, 10000) }
func BenchmarkBitClearPopCount100000(b *testing.B)  { benchmarkBitClearPopCount(b, 100000) }
func BenchmarkBitClearPopCount1000000(b *testing.B) { benchmarkBitClearPopCount(b, 1000000) }

func benchmarkPopCount(b *testing.B, n int) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < n; j++ {
			PopCount(uint64(j))
		}
	}
}

func benchmarkMyPopCount(b *testing.B, n int) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < n; j++ {
			MyPopCount(uint64(j))
		}
	}
}

func benchmarkDirtyPopCount(b *testing.B, n int) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < n; j++ {
			DirtyPopCount(uint64(j))
		}
	}
}

func benchmarkBitClearPopCount(b *testing.B, n int) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < n; j++ {
			BitClearPopCount(uint64(j))
		}
	}
}

func TestPopCount(t *testing.T) {
	var tests = []struct {
		bits     uint64
		expected int
	}{
		{1, 1},
		{2, 1},
		{3, 2},
		{8, 1},
	}

	for _, test := range tests {
		if p := PopCount(test.bits); p != test.expected {
			t.Errorf("Actual: %v, Expected:  %v not much", p, test.expected)
		}
	}
}

func TestMyPopCount(t *testing.T) {
	var tests = []struct {
		bits     uint64
		expected int
	}{
		{1, 1},
		{2, 1},
		{3, 2},
		{8, 1},
	}

	for _, test := range tests {
		if p := MyPopCount(test.bits); p != test.expected {
			t.Errorf("Actual: %v, Expected:  %v not much", p, test.expected)
		}
	}
}

func TestDirtyPopCount(t *testing.T) {
	var tests = []struct {
		bits     uint64
		expected int
	}{
		{1, 1},
		{2, 1},
		{3, 2},
		{8, 1},
	}

	for _, test := range tests {
		if p := DirtyPopCount(test.bits); p != test.expected {
			t.Errorf("Actual: %v, Expected:  %v not much", p, test.expected)
		}
	}
}

func TestBitClearPopCount(t *testing.T) {
	var tests = []struct {
		bits     uint64
		expected int
	}{
		{1, 1},
		{2, 1},
		{3, 2},
		{8, 1},
	}

	for _, test := range tests {
		if p := BitClearPopCount(test.bits); p != test.expected {
			t.Errorf("Actual: %v, Expected:  %v not much", p, test.expected)
		}
	}
}
