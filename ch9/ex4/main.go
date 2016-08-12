// Copyright © 2016 shoarai

package main

import (
	"fmt"
	"log"
	"os"
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

	fmt.Printf("Number of pipline: %d\n", num)

	// Create pipeline
	start := time.Now()
	in, out := createPipeline(num)
	fmt.Printf("Create: %s elapsed\n", time.Since(start))

	// Start pipeline
	start = time.Now()
	in <- struct{}{}
	<-out
	fmt.Printf("Transmission: %s elapsed\n", time.Since(start))
}

func createPipeline(num int) (chan<- struct{}, <-chan struct{}) {
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

	return firstIn, *lastOut
}

func pipeline(in <-chan struct{}, out chan<- struct{}) {
	go func() {
		out <- <-in
	}()
}
