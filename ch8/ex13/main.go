// Copyright © 2016 shoarai

// Chat is a server that lets clients chat with each other.
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

type client chan<- string // an outgoing message channel

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string) // all incoming client messages
)

func broadcaster() {
	clients := make(map[client]bool) // all connected clients
	for {
		select {
		case msg := <-messages:
			// Broadcast incoming message to all
			// clients' outgoing message channels.
			for cli := range clients {
				cli <- msg
			}

		case cli := <-entering:
			clients[cli] = true

		case cli := <-leaving:
			delete(clients, cli)
			close(cli)
		}
	}
}

func setTimeout(onTimeout func(), duration time.Duration) chan<- struct{} {
	done := make(chan struct{})

	go func() {
		ticker := time.NewTicker(duration)
		for {
			select {
			case <-done:
				ticker = time.NewTicker(duration)
			case <-ticker.C:
				onTimeout()
			}
		}
	}()

	return done
}

func handleConn(conn net.Conn) {
	ch := make(chan string) // outgoing client messages
	go clientWriter(conn, ch)

	who := conn.RemoteAddr().String()
	ch <- "You are " + who
	messages <- who + " has arrived"
	entering <- ch

	restartTimer := setTimeout(func() {
		conn.Close()
	}, 5*time.Minute)

	input := bufio.NewScanner(conn)
	for input.Scan() {
		restartTimer <- struct{}{}
		messages <- who + ": " + input.Text()
	}

	leaving <- ch
	messages <- who + " has left"
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg) // NOTE: ignoring network errors
	}
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}
