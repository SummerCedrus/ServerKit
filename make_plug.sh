#!/bin/bash
COMPILE_TS=$(date +%s)
cd ./plugins
files=$(ls ./)

for filename in $files
do
	mv $filename ${filename%%.*}_${COMPILE_TS}.so
	echo "$filename -> ${filename%%.*}_${COMPILE_TS}.so"
done
