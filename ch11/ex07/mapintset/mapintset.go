package mapintset

import (
	"bytes"
	"fmt"
	"sort"
)

type MapIntSet struct {
	m map[int]bool
}

func NewMapIntSet() *MapIntSet {
	return &MapIntSet{map[int]bool{}}
}

func (s *MapIntSet) Has(x int) bool {
	return s.m[x]
}

func (s *MapIntSet) Add(x int) {
	s.m[x] = true
}

func (s *MapIntSet) AddAll(nums ...int) {
	for _, x := range nums {
		s.m[x] = true
	}
}

func (s *MapIntSet) Len() int {
	return len(s.m)
}

func (s *MapIntSet) Remove(x int) {
	delete(s.m, x)
}

func mapKeys(m map[int]bool) []int {
	var keys []int
	for k := range m {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	return keys
}

func (s *MapIntSet) UnionWith(t MapIntSet) {
	keys := mapKeys(t.m)
	for _, k := range keys {
		s.m[k] = true
	}
}

func (s *MapIntSet) Clear() {
	s.m = make(map[int]bool)
}

func (s *MapIntSet) Copy() *MapIntSet {
	copy := make(map[int]bool)
	for k, v := range s.m {
		copy[k] = v
	}
	return &MapIntSet{copy}
}

func (s *MapIntSet) Ints() []int {
	ints := make([]int, 0, len(s.m))
	for x := range s.m {
		ints = append(ints, x)
	}
	sort.Ints(ints)
	return ints
}

func (s *MapIntSet) String() string {
	b := &bytes.Buffer{}
	b.WriteByte('{')
	for i, x := range s.Ints() {
		if i != 0 {
			b.WriteByte(' ')
		}
		fmt.Fprintf(b, "%d", x)
	}
	b.WriteByte('}')
	return b.String()
}
