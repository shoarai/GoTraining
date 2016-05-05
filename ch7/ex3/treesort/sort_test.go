// Copyright © 2016 shoarai

package treesort

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"

	"../treesort"
)

func TestSort(t *testing.T) {
	data := make([]int, 50)
	for i := range data {
		data[i] = rand.Int() % 50
	}
	treesort.Sort(data)
	if !sort.IntsAreSorted(data) {
		t.Errorf("not sorted: %v", data)
	}
}

func TestTreeString(t *testing.T) {
	var tr *tree
	tr = add(tr, 5)
	tr = add(tr, 10)
	tr = add(tr, 2)
	tr = add(tr, 1)
	tr = add(tr, 3)
	expect := "1 2 3 5 10"

	actual := fmt.Sprintf("%s", tr)
	if actual != expect {
		t.Errorf("tree.string() = %s, want %s", actual, expect)
	}
}
