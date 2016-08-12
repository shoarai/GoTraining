// Copyright © 2016 shoarai

// The jpeg command reads a PNG image from the standard input
// and writes it as a JPEG image to the standard output.
package main

import (
	"flag"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"os"
)

func main() {
	var format = flag.String("format", "jpeg", "output format")
	flag.Parse()

	if err := convertImageFormat(os.Stdin, os.Stdout, *format); err != nil {
		fmt.Fprintf(os.Stderr, "jpeg: %v\n", err)
		os.Exit(1)
	}
}

func convertImageFormat(in io.Reader, out io.Writer, format string) error {
	img, kind, err := image.Decode(in)
	if err != nil {
		return err
	}
	fmt.Fprintln(os.Stderr, "Input format =", kind)
	fmt.Fprintln(os.Stderr, "Output format =", format)

	switch format {
	case "jpeg":
		return jpeg.Encode(out, img, &jpeg.Options{Quality: 95})
	case "png":
		return png.Encode(out, img)
	case "gif":
		return gif.Encode(out, img, &gif.Options{})
	default:
		return fmt.Errorf("Unsupported format")
	}
}
