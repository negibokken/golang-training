#!/bin/sh

cd json
go test -v -bench=.
cd ..
