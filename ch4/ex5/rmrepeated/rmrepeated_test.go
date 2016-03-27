// Copyright © 2016 shoarai
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package rmrepeated_test

import (
	"testing"

	"../rmrepeated"
)

type testData struct {
	input    []string
	expected []string
}

// -- Test --

func TestRemoveRepeated(t *testing.T) {
	d := [...]testData{{
		input:    []string{"", ""},
		expected: []string{""},
	}, {
		[]string{"text", "text", "text"},
		[]string{"text"},
	}, {
		[]string{"01234", "12345", "01234", "01234"},
		[]string{"01234", "12345", "01234"},
	}}

	for _, v := range d {
		test(t, v.input, v.expected)
	}
}

func test(t *testing.T, s, e []string) {
	c := rmrepeated.RemoveRepeated(s)

	if len(c) != len(e) {
		t.Errorf("RemoveRepeated(%s) = %s, want %s", s, c, e)
		return
	}
	for i := range c {
		if c[i] != e[i] {
			t.Errorf("RemoveRepeated(%s) = %s, want %s", s, c, e)
			break
		}
	}
}
