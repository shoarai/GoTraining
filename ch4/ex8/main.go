// Copyright © 2016 shoarai
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Charcount computes counts of Unicode characters.
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {
	counts := make(map[rune]int)    // counts of Unicode characters
	types := make(map[string]int)   // counts of Unicode character's types
	var utflen [utf8.UTFMax + 1]int // count of lengths of UTF-8 encodings
	invalid := 0                    // count of invalid UTF-8 characters

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune() // returns rune, nbytes, error
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		counts[r]++
		utflen[n]++

		if unicode.IsControl(r) {
			types["control"]++
		}
		if unicode.IsDigit(r) {
			types["digit"]++
		}
		if unicode.IsGraphic(r) {
			types["graphic"]++
		}
		if unicode.IsLetter(r) {
			types["letter"]++
		}
		if unicode.IsLower(r) {
			types["lower"]++
		}
		if unicode.IsMark(r) {
			types["mark"]++
		}
		if unicode.IsNumber(r) {
			types["number"]++
		}
		if unicode.IsPrint(r) {
			types["print"]++
		}
		if unicode.IsPunct(r) {
			types["punct"]++
		}
		if unicode.IsSpace(r) {
			types["space"]++
		}
		if unicode.IsSymbol(r) {
			types["symbol"]++
		}
		if unicode.IsTitle(r) {
			types["title"]++
		}
		if unicode.IsUpper(r) {
			types["upper"]++
		}
	}
	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Print("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	fmt.Print("\nunicode\tcount\n")
	for k, v := range types {
		fmt.Printf("%s\t%d\n", k, v)
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}
