// Copyright © 2016 shoarai

package intset

import "testing"

func initSetInit(n ...int) *IntSet {
	var x IntSet
	for _, v := range n {
		x.Add(v)
	}
	return &x
}

func TestElems(t *testing.T) {
	x := initSetInit(1, 144, 9)
	s := x.Elems()

	for i, act := range s {
		var exp bool
		if i == 1 || i == 144 || i == 9 {
			exp = true
			if act != exp {
				t.Errorf("s := x.Elems(); s[%d]= %v, want %v", i, act, exp)
			}
		} else {
			exp = false
			if act != exp {
				t.Errorf("s := x.Elems(); s[%d]= %v, want %v", i, act, exp)
			}
		}
	}
}
