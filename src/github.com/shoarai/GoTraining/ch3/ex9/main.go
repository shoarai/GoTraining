// Copyright © 2016 shoarai
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Mandelbrot emits a PNG image of the Mandelbrot fractal in full color.
//
// Mandelbrot is a PNG image of the Mandelbrot server.
// The queries are "x", "y" of axises and "scale" of the image scale
package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"math/cmplx"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/png")

	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}

	var x, y float64
	scale := 1.0
	for k, v := range r.Form {
		if k == "x" {
			if val, err := strconv.ParseFloat(v[0], 64); err == nil {
				x = val
			}
		} else if k == "y" {
			if val, err := strconv.ParseFloat(v[0], 64); err == nil {
				y = val
			}
		} else if k == "scale" {
			if val, err := strconv.ParseFloat(v[0], 64); err == nil {
				scale = val
			}
		}
	}
	svg(w, x, y, scale)
}

func svg(w http.ResponseWriter, x, y, scale float64) {
	const (
		width, height = 1024, 1024
	)

	base := 2.0
	xmin := -base / scale
	ymin := xmin
	xmax := base / scale
	ymax := xmax

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin + y
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin + x
			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, mandelbrot(z))
		}
	}

	png.Encode(w, img) // NOTE: ignoring errors
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
