// Copyright © 2016 shoarai
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Run with "web" command-line argument for web server.

// Surface computes an SVG rendering of a 3-D surface function.
package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"strconv"
)

// Packages not needed by version in book.
import (
	"log"
	"net/http"
	"time"
)

const (
	angle = math.Pi / 6 // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

var width, height = 600, 320               // canvas size in pixels
var xyrange = 30.0                         // axis ranges (-xyrange..+xyrange)
var xyscale = float64(width) / 2 / xyrange // pixels per x or y unit
var zscale = float64(height) * 0.4         // pixels per z unit
var cells = 100                            // number of grid cells

func main() {
	// The sequence of images is deterministic unless we seed
	// the pseudo-random number generator using the current time.
	// Thanks to Randall McPherson for pointing out the omission.
	rand.Seed(time.Now().UTC().UnixNano())

	if len(os.Args) > 1 && os.Args[1] == "web" {
		http.HandleFunc("/", handler)
		log.Fatal(http.ListenAndServe("localhost:8000", nil))
		return
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")

	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}

	for k, v := range r.Form {
		if k == "width" {
			if val, err := strconv.Atoi(v[0]); err == nil {
				setWidth(val)
			}
		} else if k == "height" {
			if val, err := strconv.Atoi(v[0]); err == nil {
				setHeight(val)
			}
		} else if k == "cells" {
			if val, err := strconv.Atoi(v[0]); err == nil {
				setCells(val)
			}
		}
	}
	fmt.Fprint(w, svg(width, height))
}

func setWidth(w int) {
	width = w                              // canvas size in pixels
	xyscale = float64(width) / 2 / xyrange // pixels per x or y unit
}

func setHeight(h int) {
	height = h                     // canvas size in pixels
	zscale = float64(height) * 0.4 // pixels per z unit
}

func setCells(c int) {
	cells = c
}

func svg(width int, height int) string {
	out := fmt.Sprintf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, ok := corner(i+1, j)
			if !ok {
				continue
			}
			bx, by, ok := corner(i, j)
			if !ok {
				continue
			}
			cx, cy, ok := corner(i, j+1)
			if !ok {
				continue
			}
			dx, dy, ok := corner(i+1, j+1)
			if !ok {
				continue
			}
			out += fmt.Sprintf("<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}

	out += "</svg>"
	return out
}

func corner(i, j int) (float64, float64, bool) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/float64(cells) - 0.5)
	y := xyrange * (float64(j)/float64(cells) - 0.5)

	// Compute surface height z.
	z, ok := f(x, y)
	if !ok {
		return 0, 0, false
	}

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := float64(width)/2 + (x-y)*cos30*xyscale
	sy := float64(height)/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, true
}

func f(x, y float64) (float64, bool) {
	r := math.Hypot(x, y) // distance from (0,0)
	if r == 0 {
		return 0, false
	}
	return math.Sin(r) / r, true
}
