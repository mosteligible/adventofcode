package day02

import (
	"adventofcode/y2024/utils"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func processInput(dayDir string) [][]int {
	lines := utils.ReadLines(dayDir + "/input.txt")
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

func isLevelSafe(level []int) bool {
	decreasing := false
	if level[0] > level[1] {
		decreasing = true
	} else {
		decreasing = false
	}
	index := 0
	for _, l := range level[1:] {
		if level[index] == l {
			return false
		}
		if (decreasing && level[index] < l) || (!decreasing && level[index] > l) {
			return false
		}
		if (decreasing && level[index]-l > 3) || (!decreasing && l-level[index] > 3) {
			return false
		}
		index += 1
	}
	return true
}

func part01(data [][]int) {
	safe := 0
	for _, levels := range data {
		if isLevelSafe(levels) {
			safe++
		}
	}
	fmt.Println("Part 01:", safe)
}

func part02(data [][]int) {
	safe := 0
	var l []int
	for _, level := range data {
		for idx := range len(level) {
			l = append(level[:idx], level[idx+1:]...)
			if isLevelSafe(l) {
				safe += 1
				break
			}
		}
	}
	fmt.Println("Part 02:", safe)
}

func Run(dir string) {
	utils.PartPrinter("DAY 02")
	data := processInput(dir + "/day02")
	part01(data)
	part02(data)
}
