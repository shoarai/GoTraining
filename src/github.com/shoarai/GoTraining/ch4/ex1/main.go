// Copyright © 2016 shoarai
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Compatebit computes the number of different bit of
// the two SHA256 hashs (an array) of a string.
package main

import (
	"fmt"

	"github.com/shoarai/GoTraining/ch4/ex1/comparebit"
)

import "crypto/sha256"

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Printf("%x\n%x\n%t\n%T\n%d\n",
		c1, c2, c1 == c2, c1, comparebit.DiffBitCount(&c1, &c2))
	// Output:
	// 2d711642b726b04401627ca9fbac32f5c8530fb1903cc4db02258717921a4881
	// 4b68ab3847feda7d6c62c1fbcbeebfa35eab7351ed5e78f4ddadea5df64b8015
	// false
	// [32]uint8
	// 225
}
