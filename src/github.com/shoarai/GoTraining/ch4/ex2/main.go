// Copyright © 2016 shoarai
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// The sha command computes the SHA hash (an array) of a string.
// Example:
// 	$ go build github.com/shoarai/GoTraining/ch4/ex2/main.go
//	$ ./main x
// 	SHA256: 2d711642b726b04401627ca9fbac32f5c8530fb1903cc4db02258717921a4881
//	$ ./main x -l=384
//
//	$ ./main x -l=512
//
//	$ ./main x -l=111
// 	Non support length
//
package main

import (
	"flag"
	"fmt"
)

import "crypto/sha256"

func main() {
	var len int
	flag.IntVar(&len, "l", 256, "Hash length")
	flag.Parse()

	for _, v := range flag.Args() {
		c, ok := sha(v, len)
		if ok {
			fmt.Printf("SHA%d: %x\n", len, c)
		} else {
			fmt.Println("No support length")
		}
	}
}

func sha(x string, len int) ([32]uint8, bool) {
	switch len {
	case 256:
		return sha256.Sum256([]byte(x)), true
	// case 384:
	// 	return sha512.Sum384([]byte(x)), true
	default:
		var c [32]uint8
		return c, false
	}
}
