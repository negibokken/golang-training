package intset

import (
	"fmt"
	"sort"
	"strings"
	"testing"

	popcount "github.com/negibokken/golang-training/ch11/ex02/popcount"
)

func countBits(words []uint64) int {
	var num int
	for _, word := range words {
		num += popcount.PopCount(word)
	}
	return num
}

func isEqualSet(x IntSet, embededSet map[int]struct{}) bool {
	if countBits(x.words) != len(embededSet) {
		return false
	}
	var keys []int
	for k := range embededSet {
		keys = append(keys, k)
	}
	for _, k := range keys {
		if !x.Has(k) {
			fmt.Printf("x does not have %v", k)
			return false
		}
	}
	return true
}

func keyToString(ma map[int]struct{}) string {
	var keys []int
	for k := range ma {
		fmt.Printf(string(k))
		keys = append(keys, k)
	}
	sort.Ints(keys)
	str := fmt.Sprintf("%v", keys)
	str = strings.Replace(str, "[", "{", -1)
	str = strings.Replace(str, "]", "}", -1)
	return str
}

func TestIntSet(t *testing.T) {
	var tests = []struct {
		name string
		nums []int
	}{
		{
			"test IntSet", []int{1, 3, 4, 2, 5},
		},
		{
			"Add duplicated element into IntSet", []int{1, 3, 3, 3, 3, 4, 2, 5},
		},
		{
			"Add large and many elements into IntSet",
			[]int{1, 2, 3, 4, 5, 1024, 2048, 4096, 8192, 16384, 32768, 65536},
		},
		{
			"test Empty IntSet", []int{},
		},
	}
	for _, test := range tests {
		var x IntSet
		embededSet := make(map[int]struct{})
		// Add nums into IntSet and embeded Set by using map
		for _, num := range test.nums {
			x.Add(num)
			embededSet[num] = struct{}{}
		}
		if !isEqualSet(x, embededSet) {
			t.Errorf("Finally Set is %v want, %v", x.String(), keyToString(embededSet))
		}
	}
}
