#! /bin/sh

go build ex08.go
./ex08 --type complex64 > 64.png
./ex08 --type complex128 > 128.png
./ex08 --type big.Float > Float.png
./ex08 --type big.Rat > Rat.png
