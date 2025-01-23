package day16

import (
	"adventofcode/pkg/graph"
	"adventofcode/pkg/queue"
	"adventofcode/y2024/utils"
	"fmt"
	"log"
	"math"
)

const SAMPLE = `###############
#.......#....E#
#.#.###.#.###.#
#.....#.#...#.#
#.###.#####.#.#
#.#.#.......#.#
#.#.#####.###.#
#...........#.#
###.#.#####.#.#
#...#.....#.#.#
#.#.#.###.#.#.#
#.....#...#.#.#
#.###.#.#.#.#.#
#S..#.....#...#
###############`

const (
	east   = iota
	south  = iota
	west   = iota
	north  = iota
	notset = iota

	START_INDICATOR = "S"
	END_INDICATOR   = "E"
)

type Pos struct {
	Row       int
	Col       int
	Direction int
}

func (c *Pos) HashKey() string {
	return fmt.Sprintf("(%d,%d)", c.Row, c.Col)
}

func (c *Pos) NextCoordinates(
	maxRow, maxCol int, grid [][]string, obstruction string,
) []Pos {
	retval := []Pos{}

	if c.Row-1 >= 0 && (grid)[c.Row-1][c.Col] != obstruction {
		retval = append(retval, Pos{
			Row: c.Row - 1, Col: c.Col,
		})
	}
	if c.Row+1 < maxRow && (grid)[c.Row+1][c.Col] != obstruction {
		retval = append(retval, Pos{
			Row: c.Row + 1, Col: c.Col,
		})
	}
	if c.Col-1 >= 0 && (grid)[c.Row][c.Col-1] != obstruction {
		retval = append(retval, Pos{
			Row: c.Row, Col: c.Col - 1,
		})
	}
	if c.Col+1 < maxCol && (grid)[c.Row][c.Col+1] != obstruction {
		retval = append(retval, Pos{
			Row: c.Row, Col: c.Col + 1,
		})
	}

	return retval
}

func processInput(daydir string) graph.Graph {
	gridLines := utils.ReadLines(daydir + "/input.txt")
	// gridLines := strings.Split(SAMPLE, "\n")
	grid := [][]string{}
	for _, line := range gridLines {
		row := []string{}
		for _, ch := range line {
			row = append(row, string(ch))
		}
		grid = append(grid, row)
	}
	graph := graph.NewGraph(grid, "#")
	return graph
}

func getPositionOf(grid [][]string, ch string) (Pos, error) {
	for rowNum, row := range grid {
		for colNum, val := range row {
			if val == ch {
				return Pos{Row: rowNum, Col: colNum}, nil
			}
		}
	}
	return utils.GetZero[Pos](), fmt.Errorf("could not find character <%s> in grid", ch)
}

func initDistances(graph graph.Graph) map[string]int {
	distances := map[string]int{}
	for coord := range graph.Graph {
		distances[coord] = math.MaxInt
	}
	return distances
}

func calculateDistance(c1 Pos, c2 Pos) (int, int) {
	var toNeighbor int
	if c2.Row == c1.Row+1 {
		toNeighbor = south
	} else if c2.Row == c1.Row-1 {
		toNeighbor = north
	} else if c2.Col == c1.Col-1 {
		toNeighbor = west
	} else {
		toNeighbor = east
	}
	diff := utils.IntAbs(toNeighbor - c1.Direction)
	switch diff {
	case 0:
		return 1, toNeighbor
	case 1, 2:
		return 1000 * diff, toNeighbor
	case 3:
		return 1000 + 1, toNeighbor
	}
	log.Fatalf("got invalid direction change: %d | valid directions are 0,1,2,3", toNeighbor)
	return -1, -1
}

func ShortestPath(graph graph.Graph, start, end Pos) map[string]int {
	if len(graph.Graph) == 0 {
		log.Fatalf("graph has not been initialized!")
	}
	visitedNodes := map[string]bool{}
	distances := initDistances(graph)
	numRows := len(graph.Grid)
	numCols := len(graph.Grid[0])
	distances[start.HashKey()] = 0
	pq := queue.NewPriorityQueue[Pos]()
	pq.Push(queue.PriorityQueueItem[Pos]{Priority: 0, Item: start})
	for pq.Size() > 0 {
		currPos, err := pq.Pop()
		if err != nil {
			continue
		}
		if _, ok := visitedNodes[currPos.Item.HashKey()]; ok {
			continue
		}
		visitedNodes[currPos.Item.HashKey()] = true
		for _, neighbor := range currPos.Item.NextCoordinates(numRows, numCols, graph.Grid, graph.GetObstruction()) {
			dist, newOrientation := calculateDistance(currPos.Item, neighbor)
			neighbor.Direction = newOrientation
			tentativeDist := dist + currPos.Priority
			if tentativeDist <= distances[neighbor.HashKey()] {
				distances[neighbor.HashKey()] = tentativeDist
				pq.Push(queue.PriorityQueueItem[Pos]{Priority: tentativeDist, Item: neighbor})
			}
		}
	}
	return distances
}

func part01(graph graph.Graph) {
	startPos, err := getPositionOf(graph.Grid, START_INDICATOR)
	if err != nil {
		log.Fatalf("could not find start position")
	}
	endPos, err := getPositionOf(graph.Grid, END_INDICATOR)
	if err != nil {
		log.Fatalf("could not find end position")
	}
	distances := ShortestPath(graph, startPos, endPos)
	fmt.Println("PART 01:", distances[endPos.HashKey()])
}

func Run(dir string) {
	utils.PartPrinter("DAY 16")
	graph := processInput(dir + "/day16")
	part01(graph)
}
