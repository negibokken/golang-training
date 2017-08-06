#! /bin/sh

cd popcount
go test -v -bench=.
cd ..