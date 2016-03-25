// Copyright © 2016 shoarai
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Mandelbrot emits a super-sampling PNG image of the Mandelbrot fractal.
package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

func main() {
	img := mandelbrotImage()
	img = superSampling(img, 2)
	png.Encode(os.Stdout, img) // NOTE: ignoring errors
}

func superSampling(img image.Image, rate int) image.Image {
	bounds := img.Bounds()
	lowImg := image.NewRGBA(image.Rect(0, 0, bounds.Dx()/rate, bounds.Dy()/rate))
	var x, y int
	for py := bounds.Min.Y; py < bounds.Max.Y; py += rate {
		for px := bounds.Min.X; px < bounds.Max.X; px += rate {
			lowImg.Set(px/2, py/2, averageColor(img, px, py, rate))
			x++
		}
		y++
	}
	return lowImg
}

func averageColor(img image.Image, px, py, num int) color.Color {
	var red, blue, green uint32
	for i := px; i < px+num; i++ {
		for j := py; j < py+num; j++ {
			r, b, g, _ := img.At(i, j).RGBA()
			red += r
			blue += b
			green += g
		}
	}

	nums := uint32(num * num)
	return color.RGBA{
		uint8(red / nums),
		uint8(blue / nums),
		uint8(green / nums),
		255,
	}
}

func mandelbrotImage() image.Image {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, mandelbrot(z))
		}
	}

	return img
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.RGBA{255 - contrast*n, contrast * n, 50, 255}
		}
	}
	return color.Black
}
