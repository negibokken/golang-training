#! /bin/sh

go build
TZ=US/Eastern ./ex01 8001 2>> test.txt &
TZ=Europe/London ./ex01 8002 2>> test.txt &