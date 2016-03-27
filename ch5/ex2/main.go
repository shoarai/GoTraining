// Copyright © 2016 shoarai

// FindTags counts the tag in an HTML document read from standard input.
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
	tags := make(map[string]int, 0)
	countTag(tags, doc)
	for k, v := range tags {
		fmt.Printf("%s: %d\n", k, v)
	}
}

func countTag(tags map[string]int, n *html.Node) {
	if n.Type == html.ElementNode {
		tags[n.Data]++
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		countTag(tags, c)
	}
}
