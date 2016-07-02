package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Printf("CPU Number: %d", runtime.NumCPU())
	// fmt.Println(runtime.GOMAXPROCS(0))
	// fmt.Println(runtime.NumGoroutine())
	// for {
	// 	go fmt.Print(0)
	// 	fmt.Print(1)
	// }
}
