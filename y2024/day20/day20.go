package day20

import (
	"adventofcode/pkg/queue"
	"adventofcode/y2024/utils"
	"fmt"
	"math"
	"sync"
	"time"
)

const SAMPLE = `###############
#...#...#.....#
#.#.#.#.#.###.#
#S#...#.#.#...#
#######.#.#.###
#######.#.#...#
#######.#.###.#
###..E#...#...#
###.#######.###
#...###...#...#
#.#####.#.###.#
#.#...#.#.#...#
#.#.#.#.#.#.###
#...#...#...###
###############`

const (
	OBSTACLE = '#'
	START    = 'S'
	END      = 'E'

	NUM_BATCH = 16
)

type Pos struct {
	Row int
	Col int
}

func (p *Pos) HashKey() string {
	return fmt.Sprintf("(%d,%d)", p.Row, p.Col)
}

func (p *Pos) nextNodex(rowLim, colLim int, grid [][]rune, obstacle rune) []Pos {
	retval := []Pos{}
	if p.Row-1 >= 0 && grid[p.Row-1][p.Col] != obstacle {
		retval = append(retval, Pos{Row: p.Row - 1, Col: p.Col})
	}
	if p.Row+1 < rowLim && grid[p.Row+1][p.Col] != obstacle {
		retval = append(retval, Pos{Row: p.Row + 1, Col: p.Col})
	}
	if p.Col-1 >= 0 && grid[p.Row][p.Col-1] != obstacle {
		retval = append(retval, Pos{Row: p.Row, Col: p.Col - 1})
	}
	if p.Col+1 < colLim && grid[p.Row][p.Col+1] != obstacle {
		retval = append(retval, Pos{Row: p.Row, Col: p.Col + 1})
	}
	return []Pos(retval)
}

func ManhattanDistance(p1, p2 Pos) int {
	return utils.IntAbs(p1.Row-p2.Row) + utils.IntAbs(p1.Col-p2.Col)
}

type Graph struct {
	graph     map[Pos]map[Pos]int
	grid      [][]rune
	start     Pos
	end       Pos
	obstacles []Pos
}

func NewGraph(grid [][]rune) Graph {
	gridCopy := make([][]rune, len(grid))
	for row := range len(grid) {
		gridCopy[row] = make([]rune, len(grid[0]))
		copy(gridCopy[row], grid[row])
	}
	g := Graph{grid: gridCopy}
	g.fromGrid(gridCopy)
	return g
}

func (g *Graph) addEdge(c1, c2 Pos) {
	if _, ok := g.graph[c1]; !ok {
		g.graph[c1] = map[Pos]int{}
	}
	g.graph[c1][c2] = 1
}

func (g *Graph) fromGrid(grid [][]rune) {
	g.graph = map[Pos]map[Pos]int{}
	rowLim := len(grid)
	colLim := len(grid[0])
	for rowNum, row := range grid {
		for colNum, val := range row {
			currNode := Pos{Row: rowNum, Col: colNum}
			if val == OBSTACLE {
				g.obstacles = append(g.obstacles, currNode)
				continue
			}
			if val == START {
				g.start = currNode
			} else if val == END {
				g.end = currNode
			}
			for _, neighbor := range currNode.nextNodex(rowLim, colLim, grid, OBSTACLE) {
				g.addEdge(currNode, neighbor)
			}
		}
	}
}

func (g *Graph) calculateDistances() map[Pos]int {
	distances := map[Pos]int{}
	for key := range g.graph {
		distances[key] = math.MaxInt
	}
	distances[g.start] = 0
	pq := queue.NewPriorityQueue[Pos]()
	pq.Push(queue.PriorityQueueItem[Pos]{Priority: 0, Item: g.start})
	visitedNodes := map[string]bool{}
	for pq.Size() > 0 {
		currNode, err := pq.Pop()
		if err != nil {
			continue
		}
		if _, ok := visitedNodes[currNode.Item.HashKey()]; ok {
			continue
		}
		visitedNodes[currNode.Item.HashKey()] = true

		for neighbor, displacement := range g.graph[currNode.Item] {
			distToNeighbor := currNode.Priority + displacement
			if distToNeighbor < distances[neighbor] {
				distances[neighbor] = distToNeighbor
				pq.Push(queue.PriorityQueueItem[Pos]{
					Priority: distToNeighbor,
					Item:     neighbor,
				})
			}
			if neighbor == g.end {
				break
			}
		}
	}

	return distances
}

