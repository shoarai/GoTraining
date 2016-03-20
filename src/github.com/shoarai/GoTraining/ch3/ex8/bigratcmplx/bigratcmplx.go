// Copyright © 2016 shoarai
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package bigratcmplx

import (
	"math"
	"math/big"
)

type Complex struct {
	R *big.Rat
	I *big.Rat
}

func New(x, y float64) *Complex {
	var c Complex
	c.R = new(big.Rat).SetFloat64(x)
	c.I = new(big.Rat).SetFloat64(y)
	return &c
}

func Add(x, y *Complex) *Complex {
	c := New(0, 0)
	c.R.Add(x.R, y.R)
	c.I.Add(x.I, y.I)
	return c
}

func Multi(x, y *Complex) *Complex {
	c := New(0, 0)
	c.R.Sub(new(big.Rat).Mul(x.R, y.R), new(big.Rat).Mul(x.I, y.I))
	c.I.Add(new(big.Rat).Mul(x.R, y.I), new(big.Rat).Mul(x.I, y.R))
	return c
}

func Abs(v *Complex) float64 {
	vr, _ := v.R.Float64()
	vi, _ := v.I.Float64()
	return math.Hypot(vr, vi)
}
