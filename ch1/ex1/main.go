// Copyright © 2016 shoarai

// Echo prints its command-line name and arguments.
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(strings.Join(os.Args[0:], " "))
}
