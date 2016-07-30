// Copyright © 2016 shoarai

package intset_test

import (
	"testing"

	"github.com/shoarai/GoTraining/ch11/ex2/intset"
)

func toIntSet(ints []int) intset.IntSet {
	var x intset.IntSet
	for _, n := range ints {
		x.Add(n)
	}
	return x
}

func TestIntsetAdd(t *testing.T) {
	tests := []struct {
		inputs []int
		want   string
	}{
		{[]int{}, "{}"},
		{[]int{1, 10, 2}, "{1 2 10}"},
		{[]int{12, 2, 12}, "{2 12}"},
	}

	for _, test := range tests {
		x := toIntSet(test.inputs)
		if got := x.String(); got != test.want {
			t.Errorf("Inset.String() = %s, want %s", got, test.want)
		}
	}
}

func TestInsetUnionWith(t *testing.T) {
	tests := []struct {
		inputXs []int
		inputYs []int
		want    string
	}{
		{[]int{1, 10, 2}, []int{12}, "{1 2 10 12}"},
		{[]int{50, 144, 2}, []int{50, 21, 50}, "{2 21 50 144}"},
	}

	for _, test := range tests {
		x := toIntSet(test.inputXs)
		y := toIntSet(test.inputYs)
		x.UnionWith(&y)
		if got := x.String(); got != test.want {
			t.Errorf("Inset.String() = %s, want %s", got, test.want)
		}
	}
}

func TestInsetHas(t *testing.T) {
	tests := []struct {
		inputs []int
		wants  map[int]bool
	}{
		{[]int{1, 10, 2}, map[int]bool{1: true, 3: false}},
	}

	for _, test := range tests {
		x := toIntSet(test.inputs)
		for n, want := range test.wants {
			if got := x.Has(n); got != want {
				t.Errorf("Inset.Has(%s) = %q", got, want)
			}
		}
	}
}
