// Copyright © 2016 shoarai

// Package intset provides a set of integers based on a bit vector.
package intset

// An IntSet is a set of small non-negative integers.
// Its zero value represents the empty set.
type IntSet struct {
	words []uint64
}

// Elems returs slices of the set.
func (s *IntSet) Elems() []bool {
	var slice []bool
	for _, v := range s.words {
		for j := uint(0); j < 64; j++ {
			b := (v>>j)&1 == 1
			slice = append(slice, b)
		}
	}
	return slice
}

// Add adds the non-negative value x to the set.
func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= (1 << bit)
}
