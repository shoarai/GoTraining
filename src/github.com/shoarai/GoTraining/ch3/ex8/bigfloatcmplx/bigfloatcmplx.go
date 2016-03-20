// Copyright © 2016 shoarai
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package bigfloatcmplx

import (
	"math"
	"math/big"
)

type BigFloatComplex struct {
	R *big.Float
	I *big.Float
}

func New(x, y float64) *BigFloatComplex {
	var c BigFloatComplex
	c.R = big.NewFloat(x)
	c.I = big.NewFloat(y)
	return &c
}

func Add(x, y *BigFloatComplex) *BigFloatComplex {
	c := New(0, 0)
	c.R.Add(x.R, y.R)
	c.I.Add(x.I, y.I)
	return c
}

func Multi(x, y *BigFloatComplex) *BigFloatComplex {
	c := New(0, 0)
	c.R.Sub(new(big.Float).Mul(x.R, y.R), new(big.Float).Mul(x.I, y.I))
	c.I.Add(new(big.Float).Mul(x.R, y.I), new(big.Float).Mul(x.I, y.R))
	return c
}

func Abs(v *BigFloatComplex) float64 {
	vr, _ := v.R.Float64()
	vi, _ := v.I.Float64()
	return math.Hypot(vr, vi)
}
