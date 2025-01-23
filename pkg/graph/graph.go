package graph

import (
	"math"
)

type Grid[T any] struct {
	Grid [][]T
}

type Graph struct {
	Graph       map[string]map[string]int
	Grid        [][]string
	obstruction string
}

func NewGraph(grid [][]string, obstruction string) Graph {
	g := Graph{
		Graph:       map[string]map[string]int{},
		Grid:        grid,
		obstruction: obstruction,
	}
	g.FromGrid(obstruction, 1)
	return g
}

func (g *Graph) GetObstruction() string {
	return g.obstruction
}

func (g *Graph) SetObstruction(newObstruction string) {
	g.obstruction = newObstruction
}

func (g *Graph) AddEdge(c1 PositionHandler, c2 PositionHandler, weight int) {
	if _, ok := g.Graph[c1.HashKey()]; !ok {
		g.Graph[c1.HashKey()] = map[string]int{}
	}
	g.Graph[c1.HashKey()][c2.HashKey()] = math.MaxInt
}

func (g *Graph) FromGrid(obstruction string, weight int) {
	maxRow := len(g.Grid)
	maxCol := len(g.Grid)
	for row := range maxRow {
		for col := range maxCol {
			currChar := g.Grid[row][col]
			if currChar == obstruction {
				continue
			}
			currCoordinate := Coordinate{
				Row:   row,
				Col:   col,
				Value: currChar,
			}
			nextNodes := currCoordinate.NextCoordinates(maxRow, maxCol, &g.Grid, obstruction)
			for _, neighborNode := range nextNodes {
				g.AddEdge(&currCoordinate, &neighborNode, weight)
			}
		}
	}
}
