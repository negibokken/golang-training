#! /bin/sh

go build
echo "---- zip example --------"
./ex02 testdata/test.zip
echo "---- tar example --------"
./ex02 testdata/test.tar