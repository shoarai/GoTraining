package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Input command as argument\n")
		return
	}

	// Save and restore current GOMAXPROCS
	nowPROCS := runtime.GOMAXPROCS(0)
	defer runtime.GOMAXPROCS(nowPROCS)

	strs := strings.Split(os.Args[1], " ")
	cmd := strs[0]
	args := strs[1:]

	numCPU := runtime.NumCPU()
	fmt.Printf("CPU Number: %d\n", runtime.NumCPU())
	for i := 1; i <= numCPU; i++ {
		fmt.Printf("\n")
		runtime.GOMAXPROCS(i)

		start := time.Now()
		err := exec.Command(cmd, args...).Run()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("GOMAXPROCS: %d\n", runtime.GOMAXPROCS(0))
		fmt.Printf("%s elapsed\n", time.Since(start))
	}
}
