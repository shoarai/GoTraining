// Copyright © 2016 shoarai

// Package links provides a link-extraction function.
package links

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/net/html"
)

// Extract makes an HTTP GET request to the specified URL, parses
// the response as HTML, and returns the links in the HTML document.
func Extract(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("reading responce body of %s : %v", url, err)
	}

	err = writeFile(resp.Request, body)
	if err != nil {
		return nil, fmt.Errorf("writing %s file: %v", url, err)
	}

	doc, err := html.Parse(bytes.NewReader(body))
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}

	var links []string
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key != "href" {
					continue
				}
				link, err := resp.Request.URL.Parse(a.Val)
				if err != nil {
					continue // ignore bad URLs
				}
				links = append(links, link.String())
			}
		}
	}
	forEachNode(doc, visitNode, nil)
	return links, nil
}

var host string

func writeFile(req *http.Request, data []byte) error {
	if host == "" {
		host = req.URL.Host
	} else if host != req.URL.Host {
		return nil
	}

	filename := "out/" + req.URL.Host + req.URL.Path
	if req.URL.Path == "" {
		filename += "/"
	}
	if strings.HasSuffix(filename, "/") {
		filename += "index.html"
	}

	dir := filepath.Dir(filename)
	_, err := os.Stat(dir)
	if err != nil {
		os.MkdirAll(dir, 0777)
	}

	err = ioutil.WriteFile(filename, data, 0666)
	if err != nil {
		return err
	}

	return nil
}

// Copied from gopl.io/ch5/outline2.
func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}
