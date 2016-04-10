// Copyright © 2016 shoarai

// Defer demonstrates a deferred call being invoked during a panic.
package main

import "fmt"

func main() {
	fmt.Printf("%d\n", f())
}

func f() (n int) {
	defer func() {
		if err := recover(); err != nil {
			n = 10
		}
	}()
	panic("")
}
