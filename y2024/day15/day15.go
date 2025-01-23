package day15

import (
	"adventofcode/y2024/utils"
	"fmt"
	"log"
	"time"
)

const (
	ROBOT_INDICATOR = '@'
	CRATE           = 'O'
	EMPTY_SPACE     = '.'
	WALL            = '#'
)

type RobotPos struct {
	Row int
	Col int
}

func (r *RobotPos) Equal(other RobotPos) bool {
	return r.Row == other.Row && r.Col == other.Col
}

func moveRobot(direction rune, pos RobotPos) RobotPos {
	switch direction {
	case '^':
		return RobotPos{Row: pos.Row - 1, Col: pos.Col}
	case 'v':
		return RobotPos{Row: pos.Row + 1, Col: pos.Col}
	case '<':
		return RobotPos{Row: pos.Row, Col: pos.Col - 1}
	default:
		return RobotPos{Row: pos.Row, Col: pos.Col + 1}
	}
}

func showGrid(grid [][]rune) {
	for _, row := range grid {
		for _, val := range row {
			fmt.Printf("%c", val)
		}
		fmt.Println()
	}
}

func getRobotPosition(grid [][]rune) (RobotPos, error) {
	for rowNum, row := range grid {
		for colNum, val := range row {
			if val == ROBOT_INDICATOR {
				return RobotPos{Row: rowNum, Col: colNum}, nil
			}
		}
	}
	return utils.GetZero[RobotPos](), fmt.Errorf("could not find robot position")
}

func processInput(daydir string) ([][]rune, string) {
	moves := utils.Read(daydir + "/input_robot_movement.txt")

	warehouseLines := utils.ReadLines(daydir + "/input_warehouse.txt")
	grid := [][]rune{}
	for _, line := range warehouseLines {
		row := []rune{}
		for _, ch := range line {
			row = append(row, ch)
		}
		grid = append(grid, row)
	}

	return grid, moves
}

func adjustCrates(grid *[][]rune, pos RobotPos, direction rune) RobotPos {
	nextPos := moveRobot(direction, pos)
	nextCh := (*grid)[nextPos.Row][nextPos.Col]
	switch nextCh {
	case CRATE:
		if !nextPos.Equal(pos) {
			(*grid)[nextPos.Row][nextPos.Col] = CRATE
			(*grid)[pos.Row][pos.Col] = EMPTY_SPACE
		}
		return nextPos
	case WALL:
		return pos
	case EMPTY_SPACE:
		if (*grid)[pos.Row][pos.Col] == CRATE {
			(*grid)[nextPos.Row][nextPos.Col] = CRATE
			(*grid)[pos.Row][pos.Col] = EMPTY_SPACE
		}
		return nextPos
	}
	log.Fatalf("invalid 7character: %c", nextCh)
	return utils.GetZero[RobotPos]()
}

func part01(grid [][]rune, moves string) {
	robotPosition, err := getRobotPosition(grid)
	if err != nil {
		log.Fatalf("Could not find robot in grid!!")
	}
	for _, ch := range moves {
		nextRobotPos := moveRobot(ch, robotPosition)
		switch grid[nextRobotPos.Row][nextRobotPos.Col] {
		case '.':
			grid[robotPosition.Col][robotPosition.Col] = EMPTY_SPACE
			grid[nextRobotPos.Row][nextRobotPos.Col] = ROBOT_INDICATOR
			robotPosition = nextRobotPos
		case '#':
			continue
		case 'O':
			nextRobotPos = adjustCrates(&grid, robotPosition, ch)
			if !nextRobotPos.Equal(robotPosition) {
				grid[nextRobotPos.Row][nextRobotPos.Col] = ROBOT_INDICATOR
				grid[robotPosition.Row][robotPosition.Col] = EMPTY_SPACE
				robotPosition = nextRobotPos
			}
		}
		utils.ClearScreen()
		showGrid(grid)
		time.Sleep(500 * time.Millisecond)
	}
}

func Run(dir string) {
	grid, moves := processInput(dir + "/day15")
	part01(grid, moves)
}
