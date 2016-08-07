package main

import (
	"fmt"
	"time"
)

func main() {
	client := make(chan struct{})
	server := make(chan struct{})
	done := make(chan struct{})
	count := make(chan int64)

	go func() {
		countIO(client, server, done, count)
	}()

	go func() {
		for {
			<-server
			client <- struct{}{}
		}
	}()

	client <- struct{}{}
	time.Sleep(time.Second)
	close(done)

	c := <-count
	fmt.Printf("Number of cycles: %d\n", c)
}

func countIO(
	in <-chan struct{}, out chan<- struct{},
	done <-chan struct{}, count chan<- int64) {

	var c int64
	for {
		select {
		case <-in:
			c++
			out <- struct{}{}
		case <-done:
			count <- c
			break
		}
	}
}
