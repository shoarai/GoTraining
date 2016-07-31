#!/bin/sh
paths=(\
	"github.com/shoarai/GoTraining/ch11/ex7/intset"\
	"github.com/shoarai/GoTraining/ch11/ex7/intsetmap")

for path in "${paths[@]}"; do
	go test -bench=. -benchmem $path
	echo ""
done
