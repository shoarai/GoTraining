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
package ana

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", comma(os.Args[i]))
	}
}

// Anagram d
func Anagram() bool {
	return true
}

// comma inserts commas in a integer or decimal string.
func comma(s string) string {
	if s == "" {
		return s
	}

	// Write a sign
	var buf bytes.Buffer
	start := 0
	if s[0] == '+' || s[0] == '-' {
		buf.WriteByte(s[0])
		start++
	}

	// Write a integer part
	dot := strings.LastIndex(s, ".")
	if dot == -1 {
		dot = len(s)
	}
	buf.WriteString(commaInt(s[start:dot]))

	// Write a decimal part
	buf.WriteString(s[dot:])

	return buf.String()
}

// comma inserts commas in a non-negative decimal integer string.
func commaInt(s string) string {
	const digit = 3
	n := len(s)
	if n <= digit {
		return s
	}

	var buf bytes.Buffer
	for i, v := range s {
		if i != 0 && (n-i)%digit == 0 {
			buf.WriteString(",")
		}
		buf.WriteRune(v)
	}
	return buf.String()
}
