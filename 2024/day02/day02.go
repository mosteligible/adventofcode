package day02

import (
	"aoc2024/utils"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func processInput() [][]int {
	lines := utils.ReadLines("./day02/input.txt")
	parsed := [][]int{}
	for _, line := range lines {
		strnums := strings.Split(line, " ")
		nums := []int{}
		for _, n := range strnums {
			nInt, err := strconv.Atoi(n)
			if err != nil {
				log.Fatalf(
					"could not convert to integer. line: <%s> num: <%s>\n", line, n,
				)
			}
			nums = append(nums, nInt)
		}
		parsed = append(parsed, nums)
	}
	return parsed
}

func part01(data [][]int) {
	unsafe := 0
	safe := 0
	decreasing := false
	for _, levels := range data {
		if levels[0] > levels[1] {
			decreasing = true
		} else {
			decreasing = false
		}
		index := 0
		isSafe := true
		for _, l := range levels[1:] {
			if levels[index] == l {
				unsafe += 1
				isSafe = false
				break
			}
			if (decreasing && levels[index] < l) || (!decreasing && levels[index] > l) {
				unsafe += 1
				isSafe = false
				break
			}
			if (decreasing && levels[index]-l > 3) || (!decreasing && l-levels[index] > 3) {
				unsafe += 1
				isSafe = false
				break
			}
			index += 1
		}
		if isSafe {
			safe += 1
		}
	}
	fmt.Println("safe:", safe)
	fmt.Println("unsafe:", unsafe)
}

func Run() {
	utils.PartPrinter("DAY 02")
	data := processInput()
	part01(data)
}
