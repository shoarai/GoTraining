// Copyright © 2016 shoarai
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Package popcount performs number of bits count.
package popcount

// PopCountLoop64 returns the population count (number of set bits) of x.
func PopCountLoop64(x uint64) int {
	var count byte
	var i uint
	for ; i < 64; i++ {
		count += byte((x >> i) & 1)
	}
	return int(count)
}
