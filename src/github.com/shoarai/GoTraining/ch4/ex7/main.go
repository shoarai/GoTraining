// Copyright © 2016 shoarai
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Rev reverses a slice.
package main

import (
	"fmt"

	"github.com/shoarai/GoTraining/ch4/ex7/rev"
)

func main() {
	//!+array
	a := [...]byte{0, 1, 2, 3, 4, 5}
	rev.Reverse(a[:])
	fmt.Println(a) // "[5 4 3 2 1 0]"
	//!-array

	//!+slice
	s := []byte{0, 1, 2, 3, 4, 5}
	// Rotate s left by two positions.
	rev.Reverse(s[:2])
	rev.Reverse(s[2:])
	rev.Reverse(s)
	fmt.Println(s) // "[2 3 4 5 0 1]"
	//!-slice
}
