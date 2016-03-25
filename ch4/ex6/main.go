// Copyright © 2016 shoarai
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// RemoveRepeated removes repeated strings of string slice.
// Example:
// $ go buid github.com/shoarai/GoTraining/ch4/ex5/main.go
// $ ./main test test test test
// test
// $ ./main text test text text
// text test text
package main

import (
	"fmt"
	"unicode"
)

func main() {
	var b = []byte{'0', ' ', ' ', '0'}
	fmt.Println(b)
	for _, v := range b {
		if unicode.IsSpace(rune(v)) {
			fmt.Println(v)
		}
	}
	// if len(os.Args) < 2 {
	// 	return
	// }
	// fmt.Println(rmrepeated.RemoveRepeated(os.Args[1:]))
}
