package day10

import (
	"adventofcode/y2024/utils"
)

func processInput(daydir string) [][]string {
	content := utils.ReadLines(daydir + "/input.txt")
	retval := [][]string{}
	for _, line := range content {
		row := []string{}
		for _, ch := range line {
			row = append(row, string(ch))
		}
		retval = append(retval, row)
	}
	return retval
}

func part01(grid *[][]string) {}

func Run(dir string) {
	grid := processInput(dir + "/day10")
	part01(&grid)
}
