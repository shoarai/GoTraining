// Copyright © 2016 shoarai

// Xmlselect prints the text of selected elements of an XML document.
package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
)

type Node interface{} // CharData or Element

type CharData string

type Element struct {
	Type     xml.Name
	Attr     []xml.Attr
	Children []Node
}

func main() {
	dec := xml.NewDecoder(os.Stdin)
	var element Element

	var stack []*Element // stack of element names
	stack = append(stack, &element)

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
			elem := stack[len(stack)-1]
			e := Element{Type: tok.Name, Attr: tok.Attr}
			elem.Children = append(elem.Children, e)
			stack = append(stack, elem) // push

		case xml.EndElement:
			stack = stack[:len(stack)-1] // pop

		case xml.CharData:
			elem := stack[len(stack)-1]
			elem.Children = append(elem.Children, CharData(tok))
		}
	}

	printTree(element)
}

var depth int

func printTree(elem Node) {
	switch elem := elem.(type) {
	case Element:
		fmt.Printf("%*s<%s>\n", depth, "", elem.Type.Local)
		// depth++
		for _, v := range elem.Children {
			printTree(v)
		}
	case CharData:
		fmt.Println(elem)
	}
}
