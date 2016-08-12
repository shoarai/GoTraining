// Copyright © 2016 shoarai

// charcount computes counts of Unicode characters.
package charcount

import (
	"bufio"
	"io"
	"unicode"
	"unicode/utf8"
)

type CharCount struct {
	Counts  map[rune]int
	UTFlen  [utf8.UTFMax + 1]int
	Invalid int
}

func CountChar(r io.Reader) (*CharCount, error) {
	counts := make(map[rune]int)    // counts of Unicode characters
	var utflen [utf8.UTFMax + 1]int // count of lengths of UTF-8 encodings
	invalid := 0                    // count of invalid UTF-8 characters

	in := bufio.NewReader(r)
	for {
		r, n, err := in.ReadRune() // returns rune, nbytes, error
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		counts[r]++
		utflen[n]++
	}

	return &CharCount{
		counts, utflen, invalid,
	}, nil
}
