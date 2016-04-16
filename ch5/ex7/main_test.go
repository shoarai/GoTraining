// Copyright © 2016 shoarai

package main

import "testing"

func TestOutline(t *testing.T) {
	err := outline("https://golang.org/")
	if err != nil {
		t.Errorf("failed: %v", err)
	}
}
