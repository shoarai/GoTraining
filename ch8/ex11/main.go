// Copyright © 2016 shoarai

// Fetch saves the contents of a URL into a local file.
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
)

// Fetch downloads the URL and returns the
// name and length of the local file.
func fetch(url string, cancel <-chan struct{}) (filename string, n int64, err error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", 0, err
	}
	req.Cancel = cancel

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()

	local := path.Base(resp.Request.URL.Path)
	if local == "/" {
		local = "index.html"
	}
	f, err := os.Create(local)
	if err != nil {
		return "", 0, err
	}
	n, err = io.Copy(f, resp.Body)
	// Close file, but prefer error from Copy, if any.
	if closeErr := f.Close(); err == nil {
		err = closeErr
	}
	return local, n, err
}

type responce struct {
	url       string
	filenName string
	n         int64
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Input URLs as arguments")
		return
	}

	done := make(chan struct{})
	result := make(chan responce)

	for _, url := range os.Args[1:] {
		go func(url string) {
			filenName, n, err := fetch(url, done)
			if err != nil {
				fmt.Fprintf(os.Stderr, "fetch %s: %v\n", url, err)
				return
			}
			result <- responce{url, filenName, n}
		}(url)
	}

	res := <-result
	fmt.Fprintf(os.Stderr, "%s => %s (%d bytes).\n", res.url, res.filenName, res.n)
}
