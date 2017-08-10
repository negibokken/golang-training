#! /bin/sh

go build
GOMAXPROCS=1 ./ex06 > sample1.png
GOMAXPROCS=2 ./ex06 > sample2.png
GOMAXPROCS=3 ./ex06 > sample3.png
GOMAXPROCS=4 ./ex06 > sample4.png
