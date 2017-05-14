#! /bin/sh

go build ex02.go
./ex02 test
./ex02 -hash SHA384 test
./ex02 -hash SHA512 test
