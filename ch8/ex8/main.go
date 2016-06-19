// Copyright © 2016 shoarai

// Reverb2 is a TCP server that simulates an echo.
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

func handleConn(c net.Conn) {
	input := bufio.NewScanner(c)
	done := make(chan struct{})

	go func() {
		for input.Scan() {
			done <- struct{}{}
			go echo(c, input.Text(), 1*time.Second)
		}
	}()

	const timeoutSec = 10
	ticker := time.NewTicker(1 * time.Second)
	for countdown := timeoutSec; countdown > 0; countdown-- {
		select {
		case <-done:
			countdown = timeoutSec
		default:
		}
		<-ticker.C
	}
	ticker.Stop()

	// NOTE: ignoring potential errors from input.Err()
	c.Close()
	close(done)
}

func main() {
	l, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn)
	}
}
