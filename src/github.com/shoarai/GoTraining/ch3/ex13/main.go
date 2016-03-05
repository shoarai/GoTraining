// Copyright © 2016 shoarai
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Byte demonstrates byte size.
package main

import "fmt"

const (
	_ = 1 << (10 * iota)
	KB
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

func main() {
	fmt.Println("1KB =", KB, "Byte")
	fmt.Println("1MB =", MB, "Byte")
	fmt.Println("1GB =", GB, "Byte")
	fmt.Println("1TB =", TB, "Byte")
	fmt.Println("1PB =", PB, "Byte")
	fmt.Println("1EB =", EB, "Byte")
	fmt.Println("1ZB =", float64(ZB), "Byte")
	fmt.Println("1YB =", float64(YB), "Byte")
}
