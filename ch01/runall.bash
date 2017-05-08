#! /bin/sh

for dir in `ls | grep ex`; do
  cd $dir
  ./runall.bash
  cd ..
done
