package day14

import (
	"adventofcode/pkg/graph"
	"adventofcode/pkg/queue"
	"adventofcode/y2024/utils"
	"fmt"
	"strconv"
	"strings"
	"time"
)

const SAMPLE = `p=0,4 v=3,-3
p=6,3 v=-1,-3
p=10,3 v=-1,2
p=2,0 v=2,-1
p=0,0 v=1,3
p=3,0 v=-2,-2
p=7,6 v=-1,-3
p=3,0 v=-1,-2
p=9,3 v=2,3
p=7,3 v=-1,2
p=2,4 v=2,-3
p=9,5 v=-3,-3`

const (
	ROWS             = 103
	COLS             = 101
	LINEARITY_WINDOW = 12
)

type Robot struct {
	Position graph.Coordinate
	Velocity graph.Coordinate
}

func (r *Robot) String() string {
	return fmt.Sprintf(
		"Robot(Pos: %s, Vel: %s)",
		r.Position.String(), r.Velocity.String(),
	)
}

func getCoordinate(csc string) graph.Coordinate {
	coord := graph.NewCoordinate()
	numstr := strings.Split(csc, ",")
	coord.Row, _ = strconv.Atoi(numstr[1])
	coord.Col, _ = strconv.Atoi(numstr[0])
	return coord
}

func processInput(daydir string) *queue.Deque[Robot] {
	lines := utils.ReadLines(daydir + "/input.txt")
	// lines := strings.Split(SAMPLE, "\n")
	robots := queue.NewDeque[Robot]()
	for _, line := range lines {
		posvel := strings.Split(line, " ")
		posvel[0] = strings.TrimPrefix(posvel[0], "p=")
		posvel[1] = strings.TrimPrefix(posvel[1], "v=")
		robot := Robot{
			Position: getCoordinate(posvel[0]),
			Velocity: getCoordinate(posvel[1]),
		}
		robots.Append(robot)
		// fmt.Println(&robot)
	}
	return robots
}

func showGrid(grid map[graph.Coordinate]*queue.Deque[Robot]) {
	output := [ROWS][COLS]int{}
	for coord, robots := range grid {
		output[coord.Row][coord.Col] += robots.Size()
	}
	for _, row := range output {
		for _, val := range row {
			if val == 0 {
				fmt.Printf(".")
			} else {
				fmt.Printf("%d", val)
			}
		}
		fmt.Println()
	}
}

func getGrid(robots queue.Deque[Robot]) map[graph.Coordinate]*queue.Deque[Robot] {
	grid := map[graph.Coordinate]*queue.Deque[Robot]{}

	for rowNum := range ROWS {
		for colNum := range COLS {
			grid[graph.Coordinate{Row: rowNum, Col: colNum}] = queue.NewDeque[Robot]()
		}
	}
	for robots.Size() > 0 {
		r, _ := robots.Pop()
		grid[r.Position].Append(r)
	}

	return grid
}

func moveAllRobots(grid map[graph.Coordinate]*queue.Deque[Robot]) {
	newGrid := map[graph.Coordinate]*queue.Deque[Robot]{}
	for _, robots := range grid {
		for robots.Size() > 0 {
			r, err := robots.Pop()
			if err != nil {
				continue
			}
			newCoord := r.Position.Add(r.Velocity)
			if newCoord.Row >= ROWS {
				newCoord.Row -= ROWS
			} else if newCoord.Row < 0 {
				newCoord.Row += ROWS
			}
			if newCoord.Col >= COLS {
				newCoord.Col -= COLS
			} else if newCoord.Col < 0 {
				newCoord.Col += COLS
			}
			if _, ok := newGrid[newCoord]; !ok {
				newGrid[newCoord] = queue.NewDeque[Robot]()
			}
			r.Position = newCoord
			newGrid[newCoord].Append(r)
		}
	}
	for coord, robots := range newGrid {
		grid[coord] = robots
	}
}