func processInput(daydir string) Graph {
	gridLines := utils.ReadLines(daydir + "/input.txt")
	grid := [][]rune{}
	for _, line := range gridLines {
		row := []rune{}
		for _, ch := range line {
			row = append(row, ch)
		}
		grid = append(grid, row)
	}

	g := NewGraph(grid)

	return g
}

func part01(graph Graph) {
	fmt.Println("num obstacles:", len(graph.obstacles))
	originalDistances := graph.calculateDistances()
	originalDist := originalDistances[graph.end]
	fmt.Println("original distance:", originalDist)
	saves := 0
	start := time.Now()
	for _, obstacle := range graph.obstacles {
		graph.grid[obstacle.Row][obstacle.Col] = '.'
		graph.fromGrid(graph.grid)
		distances := graph.calculateDistances()
		graph.grid[obstacle.Row][obstacle.Col] = '#'
		cheatDist := distances[graph.end]
		// fmt.Printf("%d obstacle, distance: %d\n", index, cheatDist)
		if originalDist-cheatDist >= 100 {
			saves++
		}
	}
	fmt.Printf("time taken: %32f seconds\n", time.Since(start).Seconds())
	fmt.Println("num saves:", saves)
}

func calculateSavesBatched(
	obstacles []Pos,
	graph Graph,
	originalDIst int,
	isSaved chan bool,
	wg *sync.WaitGroup,
) {
	defer wg.Done()
	for _, obstacle := range obstacles {
		graph.grid[obstacle.Row][obstacle.Col] = '.'
		graph.fromGrid(graph.grid)
		distances := graph.calculateDistances()
		graph.grid[obstacle.Row][obstacle.Col] = '#'
		cheatDist := distances[graph.end]

		if originalDIst-cheatDist >= 100 {
			isSaved <- true
		}
	}
}

func part01Batched(graph Graph) {
	var right int
	var wg sync.WaitGroup
	numOstacles := len(graph.obstacles)
	fmt.Println("num obstacles:", numOstacles)
	batchSize := numOstacles / NUM_BATCH
	distances := graph.calculateDistances()
	originalDist := distances[graph.end]
	fmt.Println("original distance:", originalDist)
	saves := 0
	isSaved := make(chan bool, NUM_BATCH)
	start := time.Now()
	for i := 0; i < numOstacles; i += batchSize {
		if i+batchSize < numOstacles {
			right = i + batchSize
		} else {
			right = numOstacles
		}
		wg.Add(1)
		graphCopy := NewGraph(graph.grid)
		go calculateSavesBatched(
			graph.obstacles[i:right],
			graphCopy,
			originalDist,
			isSaved,
			&wg,
		)
	}

	go func() {
		wg.Wait()
		close(isSaved)
	}()

	for res := range isSaved {
		if res {
			saves += 1
		}
	}
	fmt.Println("timetaken:", time.Since(start).Seconds())
	fmt.Println("num saves:", saves)
}

func solveOptimized(graph Graph, numCheats int) int {
	originalDistances := graph.calculateDistances()
	saves := 0
	for coord1, dist1 := range originalDistances {
		for coord2, dist2 := range originalDistances {
			c1c2Dist := ManhattanDistance(coord1, coord2)
			if c1c2Dist > numCheats {
				continue
			}
			cheatDist := dist1 + c1c2Dist
			if dist2-cheatDist >= 100 {
				saves++
			}
		}
	}
	return saves
}

func Run(dir string) {
	utils.PartPrinter("DAY 20")
	graph := processInput(dir + "/day20")
	fmt.Println("PART 01:", solveOptimized(graph, 2))
	fmt.Println("PART 02:", solveOptimized(graph, 20))
}
