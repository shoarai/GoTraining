// Copyright © 2016 shoarai

// Bytecounter demonstrates an implementation of io.Writer that counts bytes.
package main

import (
	"bufio"
	"bytes"
	"fmt"
)

type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p)) // convert int to ByteCounter
	return len(p), nil
}

type WordCounter int

func (c *WordCounter) Write(p []byte) (int, error) {
	sc := bufio.NewScanner(bytes.NewReader(p))
	sc.Split(bufio.ScanWords)
	count := 0
	for sc.Scan() {
		count++
	}
	*c += WordCounter(count)
	return count, nil
}

type LineCounter int

func (c *LineCounter) Write(p []byte) (int, error) {
	sc := bufio.NewScanner(bytes.NewReader(p))
	count := 0
	for sc.Scan() {
		count++
	}
	*c += LineCounter(count)
	return count, nil
}

func main() {
	var bc ByteCounter
	var wc WordCounter
	var lc LineCounter
	bs := []byte("hello")
	bc.Write(bs)
	wc.Write(bs)
	lc.Write(bs)
	fmt.Println(bc) // "5", = len("hello")
	fmt.Println(wc) // "1", = Words of "hello"
	fmt.Println(lc) // "1", = Lines of "hello"

	// reset the counter
	bc = 0
	wc = 0
	lc = 0

	var name = "Dolly"
	fmt.Fprintf(&bc, "hello, %s", name)
	fmt.Fprintf(&wc, "hello, %s", name)
	fmt.Fprintf(&lc, "hello, %s", name)
	fmt.Println(bc) // "12", = len("hello, Dolly")
	fmt.Println(wc) // "2", = Words of "hello, Dolly"
	fmt.Println(lc) // "1", = Lines of "hello, Dolly"

	lc = 0
	fmt.Fprintf(&lc, "hello,\n%s\nBye", name)
	fmt.Println(lc) // "3", = Lines of "hello,\nDolly\nBye"
}
