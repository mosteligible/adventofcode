package utils

import (
	"math"
)

type Coordinate struct {
	Row   int
	Col   int
	Value string
}

func (c *Coordinate) Subtract(other Coordinate) Coordinate {
	return Coordinate{
		Row: c.Row - other.Row,
		Col: c.Col - other.Col,
	}
}

func (c *Coordinate) Add(other Coordinate) Coordinate {
	return Coordinate{
		Row: c.Row + other.Row,
		Col: c.Col + other.Col,
	}
}

func (c *Coordinate) Distance(other Coordinate) float32 {
	return float32(math.Sqrt(
		math.Pow((float64(c.Row-other.Row)), 2) + math.Pow(float64(c.Col-other.Col), 2),
	))
}

func (c *Coordinate) Multiply(n int) Coordinate {
	return Coordinate{
		Row: c.Row * n,
		Col: c.Col * n,
	}
}

func (c *Coordinate) Equals(other Coordinate) bool {
	return c.Row == other.Row && c.Col == other.Col
}

func (c *Coordinate) NextCoordinates(
	maxRow int, maxCol int, grid *Grid[string], obstruction string,
) []Coordinate {
	retval := []Coordinate{}
	// up
	if c.Row-1 >= 0 && grid.Grid[c.Row-1][c.Col] != obstruction {
		retval = append(retval, Coordinate{
			Row: c.Row - 1, Col: c.Col, Value: grid.Grid[c.Row-1][c.Col],
		})
	}
	// down
	if c.Row+1 < maxRow && grid.Grid[c.Row+1][c.Col] != obstruction {
		retval = append(retval, Coordinate{
			Row: c.Row + 1, Col: c.Col, Value: grid.Grid[c.Row+1][c.Col],
		})
	}
	// left
	if c.Col-1 >= 0 && grid.Grid[c.Row][c.Col-1] != obstruction {
		retval = append(retval, Coordinate{
			Row: c.Row, Col: c.Col - 1, Value: grid.Grid[c.Row][c.Col-1],
		})
	}
	//right
	if c.Col+1 < maxCol && grid.Grid[c.Row][c.Col+1] != obstruction {
		retval = append(retval, Coordinate{
			Row: c.Row, Col: c.Col + 1, Value: grid.Grid[c.Row][c.Col+1],
		})
	}

	return retval
}

func (c *Coordinate) Equal(other *Coordinate) bool {
	return c.Row == other.Row && c.Col == other.Col
}

type Grid[T any] struct {
	Grid [][]T
}

type Graph struct {
	Graph           map[Coordinate]map[Coordinate]int
	StartCoordinate Coordinate
	EndCoordinate   Coordinate
}

func NewGraph() Graph {
	return Graph{}
}

func (g *Graph) AddEdge(c1 Coordinate, c2 Coordinate, weight int) {
	if _, ok := g.Graph[c1]; !ok {
		g.Graph[c1] = map[Coordinate]int{}
	}
	g.Graph[c1][c2] = math.MaxInt
}

func (g *Graph) FromGrid(grid Grid[string], obstruction string, weight int) {
	maxRow := len(grid.Grid)
	maxCol := len(grid.Grid[0])
	for row := range maxRow {
		for col := range maxCol {
			currChar := grid.Grid[row][col]
			if currChar == obstruction {
				continue
			}
			currCoordinate := Coordinate{
				Row:   row,
				Col:   col,
				Value: currChar,
			}
			nextNodes := currCoordinate.NextCoordinates(maxRow, maxCol, &grid, obstruction)
			for _, neighborNode := range nextNodes {
				if _, ok := g.Graph[currCoordinate]; !ok {
					g.Graph[currCoordinate] = map[Coordinate]int{}
				}
				g.Graph[currCoordinate][neighborNode] = weight
			}
		}
	}
}
