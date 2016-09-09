// Copyright © 2016 shoarai

// Package params provides a reflection-based parser for URL parameters.
package params

import "testing"

func TestPack(t *testing.T) {
	tests := []struct {
		val  interface{}
		want string
	}{
		{struct {
			str1 string
			str2 string
		}{"golang", "programming"},
			"str1=golang&str2=programming"},
		{struct {
			str1 int `http:"str1"`
		}{1},
			"str1=1"},
		// TODO: Array
		// {struct {
		// 	str []string
		// }{[]string{"golang", "programming"}},
		// 	"str=golang&str=programming"},
	}

	for _, test := range tests {
		got := Pack(test.val)
		if got != test.want {
			t.Errorf("Pack(%v) = %v, want %v", test.val, got, test.want)
		}
	}
}
