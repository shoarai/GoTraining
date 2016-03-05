// Copyright © 2016 shoarai
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Anagram return whether or not are two arugments anagram.
//
// Example:
// 	$ go build github.com/shoarai/GoTraining/ch3/ex12/main.go
//	$ ./main text ttxe
// 	anagram
//	$ ./main text tteo
// 	non anagram
//
package main

import (
	"fmt"
	"os"

	"github.com/shoarai/GoTraining/ch3/ex12/anagram"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Input two argumens")
		return
	}

	if anagram.Anagram(os.Args[1], os.Args[2]) {
		fmt.Println("anagram")
	} else {
		fmt.Println("non anagram")
	}
}
