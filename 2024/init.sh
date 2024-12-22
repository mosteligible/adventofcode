#!/bin/bash

day=${1,}

if [[ -z "$day" ]]; then
    echo "-- which day to init for has not been provided in argument."
fi

if [[ -d "$day" ]]; then
    echo "-- day $day already exists."
    exit 1
fi

mkdir "$day"

cd "$day"
touch "__init__.py"
echo "def process_input():
    with open(\"./${day}/input.txt\", \"r\") as fp:
        content = fp.readlines()
    return content


def part01(data):
    ...


def part02(data):
    ...


def main():
    data = process_input()
    part01(data)
    part02(data)


if __name__ == \"__main__\":
    main()" > "$day.py"
touch "input.txt"
