package day12

import (
	"adventofcode/pkg/graph"
	"adventofcode/y2024/utils"
	"fmt"
)

const SAMPLE = `RRRRIICCFF
RRRRIICCCF
VVRRRCCFFF
VVRCCCJFFF
VVVVCJJCFE
VVIVCCJJEE
VVIIICJJEE
MIIIIIJJEE
MIIISIJEEE
MMMISSJEEE`

func processInput(daydir string) [][]rune {
	lines := utils.ReadLines(daydir + "/input.txt")
	// lines := strings.Split(SAMPLE, "\n")
	retval := [][]rune{}
	for _, l := range lines {
		retval = append(retval, []rune(l))
	}
	return retval
}

func countEdges(grid [][]rune, row, col, rowLim, colLim int, currPlant rune) int {
	edges := 0
	// up
	if row == 0 {
		edges++
	} else if row-1 >= 0 && grid[row-1][col] != currPlant {
		edges++
	}
	// down
	if row == rowLim-1 {
		edges++
	} else if row+1 < rowLim && grid[row+1][col] != currPlant {
		edges++
	}
	// left
	if col == 0 {
		edges++
	} else if col-1 >= 0 && grid[row][col-1] != currPlant {
		edges++
	}
	// right
	if col == colLim-1 {
		edges++
	} else if col+1 <= colLim && grid[row][col+1] != currPlant {
		edges++
	}

	return edges
}

func getEdges(
	grid [][]rune, coord graph.Coordinate, rowLim, colLim int,
	edgeTracker map[graph.Coordinate]int, visitedCoords map[graph.Coordinate]bool,
) (int, int) {
	if _, ok := visitedCoords[coord]; ok {
		return 0, 0
	}
	visitedCoords[coord] = true
	currPlant := grid[coord.Row][coord.Col]
	numEdges := edgeTracker[coord]
	totalArea := 1
	if numEdges == 4 {
		return numEdges, totalArea
	}

	// up
	if coord.Row-1 >= 0 && grid[coord.Row-1][coord.Col] == currPlant {
		edge, area := getEdges(
			grid,
			graph.Coordinate{Row: coord.Row - 1, Col: coord.Col},
			rowLim, colLim, edgeTracker, visitedCoords,
		)
		numEdges += edge
		totalArea += area
	}
	// down
	if coord.Row+1 < rowLim && grid[coord.Row+1][coord.Col] == currPlant {
		edge, area := getEdges(
			grid,
			graph.Coordinate{Row: coord.Row + 1, Col: coord.Col},
			rowLim, colLim, edgeTracker, visitedCoords,
		)
		numEdges += edge
		totalArea += area
	}
	// left
	if coord.Col-1 >= 0 && grid[coord.Row][coord.Col-1] == currPlant {
		edge, area := getEdges(
			grid,
			graph.Coordinate{Row: coord.Row, Col: coord.Col - 1},
			rowLim, colLim, edgeTracker, visitedCoords,
		)
		numEdges += edge
		totalArea += area
	}
	// right
	if coord.Col+1 < colLim && grid[coord.Row][coord.Col+1] == currPlant {
		edge, area := getEdges(
			grid,
			graph.Coordinate{Row: coord.Row, Col: coord.Col + 1},
			rowLim, colLim, edgeTracker, visitedCoords,
		)
		numEdges += edge
		totalArea += area
	}

	return numEdges, totalArea
}

func makeEdgeTracker(grid [][]rune, rowLim, colLim int) map[graph.Coordinate]int {
	edgeMap := map[graph.Coordinate]int{}
	for rowNum, row := range grid {
		for colNum, val := range row {
			coord := graph.Coordinate{Row: rowNum, Col: colNum}
			edgeMap[coord] = countEdges(grid, rowNum, colNum, rowLim, colLim, val)
		}
	}
	return edgeMap
}

func part01(grid [][]rune) {
	visitedPlantCoordinates := map[graph.Coordinate]bool{}
	rowLim := len(grid)
	colLim := len(grid[0])
	edgeTracker := makeEdgeTracker(grid, rowLim, colLim)
	price := 0
	for rowNum, row := range grid {
		for colNum := range len(row) {
			coord := graph.Coordinate{Row: rowNum, Col: colNum}
			if _, ok := visitedPlantCoordinates[coord]; ok {
				continue
			}
			perimeter, area := getEdges(grid, coord, rowLim, colLim, edgeTracker, visitedPlantCoordinates)
			price += perimeter * area
		}
	}
	fmt.Println("PART 01:", price)
}

func Run(dir string) {
	utils.PartPrinter("DAY 12")
	grid := processInput(dir + "/day12")
	part01(grid)
}
