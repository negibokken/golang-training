package main

import (
	"bytes"
	"fmt"

	"./popcount"
)

// IntSet is a set that doesn't have minus value
type IntSet struct {
	words []uint64
}

func main() {
	var s IntSet
	fmt.Println("add 1,2,3,4,5")
	s.AddAll(1, 2, 3, 4, 5)
	ss := s.Elem()
	for i, ele := range ss {
		fmt.Println(i, ele)
	}
}

// Has check in the set
func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add add new element
func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)

	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// Len returns the length of set
func (s *IntSet) Len() int {
	result := 0
	for _, w := range s.words {
		result += popcount.PopCount(w)
	}
	return result
}

// Remove deletes the element in the set
func (s *IntSet) Remove(x int) {
	word, bit := x/64, uint(x%64)
	if word < len(s.words) {
		s.words[word] &^= 1 << bit
	}
}

// Clear deletes the set
func (s *IntSet) Clear() {
	for i := range s.words {
		s.words[i] = 0
	}
}

// Copy generate copy of the set
func (s *IntSet) Copy() *IntSet {
	result := &IntSet{}
	for _, word := range s.words {
		result.words = append(result.words, word)
	}
	return result
}

// UnionWith add s and t
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// String returns string of the IntSet
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

// AddAll adds elements in the set
func (s *IntSet) AddAll(num ...int) {
	for _, n := range num {
		s.Add(n)
	}
}

// IntersectWith make intersect set
func (s *IntSet) IntersectWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] &= tword
		}
	}
	for i := len(t.words); i < len(s.words); i++ {
		s.words[i] = 0
	}
}

// DifferenceWith make different set
func (s *IntSet) DifferenceWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] &^= tword
		}
	}
}

// SymmetricDifference make symmetric difference set
func (s *IntSet) SymmetricDifference(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] ^= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// Elem returns slice of int
func (s *IntSet) Elem() (result []uint64) {
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				result = append(result, uint64(64*i+j))
			}
		}
	}
	return
}
