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

func TestUnionWith(t *testing.T) {
	x := initSetInit(1, 144, 9)
	y := initSetInit(9, 42)
	exp := "{1 9 42 144}"

	x.UnionWith(y)
	if act := x.String(); act != exp {
		t.Errorf("x.String() = %s, want %s", act, exp)
	}
}

func TestInterSectWith(t *testing.T) {
	x := initSetInit(1, 144, 9)
	y := initSetInit(9, 42)
	exp := "{9}"

	x.IntersectWith(y)
	if act := x.String(); act != exp {
		t.Errorf("x.String() = %s, want %s", act, exp)
	}
}

func TestDeffereceWith(t *testing.T) {
	x := initSetInit(1, 144, 9)
	y := initSetInit(9, 42)
	exp := "{1 144}"

	x.DefferenceWith(y)
	if act := x.String(); act != exp {
		t.Errorf("x.String() = %s, want %s", act, exp)
	}
}

func TestSystematicWith(t *testing.T) {
	x := initSetInit(1, 144, 9)
	y := initSetInit(9, 42)
	exp := "{1 42 144}"

	x.SymmetricWith(y)
	if act := x.String(); act != exp {
		t.Errorf("x.String() = %s, want %s", act, exp)
	}
}
