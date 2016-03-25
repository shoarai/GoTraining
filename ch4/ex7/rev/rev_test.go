// Copyright © 2016 shoarai
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package rev_test

import (
	"testing"

	"github.com/shoarai/GoTraining/ch4/ex7/rev"
)

type testData struct {
	input    []byte
	expected []byte
}

// -- Test --

func TestReverse(t *testing.T) {
	datum := [...]testData{{
		input:    []byte{'0', '1', '2', '3', '4'},
		expected: []byte{'4', '3', '2', '1', '0'},
	}, {
		[]byte{'a', 'b', 'c', 'd'},
		[]byte{'d', 'c', 'b', 'a'},
	}}

	for _, v := range datum {
		test(t, v.input, v.expected)
	}
}

func test(t *testing.T, b, e []byte) {
	var ans []byte
	ans = append(ans, b...)

	rev.Reverse(ans)

	if len(ans) != len(e) {
		t.Errorf("Reverse(%s) = %s, want %s", b, ans, e)
		return
	}
	for i := range ans {
		if ans[i] != e[i] {
			t.Errorf("Reverse(%s) = %s, want %s", b, ans, e)
			break
		}
	}
}
