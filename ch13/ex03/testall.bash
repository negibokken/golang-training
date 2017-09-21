#!/bin/sh

cd bzip
go test -v -bench=.
cd ..
