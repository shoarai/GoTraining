// Copyright © 2016 shoarai

// CountElement counts the tag in an HTML document read from standard input.
package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "CountElement: %v\n", err)
		os.Exit(1)
	}
	elems := make(map[string]int, 0)
	countElement(elems, doc)
	for k, v := range elems {
		fmt.Printf("%s: %d\n", k, v)
	}
}

func countElement(elems map[string]int, n *html.Node) {
	if n.Type == html.ElementNode {
		elems[n.Data]++
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		countElement(elems, c)
	}
}
