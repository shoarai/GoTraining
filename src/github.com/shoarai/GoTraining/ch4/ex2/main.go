// Copyright © 2016 shoarai
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// The sha command computes the SHA hash (an array) of a string.
// Example:
// $ go build github.com/shoarai/GoTraining/ch4/ex2/main.go
// $ ./main x
// SHA256:
//  2d711642b726b04401627ca9fbac32f5c8530fb1903cc4db02258717921a4881
// $ ./main x -l=384
// SHA384:
//  d752c2c51fba0e29aa190570a9d4253e44077a058d3297fa3a5630d5bd012622
//  f97c28acaed313b5c83bb990caa7da85
// $ ./main x -l=512
// SHA512:
//  a4abd4448c49562d828115d13a1fccea927f52b4d5459297f8b43e42da89238b
//  c13626e43dcb38ddb082488927ec904fb42057443983e88585179d50551afe62
// $ ./main x -l=111
// Non support length
//
package main

import (
	"flag"
	"fmt"
)

import (
	"crypto/sha256"
	"crypto/sha512"
)

func main() {
	var len int
	flag.IntVar(&len, "l", 256, "Hash length")
	flag.Parse()

	for _, v := range flag.Args() {
		switch len {
		case 256:
			fmt.Printf("SHA%d: %x\n", len, sha256.Sum256([]byte(v)))
		case 384:
			fmt.Printf("SHA%d: %x\n", len, sha512.Sum384([]byte(v)))
		case 512:
			fmt.Printf("SHA%d: %x\n", len, sha512.Sum512([]byte(v)))
		default:
			fmt.Println("No support length")
		}
	}
}
