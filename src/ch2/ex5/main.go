// Copyright © 2016 shoarai
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Package popcount performs number of bits count.
package popcount

// PopCountLoop64 returns the population count (number of set bits) of x.
func PopCountByLowestBit(x uint64) int {
	var count int
	for x != 0 {
		count++
		x = x & (x - 1)
	}
	return count
}
