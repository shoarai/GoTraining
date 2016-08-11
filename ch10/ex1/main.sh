#!/bin/sh
cat mandelbrot.png | go run main.go -format=jpeg > out.jpg
