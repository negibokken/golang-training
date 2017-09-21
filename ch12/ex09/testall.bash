#!/bin/sh

cd decode
go test -v -bench=.
cd ..
