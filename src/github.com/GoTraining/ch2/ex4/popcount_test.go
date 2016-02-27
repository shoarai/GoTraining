// Copyright © 2016 shoarai
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package popcount_test

import (
	"testing"

	"github.com/GoTraining/ch2/ex4"
)

// -- Test --

func TestPopCountLoop64(t *testing.T) {
	test := testPopCount(t, popcount.PopCountLoop64)
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

func BenchmarkPopCountLoop64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCountLoop64(0x1234567890ABCDEF)
	}
}

// 2.8GHz Intel Core i5
// $ go test -cpu=4 -bench=. ch2/ex4/popcount_test.go
// BenchmarkPopCountLoop64-4	20000000	       106 ns/op
