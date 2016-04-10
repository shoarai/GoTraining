// Copyright © 2016 shoarai

// The max and min programs demonstrate a variadic function.
package main

import "fmt"

func maxZero(vals ...int) (int, bool) {
	if len(vals) == 0 {
		return 0, false
	}

	m := 0
	for _, val := range vals {
		if val > m {
			m = val
		}
	}
	return m, true
}

func minZero(vals ...int) (int, bool) {
	if len(vals) == 0 {
		return 0, false
	}

	m := 0
	for _, val := range vals {
		if val < m {
			m = val
		}
	}
	return m, true
}

func max(val int, vals ...int) int {
	m := val
	for _, v := range vals {
		if v > m {
			m = v
		}
	}
	return m
}

func min(val int, vals ...int) int {
	m := val
	for _, v := range vals {
		if v < m {
			m = v
		}
	}
	return m
}

func main() {
	// maxZero
	fmt.Println(maxZero())           //  "0"
	fmt.Println(maxZero(3))          //  "3"
	fmt.Println(maxZero(1, 2, 3, 4)) //  "4"

	values := []int{1, 2, 3, 4}
	fmt.Println(maxZero(values...)) // "4"

	// minZero
	fmt.Println(minZero())           //  "0"
	fmt.Println(minZero(3))          //  "3"
	fmt.Println(minZero(1, 2, 3, 4)) //  "1"

	fmt.Println(minZero(values...)) // "1"

	// max
	fmt.Println(max(3))          //  "3"
	fmt.Println(max(1, 2, 3, 4)) //  "4"

	// min
	fmt.Println(min(3))          //  "3"
	fmt.Println(min(1, 2, 3, 4)) //  "1"
}
