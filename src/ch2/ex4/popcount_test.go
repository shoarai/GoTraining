// Copyright © 2016 shoarai
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package popcount_test

import (
	"testing"

	"ch2/ex4"
)

// -- Test --

func TestPopCountLoop64(t *testing.T) {
	count := popcount.PopCountLoop64(0)
	if count != 0 {
		t.Error(count, "out!")
	}

	count = popcount.PopCountLoop64(255)
	if count != 8 {
		t.Error(count, "out!")
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
