// Copyright © 2016 shoarai

// LimitReader reads bytes limited with its length
package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

type Reader interface {
	Read(p []byte) (n int, err error)
}

type LimitedReader struct {
	R Reader
	N int64
}

func (l *LimitedReader) Read(p []byte) (n int, err error) {
	if l.N <= 0 {
		return 0, io.EOF
	}
	if int64(len(p)) > l.N {
		p = p[0:l.N]
	}
	n, err = l.R.Read(p)
	l.N -= int64(n)
	return
}

func LimitReader(r Reader, n int64) Reader {
	return &LimitedReader{r, n}
}

func main() {
	str := "hello"
	var n int64 = 3

	r := strings.NewReader(str)
	scanner := bufio.NewScanner(LimitReader(r, n))
	var limStr string
	for scanner.Scan() {
		limStr += scanner.Text()
	}
	fmt.Printf("%q limited with length of %d is %q\n", str, n, limStr)
}
