package day13

import (
	"adventofcode/y2024/utils"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

const SAMPLE = `Button A: X+94, Y+34
Button B: X+22, Y+67
Prize: X=8400, Y=5400

Button A: X+26, Y+66
Button B: X+67, Y+21
Prize: X=12748, Y=12176

Button A: X+17, Y+86
Button B: X+84, Y+37
Prize: X=7870, Y=6450

Button A: X+69, Y+23
Button B: X+27, Y+71
Prize: X=18641, Y=10279`

var (
	X_MOVE_REGEX = regexp.MustCompile(`X\+(\d+),`)
	Y_MOVE_REGEX = regexp.MustCompile(`Y\+(\d+)`)
	X_PRIZE_MOVE = regexp.MustCompile(`X=(\d+),`)
	Y_PRIZE_MOVE = regexp.MustCompile(`Y=(\d+)`)
)

type Button struct {
	x int
	y int
}

type Prize struct {
	x int
	y int
}

type ArcadeMachine struct {
	A     Button
	B     Button
	Prize Prize
}

func newArcadeMachine() ArcadeMachine {
	return ArcadeMachine{
		A:     Button{x: -1, y: -1},
		B:     Button{x: -1, y: -1},
		Prize: Prize{x: -1, y: -1},
	}
}

func processInput(daydir string) []ArcadeMachine {
	arcadeMachines := []ArcadeMachine{}

	lines := utils.ReadLines(daydir + "/input.txt")
	// lines := strings.Split(SAMPLE, "\n")
	anArcadeMachine := newArcadeMachine()
	for _, line := range lines {
		if strings.HasPrefix(line, "Button") {
			buttonName := line[7]
			xMoveStr := X_MOVE_REGEX.FindStringSubmatch(line)[1]
			yMoveStr := Y_MOVE_REGEX.FindStringSubmatch(line)[1]
			xMoveInt, _ := strconv.Atoi(xMoveStr)
			yMoveInt, _ := strconv.Atoi(yMoveStr)
			if buttonName == byte('A') {
				anArcadeMachine.A = Button{x: xMoveInt, y: yMoveInt}
			} else {
				anArcadeMachine.B = Button{x: xMoveInt, y: yMoveInt}
			}
		} else if strings.HasPrefix(line, "Prize") {
			xMoveStr := X_PRIZE_MOVE.FindStringSubmatch(line)[1]
			yMoveStr := Y_PRIZE_MOVE.FindStringSubmatch(line)[1]
			xMoveInt, _ := strconv.Atoi(xMoveStr)
			yMoveInt, _ := strconv.Atoi(yMoveStr)
			anArcadeMachine.Prize = Prize{x: xMoveInt, y: yMoveInt}
			arcadeMachines = append(arcadeMachines, anArcadeMachine)
		}
	}

	return arcadeMachines
}

func numTokens(l1, l2 Button, prize Prize) int {
	var xmoves, ymoves int
	divisible := (l2.y*prize.x - l1.y*prize.y) % (l2.y*l1.x - l1.y*l2.x)
	if divisible == 0 {
		xmoves = (l2.y*prize.x - l1.y*prize.y) / (l2.y*l1.x - l1.y*l2.x)
	} else {
		return 0
	}
	ymoves = (prize.x*l2.x - prize.y*l1.x) / (l1.y*l2.x - l2.y*l1.x)
	return xmoves*3 + ymoves
}

func part01(arcadeMachines []ArcadeMachine) {
	totalTokens := 0
	for _, machine := range arcadeMachines {
		l1 := Button{x: machine.A.x, y: machine.B.x}
		l2 := Button{x: machine.A.y, y: machine.B.y}
		t := numTokens(l1, l2, machine.Prize)
		totalTokens += t
	}
	fmt.Println("PART 01:", totalTokens)
}

func part02(arcadeMachines []ArcadeMachine) {
	totalTokens := 0
	for _, machine := range arcadeMachines {
		machine.Prize.x += 10000000000000
		machine.Prize.y += 10000000000000
		l1 := Button{x: machine.A.x, y: machine.B.x}
		l2 := Button{x: machine.A.y, y: machine.B.y}
		t := numTokens(l1, l2, machine.Prize)
		totalTokens += t
	}
	fmt.Println("PART 02:", totalTokens)
}

func Run(dir string) {
	utils.PartPrinter("DAY 13")
	arcadeMachines := processInput(dir + "/day13")
	part01(arcadeMachines)
	part02(arcadeMachines)
}
