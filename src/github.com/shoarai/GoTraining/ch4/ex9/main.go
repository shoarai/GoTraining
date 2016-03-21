// Copyright © 2016 shoarai
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// wordfeq computes counts of Unicode words.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args[1:]) == 0 {
		fmt.Println("Input a file path as argument")
		return
	}

	path := os.Args[1]
	f, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	counts := make(map[string]int)
	input := bufio.NewScanner(f)
	input.Split(bufio.ScanWords)

	for input.Scan() {
		text := strings.ToLower(input.Text())
		counts[text]++
	}

	fmt.Printf("count\tword\n")
	for c, n := range counts {
		fmt.Printf("%d\t%q\n", n, c)
	}
}
