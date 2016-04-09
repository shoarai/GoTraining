// Copyright © 2016 shoarai

// Expand replaces text that have prefix of $.
package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {
	s := input()
	s = expand(s, func(s string) string {
		return "@" + s + "@"
	})
	fmt.Println(s)
}

func expand(s string, f func(string) string) string {
	rep := regexp.MustCompile(`\$\w*`)
	return rep.ReplaceAllStringFunc(s, func(s string) string {
		return f(s[1:])
	})
}

func input() string {
	in := bufio.NewReader(os.Stdin)
	s, _, err := in.ReadLine()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
	return string(s)
}
