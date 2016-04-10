// Copyright © 2016 shoarai

// ElementsByTagName prints the tag name of an HTML document tree.
package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	for _, url := range os.Args[1:] {
		doc, err := node(url)
		if err != nil {
			fmt.Println(err)
			continue
		}

		images := ElementsByTagName(doc, "img")
		headings := ElementsByTagName(doc, "h1", "h2", "h3", "h4")
		for _, v := range images {
			printNode(v)
		}
		for _, v := range headings {
			printNode(v)
		}
	}
}

func ElementsByTagName(n *html.Node, name ...string) []*html.Node {
	var stack []*html.Node
	if n.Type == html.ElementNode {
		for _, s := range name {
			if s == n.Data {
				stack = append(stack, n)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		s := ElementsByTagName(c, name...)
		stack = append(stack, s...)
	}
	return stack
}

func printNode(n *html.Node) {
	if n.Type == html.ElementNode {
		fmt.Printf("<%s", n.Data)
		for _, a := range n.Attr {
			fmt.Printf(" %s", a.Key)
			if a.Val != "" {
				fmt.Printf("=%s", a.Val)
			}
		}
		fmt.Printf("></%s>\n", n.Data)
	}
}

func node(url string) (*html.Node, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, err
	}

	return doc, nil
}
