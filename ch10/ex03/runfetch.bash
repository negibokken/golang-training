#! /bin/sh

cd fetch
go build
cd ..
./fetch/fetch "http://gopl.io/ch1/helloworld?go-get=1"
