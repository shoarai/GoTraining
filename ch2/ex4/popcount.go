// Copyright © 2016 shoarai
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Package popcount counts number of bits.
package popcount

// PopCountLoop64 returns the population count (number of set bits) of x.
func PopCountLoop64(x uint64) int {
	var count byte
	for i := uint(0); i < 64; i++ {
		count += byte((x >> i) & 1)
	}
	return int(count)
}
