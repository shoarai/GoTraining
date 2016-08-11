// Copyright © 2016 shoarai

// The jpeg command reads a PNG image from the standard input
// and writes it as a JPEG image to the standard output.
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
)

type Package struct {
	Dir           string // directory containing package sources
	ImportPath    string // import path of package in dir
	ImportComment string // path in import comment on package statement
	Name          string // package name
	Doc           string // package documentation string
	Target        string // install path
	Shlib         string // the shared library that contains this package (only set when -linkshared)
	Goroot        bool   // is this package in the Go root?
	Standard      bool   // is this package part of the standard Go library?
	Stale         bool   // would 'go install' do anything for this package?
	Root          string // Go root or Go path dir containing this package
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Input package as arugument")
		return
	}
	packName := os.Args[1]

	pack, err := packageList(packName)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v", pack)
}

func packageList(packName string) (*Package, error) {
	args := []string{"list", "-json", packName}
	outputs, err := exec.Command("go", args...).Output()
	if err != nil {
		return nil, err
	}

	var pack Package
	reader := bytes.NewReader(outputs)
	if err := json.NewDecoder(reader).Decode(&pack); err != nil {
		return nil, err
	}

	return &pack, nil
}
