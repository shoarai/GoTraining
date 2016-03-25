// Copyright © 2016 shoarai
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Mandelbrot emits a PNG image of the Mandelbrot fractal in full color.
// Example:
// $ github.com/shoarai/GoTraining/ch3/ex8/main.go
// $ ./main.go > mandelbrot.png
// $ ./main.go cmplx128 > mandelbrot.png
// $ ./main.go bigfloat > mandelbrot.png
// $ ./main.go bigrat > mandelbrot.png
package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"

	"github.com/shoarai/GoTraining/ch3/ex8/bigfloatcmplx"
	"github.com/shoarai/GoTraining/ch3/ex8/bigratcmplx"
)

const (
	Cmplx64 = iota
	Cmplx128
	BigFloat
	BigRat
)

func main() {
	cmplxType := Cmplx64
	if len(os.Args[1:]) > 0 {
		switch os.Args[1] {
		case "cmplx128":
			cmplxType = Cmplx128
		case "bigfloat":
			cmplxType = BigFloat
		case "bigrat":
			cmplxType = BigRat
		}
	}
	png.Encode(os.Stdout, mandelbrot(cmplxType)) // NOTE: ignoring errors
}

func mandelbrot(cmplxType int) *image.RGBA {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			var c color.Color
			switch cmplxType {
			case Cmplx64:
				c = mandelbrot64(x, y)
			case Cmplx128:
				c = mandelbrot128(x, y)
			case BigFloat:
				c = mandelbrotBigFloat(x, y)
			case BigRat:
				c = mandelbrotBigRat(x, y)
			}
			img.Set(px, py, c)
		}
	}

	return img
}

func mandelbrot64(x, y float64) color.Color {
	const iterations = 200
	const contrast = 15

	z := complex64(complex(x, y))
	var v complex64
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(complex128(v)) > 2 {
			return color.RGBA{255 - contrast*n, contrast * n, 50, 255}
		}
	}
	return color.Black
}

func mandelbrot128(x, y float64) color.Color {
	const iterations = 200
	const contrast = 15

	z := complex(x, y)
	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.RGBA{255 - contrast*n, contrast * n, 50, 255}
		}
	}
	return color.Black
}

func mandelbrotBigFloat(x, y float64) color.Color {
	const iterations = 200
	const contrast = 15

	z := bigfloatcmplx.New(x, y)
	v := bigfloatcmplx.New(0, 0)
	for n := uint8(0); n < iterations; n++ {
		v = bigfloatcmplx.Multi(v, v)
		v = bigfloatcmplx.Add(v, z)
		if bigfloatcmplx.Abs(v) > 2 {
			return color.RGBA{255 - contrast*n, contrast * n, 50, 255}
		}
	}
	return color.Black
}

func mandelbrotBigRat(x, y float64) color.Color {
	const iterations = 200
	const contrast = 15

	z := bigratcmplx.New(x, y)
	v := bigratcmplx.New(0, 0)
	for n := uint8(0); n < iterations; n++ {
		v = bigratcmplx.Multi(v, v)
		v = bigratcmplx.Add(v, z)
		if bigratcmplx.Abs(v) > 2 {
			return color.RGBA{255 - contrast*n, contrast * n, 50, 255}
		}
	}
	return color.Black
}
