// Copyright © 2016 shoarai

// Netcat1 is a read-only TCP client.
package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Input ports as argument")
		return
	}
	url := "localhost:" + os.Args[1]

	go func(url string) {
		conn, err := net.Dial("tcp", url)
		if err != nil {
			log.Fatal(err)
		}
		defer conn.Close()
		mustCopy(os.Stdout, conn)
	}(url)

	url = "localhost:" + os.Args[2]

	conn, err := net.Dial("tcp", url)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	mustCopy(os.Stdout, conn)
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