func moveRobotNSeconds(robot *Robot, iteration int) *Robot {
	robot.Position.Row = (robot.Velocity.Row + robot.Velocity.Row*iteration) % ROWS
	robot.Position.Col = (robot.Position.Col + robot.Position.Col*iteration) % COLS

	if robot.Position.Row >= ROWS {
		robot.Position.Row -= ROWS
	} else if robot.Position.Row < 0 {
		robot.Position.Row += ROWS
	}

	if robot.Position.Col >= COLS {
		robot.Position.Col -= COLS
	} else if robot.Position.Col < 0 {
		robot.Position.Col += COLS
	}

	return robot
}

func moveNSeconds(robotGrid map[graph.Coordinate]*queue.Deque[Robot], iterations int) {

}

func getSafetyFactor(grid map[graph.Coordinate]*queue.Deque[Robot]) int {
	q1, q2, q3, q4 := 0, 0, 0, 0
	for coord, robots := range grid {
		if coord.Row < ROWS/2 && coord.Col < COLS/2 {
			q1 += robots.Size()
		} else if coord.Row < ROWS/2 && coord.Col > COLS/2 {
			q2 += robots.Size()
		} else if coord.Row > ROWS/2 && coord.Col < COLS/2 {
			q3 += robots.Size()
		} else if coord.Row > ROWS/2 && coord.Col > COLS/2 {
			q4 += robots.Size()
		}
	}
	return q1 * q2 * q3 * q4
}

func part01(robots *queue.Deque[Robot]) {
	numSeconds := 100
	robotGrid := getGrid(*robots)
	for range numSeconds {
		// showGrid(robotGrid)
		moveAllRobots(robotGrid)
		// time.Sleep(250 * time.Millisecond)
	}
	// showGrid(robotGrid)
	fmt.Println("PART 01:", getSafetyFactor(robotGrid))
}

func checkLinearity(grid map[graph.Coordinate]*queue.Deque[Robot]) bool {
	coord := graph.Coordinate{Row: 0, Col: 0}
	for rowNum := range ROWS {
		coord.Row = rowNum
		for colIndex := 0; colIndex < COLS; colIndex++ {
			if colIndex+LINEARITY_WINDOW >= COLS {
				break
			}
			numEuqals := 0
			for offset := 0; offset < LINEARITY_WINDOW; offset++ {
				coord.Col = colIndex + offset
				if grid[coord].Size() != 1 {
					break
				} else {
					numEuqals++
				}
			}
			// fmt.Printf("len tocheck: %d", len(toCheck))
			if numEuqals == LINEARITY_WINDOW {
				return true
			}
		}
	}
	return false
}

func part02(robots *queue.Deque[Robot]) {
	robotGrid := getGrid(*robots)

	totalLinearityCheck := 0.0
	moveRobotsTime := 0.0
	maxMoveTobot := -1.0
	maxLinearityTime := -1.0
	for numSeconds := range 10000 {
		start := time.Now()
		moveAllRobots(robotGrid)
		delta := time.Since(start).Seconds()
		if delta > maxMoveTobot {
			moveRobotsTime = delta
		}
		moveRobotsTime += delta
		start = time.Now()
		isLinear := checkLinearity(robotGrid)
		delta = time.Since(start).Seconds()
		if delta > maxLinearityTime {
			maxLinearityTime = delta
		}
		totalLinearityCheck += delta
		if isLinear {
			// showGrid(robotGrid)
			fmt.Println("PART 02:", numSeconds)
			fmt.Println(
				"avg checklinearity time:", totalLinearityCheck/float64(numSeconds),
				" max linear time:", maxLinearityTime,
			)
			fmt.Println(
				"move robots time:", moveRobotsTime/float64(numSeconds),
				"max robot time:", maxMoveTobot,
			)
			break
		}
	}
}

func Run(dir string) {
	utils.PartPrinter("DAY 14")
	robots := processInput(dir + "/day14")
	start := time.Now()
	part01(robots)
	fmt.Println("p01 time taken:", time.Since(start).Seconds())
	part02(robots)
	fmt.Println("time taken:", time.Since(start).Seconds())
}
