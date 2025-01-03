package y2024

import (
	"adventofcode/y2024/day01"
	"adventofcode/y2024/day02"
	"adventofcode/y2024/day03"
	"adventofcode/y2024/day04"
	"adventofcode/y2024/day05"
	"adventofcode/y2024/day06"
	"adventofcode/y2024/day07"
	"adventofcode/y2024/day08"
	"adventofcode/y2024/day09"
)

func Run(wd string) {
	curdir := wd + "/y2024"
	day01.Run(curdir)
	day02.Run(curdir)
	day03.Run(curdir)
	day04.Run(curdir)
	day05.Run(curdir)
	day06.Run(curdir)
	day07.Run(curdir)
	day08.Run(curdir)
	day09.Run(curdir)
}
