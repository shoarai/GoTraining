// Copyright © 2016 shoarai

package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"strconv"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Input number of pipeline")
		return
	}

	num, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	// Create pipeline
	start := time.Now()
	firstIn := make(chan struct{})
	var in *chan struct{}
	var lastOut *chan struct{}
	for i := 0; i < num; i++ {
		if i == 0 {
			in = &firstIn
		} else {
			in = lastOut
		}

		out := make(chan struct{})
		pipeline(*in, out)
		lastOut = &out
	}

	fmt.Printf("NumGoroutine: %d\n", runtime.NumGoroutine())
	fmt.Printf("Create: %s elapsed\n", time.Since(start))

	// Start pipeline
	start = time.Now()
	firstIn <- struct{}{}
	<-(*lastOut)
	fmt.Printf("Run: %s elapsed\n", time.Since(start))
}

func pipeline(in <-chan struct{}, out chan<- struct{}) {
	go func() {
		out <- <-in
	}()
}
