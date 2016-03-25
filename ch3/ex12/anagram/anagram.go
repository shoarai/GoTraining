// Copyright © 2016 shoarai
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Package anagram returns whether two arugments are anagram or not.
package anagram

// Anagram returns whether two arugments are anagram or not.
func Anagram(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	buf := []byte(s1)

	for _, v := range s2 {
		ana := false
		for i := 0; i < len(buf); i++ {
			if v == rune(buf[i]) {
				var ok bool
				buf, ok = remove(buf, i)
				if !ok {
					panic("index error of remove")
				}
				ana = true
				break
			}
		}
		if !ana {
			return false
		}
	}

	return true
}

func remove(slice []byte, i int) ([]byte, bool) {
	if i < 0 || i >= len(slice) {
		var s []byte
		return s, false
	}
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1], true
}
