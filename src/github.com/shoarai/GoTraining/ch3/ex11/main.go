// Copyright © 2016 shoarai
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Comma prints its argument numbers with a comma at each power of 1000.
//
// Example:
// 	$ go build github.com/shoarai/GoTraining/ch3/ex11/main.go
//	$ ./main 1 -12 12.3 -1234 12345.67890
// 	1
// 	-12
// 	12.3
// 	-1,234
// 	12,345.67890
//
package main

import (
	"fmt"
	"os"

	"github.com/shoarai/GoTraining/ch3/ex11/comma"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", comma.Comma(os.Args[i]))
	}
}
