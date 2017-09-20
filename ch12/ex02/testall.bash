#! /bin/sh

cd display
go test -v -bench=.
cd ..