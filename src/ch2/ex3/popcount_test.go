// Copyright © 2016 shoarai
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package popcount_test

import (
	"testing"

	"ch2/ex3"
)

// -- Tests --

func TestPopCount(t *testing.T) {
	testPopCount(t, popcount.PopCount, 0x0, 0)
	testPopCount(t, popcount.PopCount, 0xffff, 16)
	testPopCount(t, popcount.PopCount, 0xfffefffe, 30)
	testPopCount(t, popcount.PopCount, 0xffffffffffffffff, 64)
}

func testPopCount(t *testing.T, popCount func(uint64) int, val uint64, count int) {
	c := popCount(val)
	if c != count {
		t.Errorf("PopCount(%x) = %d, want %d", val, c, count)
	}
}

// -- Benchmarks --

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCount(0x1234567890ABCDEF)
	}
}

func BenchmarkPopCountLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCountLoop(0x1234567890ABCDEF)
	}
}

// 2.8GHz Intel Core i5
// $ go test -cpu=4 -bench=. ch2/ex3/popcount_test.go
// BenchmarkPopCount-4    	300000000	        6.19 ns/op
// BenchmarkPopCountLoop-4	100000000	        15.5 ns/op
