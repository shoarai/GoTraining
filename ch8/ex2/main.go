// Copyright © 2016 shoarai

// Clock is a TCP server that periodically writes the time.
package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/shoarai/GoTraining/ch8/ex2/ftp"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Input port as argument")
		return
	}
	url := "localhost:" + os.Args[1]

	listener, err := net.Listen("tcp", url)
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go ftp.HandleConnection(conn) // handle connections concurrently
	}
}
