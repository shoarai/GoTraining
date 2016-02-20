// Copyright © 2016 shoarai
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package popcount_test

import (
	"testing"

	"ch2/ex5"
)

// -- Test --

func TestPopCountByLowestBit(t *testing.T) {
	test := testPopCount(t, popcount.PopCountByLowestBit)
	test(0x0, 0)
	test(0xffff, 16)
	test(0xfffefffe, 30)
	test(0xffffffffffffffff, 64)
}

func testPopCount(t *testing.T, popCount func(uint64) int) func(val uint64, count int) {
	return func(val uint64, count int) {
		c := popCount(val)
		if c != count {
			t.Errorf("PopCount(%x) = %d, want %d", val, c, count)
		}
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
