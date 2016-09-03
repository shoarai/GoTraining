// Copyright © 2016 shoarai

package cycle

import "testing"

func TestIsCycle(t *testing.T) {
	type link struct {
		value string
		tail  *link
	}

	a, b := link{value: "a"}, link{value: "b"}
	d := link{value: "d", tail: nil}
	a.tail, b.tail = &b, &a

	for _, test := range []struct {
		x    interface{}
		want bool
	}{
		{a, true},
		{b, true},
		{d, false},
	} {
		if IsCycle(test.x) != test.want {
			t.Errorf("IsCycle(%v) = %t", test.x, !test.want)
		}
	}
}
