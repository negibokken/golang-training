#!/bin/sh

cd bzip2
go test -v -bench=.
cd ..
