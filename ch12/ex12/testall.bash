#!/bin/sh

cd params
go test -v -bench=.
cd ..
