// Copyright © 2016 shoarai

// CountingWriter counts an implementation of io.Writer that counts bytes.
package main

import (
	"fmt"
	"io"
)

type countingWriter struct {
	writer io.Writer
	count  *int64
}

func (c *countingWriter) Write(p []byte) (int, error) {
	n, err := c.writer.Write(p)
	// if err == nil {
	// 	*c.count += int64(n)
	// }
	return n, err
}

// func CountingWriter(w io.Writer) (io.Writer, *int64) {
// 	var cntr int64
// 	retval := countingWriter{&cntr, w}
// 	return retval, retval.count
// }

func main() {
	// var b bytes.Buffer
	var w io.Writer
	var count *int64
	// w, cntr = CountingWriter(&b)

	cw := countingWriter{
		w, count,
	}
	cw.Write([]byte("hello"))
	// fmt.Println(cw)

	fmt.Fprintf(&cw, "hello, %s", "")

	// buf := &bytes.Buffer{}
	// buf.WriteString("s string")
	// _, count := CountingWriter(buf)
	// fmt.Println("a ...interface{}")
	// os.Stdout.Write("b []byte")

	// fmt.Println("a ...interface{}")
	// _, count := CountingWriter(os.Stdout)
	// fmt.Println("a ...interface{}")
	// os.Stdout.Write("b []byte")
	// fmt.Println(*count)

	// var w io.Writer
	// bs := []byte("hello")
	// w.Write(bs)

	// w, count := CountingWriter(w)
	// fmt.Println(*count)
}
