#! /bin/sh

go test -v -bench=.
GOARCH=386 go test -v -bench=.