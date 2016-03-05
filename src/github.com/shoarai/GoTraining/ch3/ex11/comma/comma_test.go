// Copyright © 2016 shoarai
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package comma_test

import (
	"testing"

	"github.com/shoarai/GoTraining/ch3/ex11/comma"
)

type testData struct {
	input    string
	expected string
}

// -- Test --

func TestComma(t *testing.T) {
	datum := [...]testData{
		{input: "1", expected: "1"},
		{"1234", "1,234"},
		{"-1234", "-1,234"},
		{"1234.5", "1,234.5"},
		{"12345.67890", "12,345.67890"},
	}

	for _, d := range datum {
		test(t, d.input, d.expected)
	}
}

func test(t *testing.T, str, exp string) {
	b := comma.Comma(str)
	if b != exp {
		t.Errorf("Comma(%q) = %s, want %s", str, b, exp)
	}
}
