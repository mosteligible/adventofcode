package y2024

import (
	"adventofcode/y2024/day01"
	"adventofcode/y2024/day03"
)

func Run(wd string) {
	curdir := wd + "/y2024"
	day01.Run(curdir)
	day03.Run(curdir)
}
