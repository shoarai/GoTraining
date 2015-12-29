// Copyright © 2016 shoarai

// Dup2 prints the count and text of lines that appear more than once
// in the input and the names of files in which the text appears.
// It reads from stdin or from a list of named files.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Line struct {
	Count int
	Files []string
}

func main() {
	lines := make(map[string]Line)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, lines, "")
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, lines, arg)
			f.Close()
		}
	}
	for line, n := range lines {
		if n.Count > 1 {
			fmt.Printf("%d\t%s\t%s\n", n.Count, line, strings.Join(n.Files, " "))
		}
	}
}

func countLines(f *os.File, lines map[string]Line, filename string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		count := lines[input.Text()]
		count.Count++
		if !stringInStrings(filename, count.Files) {
			count.Files = append(count.Files, filename)
		}
		lines[input.Text()] = count;
	}
	// NOTE: ignoring potential errors from input.Err()
}

func stringInStrings(str string, strs []string) bool {
 	for _, s := range strs {
 		if s == str {
 			return true
 		}
 	}
 	return false
 }
