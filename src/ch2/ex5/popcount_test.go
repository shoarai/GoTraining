// Copyright © 2016 shoarai
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package popcount_test

import (
	"testing"

	"ch2/ex5"
)

// -- Test --

func TestPopCountByLowestBit(t *testing.T) {
	count := popcount.PopCountByLowestBit(0)
	if count != 0 {
		t.Error(count, "out!")
	}

	count = popcount.PopCountByLowestBit(255)
	if count != 8 {
		t.Error(count, "out!")
	}
}

// -- Benchmarks --

func BenchmarkPopCountByLowestBit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCountByLowestBit(0x1234567890ABCDEF)
	}
}

// 2.8GHz Intel Core i5
// $ go test -cpu=4 -bench=. ch2/ex5/popcount_test.go
// BenchmarkPopCountByLowestBit-4	50000000	        25.6 ns/op