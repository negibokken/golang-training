#!/bin/sh

cd equalish
go test -v -bench=.
cd ..
