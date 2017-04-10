package popcount

import "testing"

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(uint64(i))
	}
}

func BenchmarkMyPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MyPopCount(uint64(i))
	}
}

func BenchmarkDirtyPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DirtyPopCount(uint64(i))
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
