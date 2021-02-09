#!/bin/bash

while IFS= read -r -d '' dir; do
  # echo $dir
	binary_name=$(echo $dir | sed -e 's/\(.* - \)*//g')
  # echo $binary_name
	go build -o bin/$binary_name "$dir"
done < <(find . -type d -name "* - *" -print0)
