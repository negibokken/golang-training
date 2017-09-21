#!/bin/sh

cd jencoder
go test -v -bench=.
cd ..
