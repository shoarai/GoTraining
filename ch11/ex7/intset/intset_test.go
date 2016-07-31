// Copyright © 2016 shoarai

package intset_test

import (
	"math"
	"math/rand"
	"testing"
	"time"

	"github.com/shoarai/GoTraining/ch11/ex7/intset"
)

func newRand() *rand.Rand {
	seed := time.Now().UnixNano()
	return rand.New(rand.NewSource(seed))
}

func BenchmarkAdd(b *testing.B) {
	var x intset.IntSet
	rng := newRand()
	for i := 0; i < b.N; i++ {
		x.Add(rng.Intn(math.MaxInt8))
	}
}

func BenchmarkUnionWith(b *testing.B) {
	var x, y intset.IntSet
	rng := newRand()
	for i := 0; i < math.MaxInt8; i++ {
		x.Add(rng.Intn(math.MaxInt8))
		y.Add(rng.Intn(math.MaxInt8))
	}
	for i := 0; i < b.N; i++ {
		x.UnionWith(&y)
	}
}

func BenchmarkHas(b *testing.B) {
	var x intset.IntSet
	rng := newRand()
	for i := 0; i < math.MaxInt8; i++ {
		x.Add(rng.Intn(math.MaxInt8))
	}
	for i := 0; i < b.N; i++ {
		x.Has(rng.Intn(math.MaxInt8))
	}
}
