// Copyright © 2016 shoarai

// Expand replaces text that have prefix of $.
package main

import "testing"

func TestExpand(t *testing.T) {
	f := func(s string) string {
		return "pre" + s
	}

	ts := map[string]string{
		"$test":         "pretest",
		"aaa bb$b $ccc": "aaa bbpreb preccc",
		"$a $ c":        "prea pre c",
	}

	for k, v := range ts {
		ex := expand(k, f)
		if ex != v {
			t.Errorf("expand(%s, function) = %s, want %s", k, ex, v)
		}
	}
}
