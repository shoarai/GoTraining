#!/bin/sh
go build main.go
./main "go run ../../ch8/ex5/mandelbrot/main.go" > measure.out
