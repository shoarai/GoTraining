// Copyright © 2016 shoarai
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Comma prints its argument numbers with a comma at each power of 1000.
//
// Example:
// 	$ go build gopl.io/ch3/comma
//	$ ./comma 1 12 123 1234 1234567890
// 	1
// 	12
// 	123
// 	1,234
// 	1,234,567,890
//
package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", comma(os.Args[i]))
	}
}

// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
	digit := 3
	n := len(s)
	if n <= digit {
		return s
	}

	var buf bytes.Buffer
	for i := 0; i < n; i++ {
		if (n-i)%digit == 0 {
			buf.WriteString(",")
		}
		buf.WriteByte(s[i])
	}
	return buf.String()
}
