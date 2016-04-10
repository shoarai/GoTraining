// Copyright © 2016 shoarai

// The stringsJoin joins variadic texts.
package main

import (
	"fmt"
	"strings"
)

func stringsJoin(seq string, vals ...string) string {
	return strings.Join(vals, " ")
}

func main() {
	seq := " "
	fmt.Println(stringsJoin(seq))
	fmt.Println(stringsJoin(seq, "a", "b", "c"))

	str := []string{"あ", "い", "う"}
	fmt.Println(stringsJoin(seq, str...))
}
