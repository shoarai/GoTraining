// Copyright © 2016 shoarai

// Crawl3 crawls web links starting with the command-line arguments.
//
// This version uses bounded parallelism.
// For simplicity, it does not address the termination problem.
//
package main

import (
	"flag"
	"fmt"
	"log"

	"gopl.io/ch5/links"
)

type link struct {
	url   string
	depth int
}

func crawl(l link) []link {
	fmt.Printf("depth: %d, url: %s\n", l.depth, l.url)
	list, err := links.Extract(l.url)
	if err != nil {
		log.Print(err)
	}
	var ls []link
	d := l.depth + 1
	for i := range list {
		l := link{list[i], d}
		ls = append(ls, l)
	}
	return ls
}

func main() {
	var depth int
	flag.IntVar(&depth, "depth", 0, "depth of links")
	flag.Parse()

	worklist := make(chan []link)  // lists of URLs, may have duplicates
	unseenLinks := make(chan link) // de-duplicated URLs

	// Add command-line arguments to worklist.
	go func() {
		var ls []link
		for _, v := range flag.Args() {
			ls = append(ls, link{v, 0})
		}
		worklist <- ls
	}()

	// Create 20 crawler goroutines to fetch each unseen link.
	for i := 0; i < 20; i++ {
		go func() {
			for link := range unseenLinks {
				foundLinks := crawl(link)
				go func() { worklist <- foundLinks }()
			}
		}()
	}

	// The main goroutine de-duplicates worklist items
	// and sends the unseen ones to the crawlers.
	seen := make(map[string]bool)
	for list := range worklist {
		for _, link := range list {
			if !seen[link.url] {
				seen[link.url] = true
				if link.depth <= depth {
					unseenLinks <- link
				}
			}
		}
	}
}
