#! /bin/sh

echo "GOMAXPROCS=1"
GOMAXPROCS=1 go test -v -bench=.
echo "GOMAXPROCS=2"
GOMAXPROCS=2 go test -v -bench=.
echo "GOMAXPROCS=3"
GOMAXPROCS=3 go test -v -bench=.
echo "GOMAXPROCS=4"
GOMAXPROCS=4 go test -v -bench=.