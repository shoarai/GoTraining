// Copyright © 2016 shoarai
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Package comparebit compares bits
package comparebit

// DiffBitCount compare bits
func DiffBitCount(x, y *[32]byte) int {
	var count int

	for i := 0; i < 32; i++ {
		for j := 0; j < 8; j++ {
			if int(x[i])>>uint(j) != int(y[i])>>uint(j) {
				count++
			}
		}
	}
	return count
}
