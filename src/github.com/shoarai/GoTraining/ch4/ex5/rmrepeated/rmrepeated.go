// Copyright © 2016 shoarai
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Package rmrepeated removes repeated strings of string slice.
package rmrepeated

// RemoveRepeated rotates a slice of ints in place.
func RemoveRepeated(str []string) []string {
	if len(str) < 0 {
		return str
	}

	s := []string{str[0]}
	for _, v := range str[1:] {
		if s[len(s)-1] != v {
			s = append(s, v)
		}
	}
	return s
}
