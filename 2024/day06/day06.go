package day06

import (
	"aoc2024/utils"
	"fmt"
	"log"
)

var POSITIONS = []string{
	"....#.....",
	".........#",
	"..........",
	"..#.......",
	".......#..",
	"..........",
	".#..^.....",
	"........#.",
	"#.........",
	"......#...",
}

const (
	east  = iota
	west  = iota
	north = iota
	south = iota
)

type Position struct {
	row int
	col int
}

func processInput() [][]rune {
	lines := utils.ReadLines("./day06/input.txt")
	retval := [][]rune{}
	for _, line := range lines {
		row := []rune{}
		for _, ch := range line {
			row = append(row, ch)
		}
		retval = append(retval, row)
	}
	return retval
}

func getGuardPositionOrientation(data [][]rune) (Position, int) {
	for rowNum, row := range data {
		for colNum, ch := range row {
			if ch == '^' {
				return Position{
					row: rowNum,
					col: colNum,
				}, north
			}
		}
	}
	log.Fatal("could not find guard!")
	return Position{}, -1
}

func part01(data [][]rune) {
	guardPos, orientation := getGuardPositionOrientation(data)
	fmt.Printf("guardPos: %v, orientation: %d\n", guardPos, orientation)
}

func Run() {
	data := processInput()
	part01(data)
}
