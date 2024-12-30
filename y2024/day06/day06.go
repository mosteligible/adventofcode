package day06

import (
	"adventofcode/y2024/utils"
	"crypto/sha256"
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
	Row   int
	Col   int
	Value int
}

func (p *Position) PosHashKey() string {
	h := sha256.New()
	h.Write([]byte(fmt.Sprintf("(%d,%d)", p.Row, p.Col)))
	return string(h.Sum(nil))
}

func processInput(dayDir string) [][]string {
	lines := utils.ReadLines(dayDir + "/input.txt")
	retval := [][]string{}
	for _, line := range lines {
		row := []string{}
		for _, ch := range line {
			row = append(row, string(ch))
		}
		retval = append(retval, row)
	}
	return retval
}

func reorient(direction int) int {
	switch direction {
	case east:
		return south
	case south:
		return west
	case west:
		return north
	default:
		return east
	}
}

func getTick(direction int) string {
	switch direction {
	case east:
		return ">"
	case west:
		return "<"
	case north:
		return "^"
	default:
		return "v"
	}
}

func getNextPos(currPos Position, direction int) Position {
	switch direction {
	case east:
		currPos.Col += 1
	case west:
		currPos.Col -= 1
	case north:
		currPos.Row -= 1
	case south:
		currPos.Row += 1
	default:
		log.Fatalf("received unknown direction %d", direction)
	}
	return currPos
}

func move(positions *[][]string, direction int, currPos Position) *[]Position {
	orderedPositions := []Position{}
	movedPositions := map[Position]bool{}
	rLimit, cLimit := len(*positions), len((*positions)[0])
	nextPos := currPos

	for {
		if _, ok := movedPositions[currPos]; !ok {
			orderedPositions = append(orderedPositions, currPos)
		}
		movedPositions[currPos] = true
		nextPos = getNextPos(currPos, direction)
		if nextPos.Row >= rLimit || nextPos.Row < 0 ||
			nextPos.Col >= cLimit || nextPos.Col < 0 {
			break
		}
		if (*positions)[nextPos.Row][nextPos.Col] == "#" {
			direction = reorient(direction)
			continue
		}
		currPos = nextPos
	}
	return &orderedPositions
}

func moveLoop(positions *[][]string, currPos Position, size Position) bool {
	movedPositions := map[Position]bool{}
	nextPos := currPos
	for {
		if _, ok := movedPositions[currPos]; ok {
			return true
		} else {
			movedPositions[currPos] = true
		}
		nextPos = getNextPos(currPos, currPos.Value)
		nextPos.Value = currPos.Value
		if nextPos.Row >= size.Row || nextPos.Row < 0 ||
			nextPos.Col >= size.Col || nextPos.Col < 0 {
			return false
		}
		if (*positions)[nextPos.Row][nextPos.Col] == "#" {
			currPos.Value = reorient(nextPos.Value)
			continue
		}
		currPos = nextPos
	}
}

func getGuardPositionOrientation(data *[][]string) (Position, int) {
	for rowNum, row := range *(data) {
		for colNum, ch := range row {
			if ch == "^" {
				return Position{
					Row: rowNum,
					Col: colNum,
				}, north
			}
		}
	}
	log.Fatal("could not find guard!")
	return Position{}, -1
}

func copyGrid(g *[][]string) *[][]string {
	copiedGrid := [][]string{}

	for rNum := range len(*g) {
		row := []string{}
		for cNum := range len((*g)[0]) {
			row = append(row, (*g)[rNum][cNum])
		}
		copiedGrid = append(copiedGrid, row)
	}

	return &copiedGrid
}

func part01(data *[][]string) *[]Position {
	guardPos, orientation := getGuardPositionOrientation(data)
	positions := move(data, orientation, guardPos)
	fmt.Println("Part 01:", len(*positions))
	return positions
}

func part02(data *[][]string) {
	positions := part01(data)
	size := Position{Row: len(*data), Col: len((*data)[0])}
	guardPos, orientation := getGuardPositionOrientation(data)
	guardPos.Value = orientation
	numLoops := 0
	for _, pos := range (*positions)[1:] {
		(*data)[pos.Row][pos.Col] = "#"
		if moveLoop(data, guardPos, size) {
			numLoops++
		}
		(*data)[pos.Row][pos.Col] = "."
	}
	fmt.Println("Part 02:", numLoops)
}

func Run(dir string) {
	utils.PartPrinter("DAY 06")
	data := processInput(dir + "/day06")
	part02(&data)
}
