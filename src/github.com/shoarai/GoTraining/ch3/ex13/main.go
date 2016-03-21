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
	fmt.Printf("1KB = %d Byte\n", KB)
	fmt.Printf("1MB = %d Byte\n", MB)
	fmt.Printf("1GB = %d Byte\n", GB)
	fmt.Printf("1TB = %d Byte\n", TB)
	fmt.Printf("1PB = %d Byte\n", PB)
	fmt.Printf("1EB = %d Byte\n", EB)
	fmt.Printf("1ZB = %f Byte\n", float64(ZB))
	fmt.Printf("1YB = %f Byte\n", float64(YB))
}
