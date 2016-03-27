// Copyright © 2016 shoarai

// CountWordsAndImages does an HTTP GET on each URL, parses the
// result as HTML, and counts the number of words and images within it.
//
// Usage:
//	countWordsAndImages url ...
package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

// visit appends to links each link found in n, and returns the result.
func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}

func main() {
	for _, url := range os.Args[1:] {
		words, images, err := CountWordsAndImages(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "countWordsAndImages: %v\n", err)
			continue
		}
		fmt.Printf("%s\nwords: %d\nimages: %d\n", url, words, images)
	}
}

// CountWordsAndImages performs an HTTP GET request for url, parses the
// response as HTML, and extracts and returns the number of words and images.
func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		err = fmt.Errorf("getting %s: %s", url, resp.Status)
		return
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("parsing %s as HTML: %v", url, err)
		return
	}
	words, images = countWordsAndImages(doc)
	return
}

func countWordsAndImages(n *html.Node) (words, images int) {
	if n.Type == html.TextNode {
		r := strings.NewReader(n.Data)
		in := bufio.NewScanner(r)
		in.Split(bufio.ScanWords)
		for in.Scan() {
			words++
		}
	}
	if n.Type == html.ElementNode && n.Data == "img" {
		images++
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		w, i := countWordsAndImages(c)
		words += w
		images += i
	}
	return
}
