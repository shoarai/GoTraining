// Copyright © 2016 shoarai
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Package rmrepeated removes repeated strings of string slice.
package rmrepeated

import "unicode"

// ConvSpaces replace unicode spaces into utf8 space.
func ConvSpaces(b []byte) []byte {
	if len(b) <= 1 {
		return b
	}

	rs := []rune(string(b))
	var count int
	for i := 0; i < len(rs); i++ {
		if unicode.IsSpace(rs[i]) {
			count++
		} else {
			if count >= 2 {
				rs[i-1] = ' '
				rs, _ = remove(rs, i-count, i-2)
				i -= count
			}
			count = 0
		}
	}

	return []byte(string(rs))
}

func remove(slice []rune, start, end int) ([]rune, bool) {
	if start < 0 || end >= len(slice) {
		return slice, false
	}
	copy(slice[start:], slice[end+1:])
	return slice[:len(slice)-1], true
}
