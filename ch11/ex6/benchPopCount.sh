#!/bin/sh
paths=(\
	"gopl.io/ch2/popcount"\
	"github.com/shoarai/GoTraining/ch2/ex4"\
	"github.com/shoarai/GoTraining/ch2/ex5")
	
for path in "${paths[@]}"; do
	go test -bench=. -benchmem $path
	echo ""
done
