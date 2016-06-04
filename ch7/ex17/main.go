// Copyright © 2016 shoarai

// Xmlselect prints the text of selected elements of an XML document.
package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	dec := xml.NewDecoder(os.Stdin)
	var stack []xml.StartElement // stack of element names
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "xmlselect: %v\n", err)
			os.Exit(1)
		}
		switch tok := tok.(type) {
		case xml.StartElement:
			stack = append(stack, tok) // push
		case xml.EndElement:
			stack = stack[:len(stack)-1] // pop
		case xml.CharData:
			if containsAll(stack, os.Args[1:]) {
				var str string
				for _, v := range stack {
					str += v.Name.Local + " "
				}
				fmt.Printf("%s: %s\n", str, tok)
			}
		}
	}
}

// containsAll reports whether x contains the elements of y, in order.
func containsAll(x []xml.StartElement, y []string) bool {
	for len(y) <= len(x) {
		if len(y) == 0 {
			return true
		}

		if strings.HasPrefix(y[0], "#") {
			for _, v := range x[0].Attr {
				if v.Name.Local == "id" && "#"+v.Value == y[0] {
					y = y[1:]
				}
			}
		} else if strings.HasPrefix(y[0], ".") {
			for _, v := range x[0].Attr {
				if v.Name.Local == "class" && "."+v.Value == y[0] {
					y = y[1:]
				}
			}
		} else {
			if x[0].Name.Local == y[0] {
				y = y[1:]
			}
		}

		x = x[1:]
	}
	return false
}
