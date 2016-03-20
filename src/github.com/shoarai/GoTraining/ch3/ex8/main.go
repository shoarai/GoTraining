// Copyright © 2016 shoarai
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Mandelbrot emits a PNG image of the Mandelbrot fractal in full color.
package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"

	"github.com/shoarai/GoTraining/ch3/ex8/bigfloatcmplx"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			// Image point (px, py) represents complex value z.
			img.Set(px, py, mandelbrotBigFloat(x, y))
		}
	}
	png.Encode(os.Stdout, img) // NOTE: ignoring errors
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
