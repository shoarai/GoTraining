// Copyright © 2016 shoarai
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Package rmrepeated removes repeated strings of string slice.
package rmrepeated

// RemoveRepeated rotates a slice of ints in place.
func RemoveRepeated(str []string) []string {
	if len(str) < 1 {
		return str
	}

	for i := 1; i < len(str); i++ {
		if str[i-1] == str[i] {
			str, _ = remove(str, i)
		}
	}
	return str
}

func remove(slice []string, i int) ([]string, bool) {
	if i < 0 || i >= len(slice) {
		var s []string
		return s, false
	}
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1], true
}
