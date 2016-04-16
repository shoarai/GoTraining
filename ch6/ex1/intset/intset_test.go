// Copyright © 2016 shoarai

package intset

import (
	"fmt"
	"testing"
)

func initSetInit(n ...int) *IntSet {
	var x IntSet
	for _, v := range n {
		x.Add(v)
	}
	return &x
}

func TestLen(t *testing.T) {
	x := initSetInit(1, 144, 9)
	exp := 3

	if act := x.Len(); act != exp {
		t.Errorf("x.Len() = %d, want %d", act, exp)
	}
}

func TestRemove(t *testing.T) {
	x := initSetInit(1, 144, 9)
	in := 144
	exp := "{1 9}"

	x.Remove(in)
	if act := x.String(); act != exp {
		t.Errorf("x.String() = %s, want %s", act, exp)
	}
}

func TestClear(t *testing.T) {
	x := initSetInit(1, 144, 9)
	exp := "{}"

	x.Clear()
	if act := x.String(); act != exp {
		t.Errorf("x.String() = %s, want %s", act, exp)
	}
}

func TestCopy(t *testing.T) {
	x := initSetInit(1, 144, 9)
	exp := x.String()

	c := x.Copy()
	x.Add(10) // dummy data to verify object is copied
	if act := c.String(); act != exp {
		t.Errorf("x.String() = %s, want %s", act, exp)
	}
}

func TestAdd(t *testing.T) {
	var x IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	expected := "{1 9 144}"
	if x.String() != expected {
		t.Errorf("x.String() = %s, want %s", x.String(), expected)
	}
}

func Example_one() {
	//!+main
	var x, y IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	fmt.Println(x.String()) // "{1 9 144}"

	y.Add(9)
	y.Add(42)
	fmt.Println(y.String()) // "{9 42}"

	x.UnionWith(&y)
	fmt.Println(x.String()) // "{1 9 42 144}"

	fmt.Println(x.Has(9), x.Has(123)) // "true false"
	//!-main

	// Output:
	// {1 9 144}
	// {9 42}
	// {1 9 42 144}
	// true false
}

func Example_two() {
	var x IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	x.Add(42)

	//!+note
	fmt.Println(&x)         // "{1 9 42 144}"
	fmt.Println(x.String()) // "{1 9 42 144}"
	fmt.Println(x)          // "{[4398046511618 0 65536]}"
	//!-note

	// Output:
	// {1 9 42 144}
	// {1 9 42 144}
	// {[4398046511618 0 65536]}
}
