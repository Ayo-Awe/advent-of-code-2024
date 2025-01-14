#!/bin/bash

day=$1

if [ -z "$day" ]; then
    echo "Usage: $0 <day>"
    exit 1
fi

if [[ ! "$day" =~ ^[0-9]+$ ]]; then
    echo "Invalid argument: $day is not a number"
    exit 1
fi

folder="day_$day"
main_file="$folder/main.go"
input_file="$folder/input.txt"

if [ ! -e "$folder" ]; then
    mkdir "$folder"
fi

if [ ! -e "$main_file" ]; then
    echo -e "package main\n\nfunc main() {\n\n}" > "$main_file"
fi

# touch does nothing if the file already exists
touch $input_file
