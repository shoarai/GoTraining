// Copyright © 2016 shoarai
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Rev reverses a slice.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	//!+array
	a := [6]int{0, 1, 2, 3, 4, 5}
	reverse(&a, 0, 5)
	fmt.Println(a) // "[5 4 3 2 1 0]"
	//!-array

	//!+slice
	s := []int{0, 1, 2, 3, 4, 5}
	// Rotate s left by two positions.
	var sa [6]int
	copy(sa[:], s[:6])
	reverse(&sa, 0, 1)
	reverse(&sa, 2, 5)
	reverse(&sa, 0, 5)
	s = sa[:]
	fmt.Println(s) // "[2 3 4 5 0 1]"
	//!-slice

	// Interactive test of reverse.
	input := bufio.NewScanner(os.Stdin)
outer:
	for input.Scan() {
		var ints []int
		for _, s := range strings.Fields(input.Text()) {
			x, err := strconv.ParseInt(s, 10, 64)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				continue outer
			}
			ints = append(ints, int(x))
		}
		l := len(ints)
		if l > 6 {
			fmt.Println("Input array in length of equal or smaller than six")
			continue outer
		}
		var inta [6]int
		copy(inta[:], ints[:l])
		reverse(&inta, 0, l-1)
		ints = inta[:l]
		fmt.Printf("%v\n", ints)
	}
	// NOTE: ignoring potential errors from input.Err()
}

// reverse reverses a array of ints in place.
func reverse(s *[6]int, start, end int) {
	for i, j := start, end; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
