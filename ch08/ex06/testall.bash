#! /bin/sh

go test -v -bench=.
cd no_goroutine