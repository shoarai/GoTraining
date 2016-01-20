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
	GetOSArgsInefficient()
	timeLong := time.Since(start)
	start = time.Now()
	GetOSArgs()
	timeShort := time.Since(start)
	fmt.Println(timeLong, " - ", timeShort, " = ", timeLong-timeShort)
}

func GetOSArgsInefficient() string {
	var s, sep string
	for i := 0; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	return s
}

func GetOSArgs() string {
	return strings.Join(os.Args[0:], " ")
}
