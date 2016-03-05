// Copyright © 2016 shoarai
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Package comma prints its argument numbers with a comma at each power of 1000.
package comma

import (
	"bytes"
	"strings"
)

// Comma inserts commas in a integer or decimal string.
func Comma(s string) string {
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
