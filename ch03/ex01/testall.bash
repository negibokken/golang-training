#! /bin/sh

go test -v -bench=. | grep -Ev "^./ex01"
