// Copyright © 2016 shoarai

// Package memotest provides common functions for
// testing various designs of the memo package.
package memotest

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"testing"
	"time"
)

//!+httpRequestBody
func httpGetBody(url string) (interface{}, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

//!-httpRequestBody

var HTTPGetBody = httpGetBody

func incomingURLs() <-chan string {
	ch := make(chan string)
	go func() {
		for _, url := range []string{
			"https://golang.org",
			"https://godoc.org",
			"https://play.golang.org",
			"http://gopl.io",
			"https://golang.org",
			"https://godoc.org",
			"https://play.golang.org",
			"http://gopl.io",
		} {
			ch <- url
		}
		close(ch)
	}()
	return ch
}

type M interface {
	Get(key string, cancel <-chan struct{}) (interface{}, error)
}

// m := memo.New(httpGetBody)
func Sequential(t *testing.T, m M) {
	for url := range incomingURLs() {
		start := time.Now()
		cancel := make(chan struct{})
		value, err := m.Get(url, cancel)
		if err != nil {
			log.Print(err)
			continue
		}
		fmt.Printf("%s, %s, %d bytes\n",
			url, time.Since(start), len(value.([]byte)))
	}
}

// m := memo.New(httpGetBody)
func Concurrent(t *testing.T, m M) {
	var n sync.WaitGroup
	for url := range incomingURLs() {
		n.Add(1)
		go func(url string) {
			defer n.Done()
			start := time.Now()
			cancel := make(chan struct{})
			value, err := m.Get(url, cancel)
			if err != nil {
				log.Print(err)
				return
			}
			fmt.Printf("%s, %s, %d bytes\n",
				url, time.Since(start), len(value.([]byte)))
		}(url)
	}
	n.Wait()
}

// m := memo.New(httpGetBody)
func Cancel(t *testing.T, m M) {
	url := "https://golang.org"

	cancel := make(chan struct{})
	go func() {
		cancel <- struct{}{}
	}()

	value, err := m.Get(url, cancel)
	if err == nil {
		t.Errorf("Get(%q) = , nil", value)
	}
}
