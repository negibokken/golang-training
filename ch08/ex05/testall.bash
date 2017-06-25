#! /bin/sh

echo "--- with goroutine ---"
go test -v -bench=.
cd no_goroutine
echo "--- without goroutine ---"
go test -v -bench=.