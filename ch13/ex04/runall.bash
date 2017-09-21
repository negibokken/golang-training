#! /bin/sh
go build
cat test.txt | ./ex04 > out.bz2
