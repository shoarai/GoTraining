// Copyright © 2016 shoarai

package strings_test

import (
	"strings"
	"testing"
)

func TestSplit(t *testing.T) {
	tests := []struct {
		str, separator string
		want           int
	}{
		{"a:b:c", ":", 3},
		{"ab. cde:.fg.h", ".", 4},
		{"", ":", 1},
		{"&ab&", "&", 3},
	}

	for _, test := range tests {
		words := strings.Split(test.str, test.separator)
		if got := len(words); got != test.want {
			t.Errorf("Split(%q, %q) returned %d words, want %d",
				test.str, test.separator, got, test.want)
		}
	}
}
