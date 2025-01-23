package day10

import (
	"adventofcode/pkg/graph"
	"adventofcode/y2024/utils"
	"fmt"
	"slices"
)

const SAMPLE = `89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732`

var ORIGIN_DESTINATION_TRACKER = map[graph.Coordinate][]graph.Coordinate{}

func processInput(daydir string) [][]int {
	content := utils.ReadLines(daydir + "/input.txt")
	retval := [][]int{}
	for _, line := range content {
		row := []int{}
		for _, ch := range line {
			row = append(row, int(ch-'0'))
		}
		retval = append(retval, row)
	}
	return retval
}

func isDestinationVisited(row, col int) bool {
	val := graph.Coordinate{Row: row, Col: col}
	arr, ok := ORIGIN_DESTINATION_TRACKER[val]
	if !ok {
		return false
	}
	return slices.Contains(arr, val)
}

func trailPathVisited(trail []graph.Coordinate, point graph.Coordinate) bool {
	for _, coord := range trail {
		if coord.Equals(point) {
			return true
		}
	}
	return false
}

func trailheads(
	data [][]int,
	row, col, prevNo int,
	size, origin graph.Coordinate, start bool,
	trailtracker []graph.Coordinate,
	isp01 bool,
) int {
	currNo := data[row][col]
	currCoord := graph.Coordinate{Row: row, Col: col}
	numRight, numLeft, numUp, numDown := 0, 0, 0, 0
	if !start {
		if currNo-prevNo != 1 {
			return 0
		}
		if currNo == 9 && prevNo != 0 {
			if isDestinationVisited(row, col) {
				return 0
			}
			ORIGIN_DESTINATION_TRACKER[origin] = append(
				ORIGIN_DESTINATION_TRACKER[origin],
				currCoord,
			)
			trailtracker = append(trailtracker, currCoord)
			return 1
		}
	} else {
		start = false
	}

	if isp01 && trailPathVisited(trailtracker, graph.Coordinate{Row: row, Col: col}) {
		return 0
	}
	if row-1 >= 0 {
		numUp = trailheads(data, row-1, col, currNo, size, origin, start, trailtracker, isp01)
		if numUp > 0 {
			trailtracker = append(trailtracker, graph.Coordinate{Row: row, Col: col})
		}
	}

	if row+1 < size.Row {
		numDown = trailheads(data, row+1, col, currNo, size, origin, start, trailtracker, isp01)
		if numDown > 0 {
			trailtracker = append(trailtracker, graph.Coordinate{Row: row, Col: col})
		}
	}

	if col-1 >= 0 {
		numLeft = trailheads(data, row, col-1, currNo, size, origin, start, trailtracker, isp01)
	}

	if col+1 < size.Col {
		numRight = trailheads(data, row, col+1, currNo, size, origin, start, trailtracker, isp01)
	}

	return numRight + numLeft + numDown + numUp
}

func solve(grid *[][]int, isP01 bool) {
	size := graph.Coordinate{Row: len(*grid), Col: len((*grid)[0])}
	numTrails := 0
	trailtracker := []graph.Coordinate{}
	for rowNum, row := range *grid {
		for colNum, val := range row {
			if val != 0 {
				continue
			}
			origin := graph.Coordinate{Row: rowNum, Col: colNum}
			ORIGIN_DESTINATION_TRACKER[origin] = []graph.Coordinate{}
			nTrail := trailheads(*grid, rowNum, colNum, 0, size, origin, true, trailtracker, isP01)
			numTrails += nTrail
		}
	}
	fmt.Println("PART 01:", numTrails)
}

func Run(dir string) {
	utils.PartPrinter("DAY 10")
	grid := processInput(dir + "/day10")
	solve(&grid, true)
	solve(&grid, false)
}
