package day08

import (
	"adventofcode/pkg/graph"
	"adventofcode/y2024/utils"
	"fmt"
)

const (
	EPSILON = 0.009
	SAMPLE  = `............
........0...
.....0......
.......0....
....0.......
......A.....
............
............
........A...
.........A..
............
............`
)

func processInput(dayDir string) *[][]string {
	grid := [][]string{}

	lines := utils.ReadLines(dayDir + "/input.txt")
	// lines := strings.Split(SAMPLE, "\n")
	for _, l := range lines {
		row := []string{}
		for _, r := range l {
			row = append(row, string(r))
		}
		grid = append(grid, row)
	}

	return &grid
}

func getAntennaCoordinates(grid *[][]string) map[string][]graph.Coordinate {
	coordinates := map[string][]graph.Coordinate{}
	for rowNum, row := range *grid {
		for colNum, colVal := range row {
			coord := graph.Coordinate{Row: rowNum, Col: colNum}
			if colVal != "." {
				if _, ok := coordinates[colVal]; !ok {
					coordinates[colVal] = []graph.Coordinate{}
				}
				coordinates[colVal] = append(coordinates[colVal], coord)
			}
		}
	}
	return coordinates
}

func getAntinodePositions(
	nodeCoordinates []graph.Coordinate,
	size graph.Coordinate,
	antennas map[graph.Coordinate]bool,
	antinodes *map[graph.Coordinate]bool,
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
	size := graph.Coordinate{
		Row: len(grid),
		Col: len(grid[0]),
	}
	antennaCoordinates := getAntennaCoordinates(&grid)
	antennas := map[graph.Coordinate]bool{}
	uniqueAntinodes := map[graph.Coordinate]bool{}
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
