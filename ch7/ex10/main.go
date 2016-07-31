// Copyright © 2016 shoarai

// Sorting sorts a music playlist into a variety of orders.
package main

import (
	"fmt"
	"sort"
)

func main() {
	text := Text{[]byte("abcdefg")}
	palindorome := isPalindorome(text)
	fmt.Printf("%s palindorome is %t\n", text, palindorome)

	text = Text{[]byte("abccba")}
	palindorome = isPalindorome(text)
	fmt.Printf("%s palindorome is %t\n", text, palindorome)
}

func isPalindorome(s sort.Interface) bool {
	for i := 0; i < s.Len()/2; i++ {
		j := s.Len() - i - 1
		if s.Less(i, j) || s.Less(i, j) {
			return false
		}
	}
	return true
}

type Text struct {
	bytes []byte
}

func (x Text) Len() int           { return len(x.bytes) }
func (x Text) Less(i, j int) bool { return x.bytes[i] < x.bytes[j] }
func (x Text) Swap(i, j int)      { x.bytes[i], x.bytes[j] = x.bytes[j], x.bytes[i] }
