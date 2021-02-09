#!/bin/bash

time while IFS= read -r -d '' dir; do
  binary_name=$(echo $dir | sed -e 's/\(.* - \)*//g')
	go build -o bin/$binary_name "$dir"
done < <(find . -type d -name "* - *" -print0)
