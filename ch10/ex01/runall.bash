#! /bin/sh

go build
cat gopher.jpg|./ex01 -type=png > outjp.png
cat gopher.jpg|./ex01 -type=gif> outjg.gif
cat gopher.png|./ex01 -type=jpeg > outpj.jpeg
cat gopher.png|./ex01 -type=gif > outpg.gif
cat gopher.gif|./ex01 -type=jpeg > outgj.jpeg
cat gopher.gif|./ex01 -type=png> outgp.png