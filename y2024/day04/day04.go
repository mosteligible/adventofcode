package day04

import (
	"adventofcode/pkg/graph"
	"adventofcode/y2024/utils"
	"fmt"
)

func processInput(daydir string) [][]string {
	content := utils.ReadLines(daydir + "/input.txt")
	grid := [][]string{}
	for _, line := range content {
		row := []string{}
		for _, ch := range line {
			row = append(row, string(ch))
		}
		grid = append(grid, row)
	}
	return grid
}

func getWords(characters *[][]string, origin graph.Coordinate, rLimit int, cLimit int) [8]string {
	wordLeft, wordRight, wordUp, wordDown, leftUp, leftDown, rightUp, rightDown := "", "", "", "", "", "", "", ""

	// check up
	if origin.Row-3 >= 0 {
		wordUp = ((*characters)[origin.Row][origin.Col] + (*characters)[origin.Row-1][origin.Col] +
			(*characters)[origin.Row-2][origin.Col] + (*characters)[origin.Row-3][origin.Col])
	}

	// check down
	if origin.Row+3 < rLimit {
		wordDown = string((*characters)[origin.Row][origin.Col] + (*characters)[origin.Row+1][origin.Col] +
			(*characters)[origin.Row+2][origin.Col] + (*characters)[origin.Row+3][origin.Col])
	}

	// check left
	if origin.Col-3 >= 0 {
		wordLeft = ((*characters)[origin.Row][origin.Col] + (*characters)[origin.Row][origin.Col-1] +
			(*characters)[origin.Row][origin.Col-2] + (*characters)[origin.Row][origin.Col-3])
	}

	// check right
	if origin.Col+3 < cLimit {
		wordRight = ((*characters)[origin.Row][origin.Col] + (*characters)[origin.Row][origin.Col+1] +
			(*characters)[origin.Row][origin.Col+2] + (*characters)[origin.Row][origin.Col+3])
	}

	// diagonally left up
	if origin.Row-3 >= 0 && origin.Col-3 >= 0 {
		leftUp = ((*characters)[origin.Row][origin.Col] + (*characters)[origin.Row-1][origin.Col-1] +
			(*characters)[origin.Row-2][origin.Col-2] + (*characters)[origin.Row-3][origin.Col-3])
	}

	// diagonally left down
	if origin.Row+3 < rLimit && origin.Col-3 >= 0 {
		leftDown = ((*characters)[origin.Row][origin.Col] + (*characters)[origin.Row+1][origin.Col-1] +
			(*characters)[origin.Row+2][origin.Col-2] + (*characters)[origin.Row+3][origin.Col-3])
	}

	// diagonally right up
	if origin.Row-3 >= 0 && origin.Col+3 < cLimit {
		rightUp = ((*characters)[origin.Row][origin.Col] + (*characters)[origin.Row-1][origin.Col+1] +
			(*characters)[origin.Row-2][origin.Col+2] + (*characters)[origin.Row-3][origin.Col+3])
	}

	// diagonally right down
	if origin.Row+3 < rLimit && origin.Col+3 < cLimit {
		rightDown = ((*characters)[origin.Row][origin.Col] + (*characters)[origin.Row+1][origin.Col+1] +
			(*characters)[origin.Row+2][origin.Col+2] + (*characters)[origin.Row+3][origin.Col+3])
	}

	return [8]string{
		wordLeft, wordRight, wordUp, wordDown, leftUp, leftDown, rightUp, rightDown,
	}
}

func isXmas(grid *[][]string, origin graph.Coordinate) bool {
	leftTopRightDown := (*grid)[origin.Row-1][origin.Col-1] + (*grid)[origin.Row][origin.Col] + (*grid)[origin.Row+1][origin.Col+1]
	rightTopLeftDown := (*grid)[origin.Row-1][origin.Col+1] + (*grid)[origin.Row][origin.Col] + (*grid)[origin.Row+1][origin.Col-1]

	revLtRd := utils.ReverseString(leftTopRightDown)
	revRtLd := utils.ReverseString(rightTopLeftDown)
	if (leftTopRightDown == "MAS" || revLtRd == "MAS") && (rightTopLeftDown == "MAS" || revRtLd == "MAS") {
		return true
	}
	return false
}

func part01(grid *[][]string) {
	rows := len(*grid)
	columns := len((*grid)[0])
	found := 0
	for row := range rows {
		for col := range columns {
			if (*grid)[row][col] != "X" {
				continue
			}
			origin := graph.Coordinate{Row: row, Col: col}
			words := getWords(grid, origin, rows, columns)
			for _, w := range words {
				if w == "XMAS" {
					found += 1
				}
			}
		}
	}
	fmt.Println("Part 01:", found)
}

func part02(grid *[][]string) {
	rows := len(*grid)
	columns := len((*grid)[0])
	found := 0
	for row := 1; row < rows-1; row++ {
		for col := 1; col < columns-1; col++ {
			if (*grid)[row][col] != "A" {
				continue
			}
			if isXmas(grid, graph.Coordinate{Row: row, Col: col}) {
				found++
			}
		}
	}
	fmt.Println("Part 02:", found)
}

func Run(dir string) {
	utils.PartPrinter("DAY 04")
	data := processInput(dir + "/day04")
	part01(&data)
	part02(&data)
}
