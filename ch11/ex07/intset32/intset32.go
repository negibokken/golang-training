package intset32

import (
	"bytes"
	"fmt"
)

type IntSet interface {
	Has(x int) bool
	Add(x int)
	AddAll(nums ...int)
	UnionWith(t IntSet)
	Len() int
	Remove(x int)
	Clear()
	Copy() IntSet
	String() string
	Ints() []int
}

type BitIntSet struct {
	words []uint32
}

func NewBitIntSet() IntSet {
	return &BitIntSet{}
}

func (s *BitIntSet) Has(x int) bool {
	word, bit := x/32, uint(x%32)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

func (s *BitIntSet) Add(x int) {
	word, bit := x/32, uint(x%32)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

func (s *BitIntSet) AddAll(nums ...int) {
	for _, n := range nums {
		s.Add(n)
	}
}

func (s *BitIntSet) UnionWith(t IntSet) {
	if bis, ok := t.(*BitIntSet); ok {
		for i, tword := range bis.words {
			if i < len(s.words) {
				s.words[i] |= tword
			} else {
				s.words = append(s.words, tword)
			}
		}
	} else {
		for _, i := range t.Ints() {
			s.Add(i)
		}
	}
}

func popcount(x uint32) int {
	count := 0
	for x != 0 {
		count++
		x &= x - 1
	}
	return count
}

func (s *BitIntSet) Len() int {
	count := 0
	for _, word := range s.words {
		count += popcount(word)
	}
	return count
}

func (s *BitIntSet) Remove(x int) {
	word, bit := x/32, uint(x%32)
	s.words[word] &^= 1 << bit
}

func (s *BitIntSet) Clear() {
	for i := range s.words {
		s.words[i] = 0
	}
}

func (s *BitIntSet) Copy() IntSet {
	new := &BitIntSet{}
	new.words = make([]uint32, len(s.words))
	copy(new.words, s.words)
	return new
}

func (s *BitIntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 32; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 32*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

func (s *BitIntSet) Ints() []int {
	var ints []int
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 32; j++ {
			if word&(1<<uint(j)) != 0 {
				ints = append(ints, 32*i+j)
			}
		}
	}
	return ints
}