// Copyright © 2016 shoarai
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Rev reverses a slice.
package rmrepeated

import (
	"reflect"
	"testing"

	"github.com/shoarai/GoTraining/ch4/ex6/rmrepeated"
)

type testData struct {
	input    []byte
	expected []byte
}

// -- Test --

func TestConvSpaces(t *testing.T) {
	datum := [...]testData{
		{input: []byte("t　　est"), expected: []byte("t est")},
		{input: []byte("あ　　いう　え　お"), expected: []byte("あ いう　え　お")},
	}

	for _, v := range datum {
		c := rmrepeated.ConvSpaces(v.input)
		if !reflect.DeepEqual(v.expected[:], c) {
			t.Errorf("ConvSpaces(%s) = %s, want %s", v.input, c, v.expected)
		}
	}
}
