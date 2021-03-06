// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package memo_test

import (
	"testing"

	"github.com/shoarai/GoTraining/ch9/ex3"
	"github.com/shoarai/GoTraining/ch9/ex3/memotest"
)

var httpGetBody = memotest.HTTPGetBody

func Test(t *testing.T) {
	m := memo.New(httpGetBody)
	defer m.Close()
	memotest.Sequential(t, m)
}

func TestConcurrent(t *testing.T) {
	m := memo.New(httpGetBody)
	defer m.Close()
	memotest.Concurrent(t, m)
}

func TestCancel(t *testing.T) {
	m := memo.New(httpGetBody)
	defer m.Close()
	memotest.Cancel(t, m)
}
