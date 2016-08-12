#!/bin/sh
cat mandelbrot.png | go run main.go -format=jpeg > out.jpg
cat mandelbrot.png | go run main.go -format=png > out.png
cat mandelbrot.png | go run main.go -format=gif > out.gif
