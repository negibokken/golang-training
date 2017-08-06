package mapintset

import (
	"math/rand"
	"testing"
)

var intset = &MapIntSet{make(map[int]bool)}
var intsetb = &MapIntSet{make(map[int]bool)}

const max = 100000

func addData(n int) {
	for i := 0; i < n; i++ {
		intset.Add(rand.Intn(max))

	}
}

func addDataB(n int) {
	for i := 0; i < n; i++ {
		intsetb.Add(rand.Intn(max))

	}
}

func benchmarkHas(b *testing.B, n int) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < n; j++ {
			intset.Has(rand.Intn(max))
		}
	}
}

func BenchmarkHas100(b *testing.B)   { benchmarkHas(b, 100) }
func BenchmarkHas1000(b *testing.B)  { benchmarkHas(b, 1000) }
func BenchmarkHas10000(b *testing.B) { benchmarkHas(b, 10000) }

func benchmarkAdd(b *testing.B, n int) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < n; j++ {
			intset.Add(rand.Intn(max))
		}
	}
	intset.Clear()
}

func BenchmarkAdd100(b *testing.B)   { benchmarkAdd(b, 100) }
func BenchmarkAdd1000(b *testing.B)  { benchmarkAdd(b, 1000) }
func BenchmarkAdd10000(b *testing.B) { benchmarkAdd(b, 10000) }

func randInts(n int) []int {
	ints := make([]int, n)
	for i := 0; i < n; i++ {
		ints[i] = rand.Intn(max)
	}
	return ints
}

func benchmarkAddAll(b *testing.B, n int) {
	for i := 0; i < b.N; i++ {
		nums := randInts(n)
		for j := 0; j < n; j++ {
			intset.AddAll(nums...)
		}
	}
	intset.Clear()
}

func BenchmarkAddAll100(b *testing.B)   { benchmarkAddAll(b, 100) }
func BenchmarkAddAll1000(b *testing.B)  { benchmarkAddAll(b, 1000) }
func BenchmarkAddAll10000(b *testing.B) { benchmarkAddAll(b, 10000) }

func benchmarkUnionWith(b *testing.B, n int) {
	addData(n)
	addDataB(n)
	for i := 0; i < b.N; i++ {
		intset.UnionWith(*intsetb)
	}
}

func BenchmarkUnionWith100(b *testing.B)   { benchmarkUnionWith(b, 100) }
func BenchmarkUnionWith1000(b *testing.B)  { benchmarkUnionWith(b, 1000) }
func BenchmarkUnionWith10000(b *testing.B) { benchmarkUnionWith(b, 10000) }

func benchmarkString(b *testing.B, n int) {
	addData(n)
	for i := 0; i < b.N; i++ {
		intset.String()
	}
}

func BenchmarkString100(b *testing.B)   { benchmarkString(b, 100) }
func BenchmarkString1000(b *testing.B)  { benchmarkString(b, 1000) }
func BenchmarkString10000(b *testing.B) { benchmarkString(b, 10000) }
