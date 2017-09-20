#! /bin/sh

cd display
go test -v -bench=.
cd ..
cd format
go test -v -bench=.
cd ..