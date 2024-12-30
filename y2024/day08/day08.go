package day08

import (
	"adventofcode/y2024/utils"
	"fmt"
)

const (
	EPSILON = 0.009
)

func processInput(dayDir string) *[][]string {
	grid := [][]string{}

	lines := utils.ReadLines(dayDir + "/input.txt")
	for _, l := range lines {
		row := []string{}
		for _, r := range l {
			row = append(row, string(r))
		}
		grid = append(grid, row)
	}

	return &grid
}

func getAntennaCoordinates(grid *[][]string) map[string][]utils.Coordinate {
	coordinates := map[string][]utils.Coordinate{}
	for rowNum, row := range *grid {
		for colNum, colVal := range row {
			coord := utils.Coordinate{Row: rowNum, Col: colNum}
			if colVal != "." {
				if _, ok := coordinates[colVal]; !ok {
					coordinates[colVal] = []utils.Coordinate{}
				}
				coordinates[colVal] = append(coordinates[colVal], coord)
			}
		}
	}
	return coordinates
}

func getAntinodePositions(
	nodeCoordinates []utils.Coordinate,
	size utils.Coordinate,
	antennas map[utils.Coordinate]bool,
	antinodes *map[utils.Coordinate]bool,
) {
	for coordIndex, c0 := range nodeCoordinates {
		for _, c1 := range nodeCoordinates[coordIndex+1:] {
			delta := c0.Subtract(c1)
			potentialAntinode0 := c0.Add(delta)
			potentialAntinode1 := c0.Subtract(delta.Multiply(2))
			_, pa0Ok := antennas[potentialAntinode0]
			if potentialAntinode0.Row >= 0 && potentialAntinode0.Row < size.Row &&
				potentialAntinode0.Col >= 0 && potentialAntinode0.Col < size.Col &&
				!pa0Ok {
				(*antinodes)[potentialAntinode0] = true
			}
			_, pa1Ok := antennas[potentialAntinode1]
			if potentialAntinode1.Row >= 0 && potentialAntinode1.Row < size.Row &&
				potentialAntinode1.Col >= 0 && potentialAntinode1.Col < size.Col &&
				!pa1Ok {
				(*antinodes)[potentialAntinode1] = true
			}
		}
	}
}

func part01(grid [][]string) {
	size := utils.Coordinate{
		Row: len(grid),
		Col: len(grid[0]),
	}
	antennaCoordinates := getAntennaCoordinates(&grid)
	antennas := map[utils.Coordinate]bool{}
	uniqueAntinodes := map[utils.Coordinate]bool{}
	for _, val := range antennaCoordinates {
		for _, a := range val {
			antennas[a] = true
		}
	}
	for _, coordinates := range antennaCoordinates {
		getAntinodePositions(
			coordinates, size, antennas, &uniqueAntinodes,
		)
	}
	fmt.Println("Part 01:", len(uniqueAntinodes))
}

func Run(dir string) {
	utils.PartPrinter("DAY 08")
	grid := processInput(dir + "/day08")
	part01(*grid)
}
