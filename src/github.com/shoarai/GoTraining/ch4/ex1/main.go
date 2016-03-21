// Copyright © 2016 shoarai
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Compatebit computes the number of different bit of
// the two SHA256 hashs (an array) of a string.
// Example:
// $ go build github.com/shoarai/GoTraining/ch4/ex1/main.go
// $ ./main.go
// Input value 1: x
// Input value 2: X
// Hash1: 2d711642b726b04401627ca9fbac32f5c8530fb1903cc4db02258717921a4881
// Hash2: 4b68ab3847feda7d6c62c1fbcbeebfa35eab7351ed5e78f4ddadea5df64b8015
// Diff bit: 225
package main

import (
	"fmt"

	"github.com/shoarai/GoTraining/ch4/ex1/comparebit"
)

import "crypto/sha256"

func main() {
	var v1 string
	var v2 string
	fmt.Printf("Input value 1: ")
	fmt.Scan(&v1)
	fmt.Printf("Input value 2: ")
	fmt.Scan(&v2)

	c1 := sha256.Sum256([]byte(v1))
	c2 := sha256.Sum256([]byte(v2))
	fmt.Printf("Hash1: %x\nHash2: %x\nDiff bit: %d\n",
		c1, c2, comparebit.DiffBitCount(&c1, &c2))
}
