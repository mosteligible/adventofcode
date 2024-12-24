package day07

import (
	"aoc2024/utils"
	"fmt"
	"strconv"
	"strings"
)

type Input struct {
	target  int
	numbers []int
}

func (i Input) String() string {
	return fmt.Sprintf(
		"Input(target: %d, numbers: %v)", i.target, i.numbers,
	)
}

var SAMPLE = []string{
	"190: 10 19",
	"3267: 81 40 27",
	"83: 17 5",
	"156: 15 6",
	"7290: 6 8 6 15",
	"161011: 16 10 13",
	"192: 17 8 14",
	"21037: 9 7 18 13",
	"292: 11 6 16 20",
}

func processInput() []Input {
	lines := utils.ReadLines("./day07/input.txt")
	// lines = SAMPLE
	parsed := []Input{}
	for _, line := range lines {
		splitLines := strings.Split(line, ": ")
		intTarget, _ := strconv.Atoi(splitLines[0])
		nums := strings.Split(splitLines[1], " ")
		input := Input{target: intTarget, numbers: []int{}}
		for _, num := range nums {
			intNum, _ := strconv.Atoi(num)
			input.numbers = append(input.numbers, intNum)
		}
		parsed = append(parsed, input)
	}
	return parsed
}

func operate(nums []int, currIndex int, res int, target int) bool {
	if currIndex >= len(nums) && res != target {
		return false
	}
	if currIndex >= len(nums) && res == target {
		return true
	}

	sumRes := nums[currIndex] + res
	sumOk := operate(nums, currIndex+1, sumRes, target)
	if sumOk {
		return true
	}
	prodRes := nums[currIndex] * res
	prodOk := operate(nums, currIndex+1, prodRes, target)
	return prodOk
}

func tenPower(num int) int {
	ten := 10
	for num%ten != num {
		ten *= 10
	}
	return ten
}

func operateExtra(nums []int, currIndex int, res int, target int) bool {
	if currIndex >= len(nums) && res != target {
		return false
	}
	if currIndex >= len(nums) && res == target {
		return true
	}

	sumRes := nums[currIndex] + res
	sumOk := operateExtra(nums, currIndex+1, sumRes, target)
	if sumOk {
		return true
	}
	tenMultiplier := tenPower(nums[currIndex])
	concatRes := res*tenMultiplier + nums[currIndex]
	concatOk := operateExtra(nums, currIndex+1, concatRes, target)
	if concatOk {
		return true
	}
	prodRes := nums[currIndex] * res
	prodOk := operateExtra(nums, currIndex+1, prodRes, target)
	return prodOk
}

func part01(input []Input) {
	total := 0
	for _, d := range input {
		if operate(d.numbers[1:], 0, d.numbers[0], d.target) {
			total += d.target
		}
	}
	fmt.Printf("xx PART 01 total: %d\n", total)
}

func part02(input []Input) {
	total := 0
	for _, d := range input {
		if operateExtra(d.numbers[1:], 0, d.numbers[0], d.target) {
			total += d.target
		}
	}
	fmt.Printf("XX PART 02 total: %d\n", total)
}

func Run() {
	data := processInput()
	part01(data)
	part02(data)
}
