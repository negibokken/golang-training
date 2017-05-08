#! /bin/sh

for dir in `ls | grep ex`; do
  cd $dir
  ./testall.bash
  cd ..
done
