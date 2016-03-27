// Copyright © 2016 shoarai

// Findlinks1 prints the links in an HTML document read from standard input.
package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	textNode(doc, false)
}

func textNode(n *html.Node, ignore bool) {
	if n.Type == html.ElementNode {
		if n.Data == "script" {
			ignore = true
		} else {
			ignore = false
		}
	} else if n.Type == html.TextNode && !ignore {
		fmt.Printf("%s\n", n.Data)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		textNode(c, ignore)
	}
}
