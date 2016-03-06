// Copyright © 2016 shoarai
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Rev reverses a slice.
package rmrepeated

import (
	"testing"

	"github.com/shoarai/GoTraining/ch4/ex5/rmrepeated"
)

type testData struct {
	input    []string
	expected []string
}

// -- Test --

func TestRemoveRepeated(t *testing.T) {
	datum := [...]testData{
		{input: []string{"", ""}, expected: []string{""}},
		{[]string{"text", "text", "text"}, []string{"text"}},
		{[]string{"01234", "12345", "01234", "01234"}, []string{"01234", "12345", "01234"}},
	}

	for _, v := range datum {
		c := rmrepeated.RemoveRepeated(v.input)
		if len(c) != len(v.expected) {
			t.Errorf("RemoveRepeated(%s) = %s, want %s", v.input, c, v.expected)
		}
		for i := range c {
			if i < len(v.expected) && c[i] != v.expected[i] {
				t.Errorf("RemoveRepeated(%s) = %s, want %s", v.input, c, v.expected)
				break
			}
		}
	}
}
