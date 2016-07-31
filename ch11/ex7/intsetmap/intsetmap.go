// Copyright © 2016 shoarai

// Package intsetmap provides a set of integers based on a map.
package intsetmap

import (
	"bytes"
	"strconv"
)

// An IntSet is a set of small non-negative integers.
// Its zero value represents the empty set.
type IntSet struct {
	words map[int]bool
}

// Has reports whether the set contains the non-negative value x.
func (s *IntSet) Has(x int) bool {
	_, ok := s.words[x]
	return ok
}

// Add adds the non-negative value x to the set.
func (s *IntSet) Add(x int) {
	if len(s.words) == 0 {
		s.words = make(map[int]bool)
	}
	s.words[x] = true
}

// UnionWith sets s to the union of s and t.
func (s *IntSet) UnionWith(t *IntSet) {
	for n, b := range t.words {
		s.words[n] = b
	}
}

// String returns the set as a string of the form "{1 2 3}".
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	first := true
	for n := range s.words {
		if first {
			buf.WriteByte(' ')
			first = false
		}
		buf.WriteString(strconv.Itoa(n))
	}
	buf.WriteByte('}')
	return buf.String()
}
