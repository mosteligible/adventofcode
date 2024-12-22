package day19

import (
	"aoc2024/utils"
	"fmt"
	"strings"
)

const SAMPLE = `r, wr, b, g, bwu, rb, gb, br

brwrr
bggr
gbbr
rrbgbr
ubwu
bwurrg
brgr
bbrgwb`

var MEMO = map[string]bool{}
var MEMO_WAYS = map[string]int{}

type Input struct {
	Puzzles []string
	Stripes []string
}

func (i *Input) ProcessInput() {
	content := utils.Read("./day19/input.txt")
	// content = SAMPLE
	input := strings.Split(content, "\n\n")
	i.Stripes = strings.Split(input[0], ", ")
	i.Puzzles = strings.Split(input[1], "\n")
}

func NewInput() Input {
	i := Input{}
	i.ProcessInput()
	return i
}

func matchStripes(stripes []string, puzzleStr string) bool {
	if _, ok := MEMO[puzzleStr]; ok {
		return MEMO[puzzleStr]
	}

	if puzzleStr == "" {
		return true
	}

	for _, word := range stripes {
		if strings.HasPrefix(puzzleStr, word) {
			matched := matchStripes(stripes, puzzleStr[len(word):])
			if matched {
				MEMO[puzzleStr] = true
				return true
			}
		}
	}
	MEMO[puzzleStr] = false
	return false
}

func countWays(stripes []string, puzzleStr string) int {
	if _, ok := MEMO_WAYS[puzzleStr]; ok {
		return MEMO_WAYS[puzzleStr]
	}
	if puzzleStr == "" {
		return 1
	}

	ways := 0
	for _, word := range stripes {
		if strings.HasPrefix(puzzleStr, word) {
			ways += countWays(stripes, puzzleStr[len(word):])
		}
	}
	MEMO_WAYS[puzzleStr] = ways
	return ways
}

func Part01(input Input) {
	numSolvables := 0
	for _, puzzle := range input.Puzzles {
		isSolvable := matchStripes(input.Stripes, puzzle)
		fmt.Printf("pzl: %s, ", puzzle)
		if isSolvable {
			numSolvables++
		}
		fmt.Printf("solvable: %t\n", isSolvable)
	}
	fmt.Println(" [o] Solvables:", numSolvables)
}

func Part02(input Input) {
	utils.PartPrinter("Part 02")
	numWays := 0
	for _, puzzle := range input.Puzzles {
		ways := countWays(input.Stripes, puzzle)
		numWays += ways
	}
	fmt.Println(" [o] Ways:", numWays)
}

func Run() {
	input := NewInput()
	Part01(input)
	Part02(input)
}
