// Copyright © 2016 shoarai
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Echo prints the difference in running time
// between the inefficient version and the one
package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	getOSArgsInefficient()
	timeLong := time.Since(start)
	start = time.Now()
	getOSArgs()
	timeShort := time.Since(start)
	fmt.Println(timeLong, " - ", timeShort, " = ", timeLong-timeShort)
}

func getOSArgsInefficient() string {
	var s, sep string
	for i := 0; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	return s
}

func getOSArgs() string {
	return strings.Join(os.Args[0:], " ")
}
