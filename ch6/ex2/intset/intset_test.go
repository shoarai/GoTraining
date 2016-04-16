// Copyright © 2016 shoarai

package intset

import "testing"

func TestAddAll(t *testing.T) {
	var x IntSet
	exp := "{1 9 144}"
	x.AddAll(1, 144, 9)

	if act := x.String(); act != exp {
		t.Errorf("x.String() = %s, want %s", act, exp)
	}
}
