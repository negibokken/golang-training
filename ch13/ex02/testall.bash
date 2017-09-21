#!/bin/sh

cd cyclic
go test -v -bench=.
cd ..
