// Copyright © 2016 shoarai
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Echo prints the index and value of each of
// its command-line arguments, one per line.
package main

import (
	"fmt"
	"os"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Println(i, os.Args[i])
	}

	v := 1

	if v > 0 {
		fmt.Println("ffffff1f")

	} else if v > 0 {
		fmt.Println("fffffff")

	}

	switch true {
	case v > 0:
		fmt.Println("eee")
	case 1 > 0:
		fmt.Println("aaa")
	default:
		fmt.Println("fffffff")
	}
}
