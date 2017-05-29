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
	fmt.Println("add 1")
	s.Add(1)
	fmt.Println(s.String())
	fmt.Println(s.Len())
	fmt.Println("remove 1")
	s.Remove(1)
	fmt.Println(s.String())
	fmt.Println("add 1")
	s.Add(1)
	fmt.Println("clear")
	fmt.Println(s.String())
	fmt.Println("add 1")
	s.Add(1)
	fmt.Println("add 2")
	s.Add(2)
	fmt.Println("copy")
	t := s.Copy()
	fmt.Println("copy s -> t")
	fmt.Println(t.String())
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
