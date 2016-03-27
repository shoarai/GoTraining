// Copyright © 2016 shoarai
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package comparebit_test

import (
	"testing"

	"../comparebit"
)

func TestDiffBitCount(t *testing.T) {
	var b0 [32]byte
	for i := range b0 {
		b0[i] = byte(0)
	}

	test(t, &b0, &b0, 0)

	var b1 [32]byte
	for i := range b1 {
		b1[i] = byte(1)
	}

	test(t, &b0, &b1, 32)
}

func test(t *testing.T, x, y *[32]byte, count int) {
	c := comparebit.DiffBitCount(x, y)
	if c != count {
		t.Errorf("CompareBit(%x, %x) = %d, want %d\n", *x, *y, c, count)
	}
}
