// Copyright © 2016 shoarai

// Echo1 prints the index and value of each of
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
}