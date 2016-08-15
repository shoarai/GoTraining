// Copyright © 2016 shoarai

package main

import (
	"fmt"
	"log"
	"os"

	"github.com/shoarai/GoTraining/ch10/ex2/archive"
	_ "github.com/shoarai/GoTraining/ch10/ex2/archive/tar"
	_ "github.com/shoarai/GoTraining/ch10/ex2/archive/zip"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Input file path as argument")
		return
	}
	path := os.Args[1]

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	err = archive.Read(file, os.Stdout)
	if err != nil {
		log.Fatal(err)
	}
}
