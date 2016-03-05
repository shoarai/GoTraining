// Copyright © 2016 shoarai
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package anagram_test

import (
	"testing"

	"github.com/shoarai/GoTraining/ch3/ex12/anagram"
)

type testData struct {
	input    []string
	expected bool
}

// -- Test --

func TestAnagram(t *testing.T) {
	datum := [...]testData{
		{input: []string{"", ""}, expected: true},
		{input: []string{"text", "text"}, expected: true},
		{input: []string{"anagram", "graaman"}, expected: true},
		{input: []string{"8789534", "9878435"}, expected: true},
		{input: []string{"test", "tess"}, expected: false},
		{input: []string{"tex", "te"}, expected: false},
	}
	for _, d := range datum {
		test(t, d.input[0], d.input[1], d.expected)
	}
}

func test(t *testing.T, str1, str2 string, exp bool) {
	b := anagram.Anagram(str1, str2)
	if b != exp {
		t.Errorf("Anagram(%q, %q) = %t, want %t", str1, str2, b, exp)
	}
}
