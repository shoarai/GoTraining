// Copyright © 2016 shoarai

// charcount computes counts of Unicode characters.
package charcount_test

import (
	"fmt"
	"strings"
	"testing"
	"unicode/utf8"

	"github.com/shoarai/GoTraining/ch11/ex1/charcount"
)

func TestCountChar(t *testing.T) {
	tests := []struct {
		input string
		want  charcount.CharCount
	}{
		{"a", charcount.CharCount{
			map[rune]int{'a': 1},
			[utf8.UTFMax + 1]int{0, 1},
			0},
		},
		{"あいあ", charcount.CharCount{
			map[rune]int{'あ': 2, 'い': 1},
			[utf8.UTFMax + 1]int{0, 0, 0, 3},
			0},
		},
	}

	for _, test := range tests {
		reader := strings.NewReader(test.input)
		charCount := charcount.CountChar(reader)

		if !isSameCharCount(charCount, test.want) {
			t.Errorf("CountChar(%s) = %v", test.input, charCount)
			printCharCount(charCount)
		}
	}
}

func isSameCharCount(charCount1, charCount2 charcount.CharCount) bool {
	for c, n1 := range charCount1.Counts {
		if n2, ok := charCount2.Counts[c]; !ok || n1 != n2 {
			return false
		}
	}
	if charCount1.UTFlen != charCount2.UTFlen {
		return false
	}
	if charCount1.Invalid != charCount2.Invalid {
		return false
	}
	return true
}

func printCharCount(charCount charcount.CharCount) {
	fmt.Printf("rune\tcount\n")
	for c, n := range charCount.Counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Print("\nlen\tcount\n")
	for i, n := range charCount.UTFlen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	if charCount.Invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", charCount.Invalid)
	}
}
