// Copyright © 2016 shoarai

// Bytecounter demonstrates an implementation of io.Writer that counts bytes.
package main

import "bufio"

type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p)) // convert int to ByteCounter
	return len(p), nil
}

type WordCounter int

func (c *WordCounter) Write(p []byte) (int, error) {
	// input := bufio.NewScanner(f)
	// input.Split(bufio.ScanWords)
	//
	// for input.Scan() {
	// 	text := strings.ToLower(input.Text())
	// 	counts[text]++
	// }

	words, _, err := bufio.ScanWords(p, true)
	if err != nil {
		return 0, err
	}
	*c += WordCounter(words) // convert int to WordCounter
	return words, nil
}

type LineCounter int

func (c *LineCounter) Write(p []byte) (int, error) {
	// words := bufio.ScanWords(p)
	// *c += LineCounter(words) // convert int to LineCounter
	return len(p), nil
}
