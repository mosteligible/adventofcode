package day03

import (
	"adventofcode/y2024/utils"
	"fmt"
	"log"
	"regexp"
	"strconv"
)

const (
	MUL_DIGIT_PATTERN = `mul\((\d{1,3},\d{1,3})\)`
	COND_PATTERN      = `(do\(\))|(don't\(\))|mul\((\d{1,3},\d{1,3})\)`

	DO   = -1
	DONT = -2
)

var (
	MUL_DIGIT_REGEX = regexp.MustCompile(`\d+`)
	COND_REGEX      = regexp.MustCompile(COND_PATTERN)
)

func processInput(dayDir string) [][2]int {
	content := utils.ReadLines(dayDir + "/input.txt")
	lineNumsProcessed := [][2]int{}
	for _, line := range content {
		lineNums := COND_REGEX.FindAllString(line, -1)
		for _, nums := range lineNums {
			if nums == "do()" {
				lineNumsProcessed = append(lineNumsProcessed, [2]int{-1, -1})
			} else if nums == "don't()" {
				lineNumsProcessed = append(lineNumsProcessed, [2]int{-2, -2})
			} else {
				integers := MUL_DIGIT_REGEX.FindAllString(nums, -1)
				d1, err := strconv.Atoi(integers[0])
				if err != nil {
					log.Fatalf("Could not convert <%s> to integer from <%s>", integers[0], nums)
				}
				d2, err := strconv.Atoi(integers[1])
				if err != nil {
					log.Fatalf("Could not convert <%s> to integer from <%s>", integers[1], nums)
				}
				lineNumsProcessed = append(lineNumsProcessed, [2]int{d1, d2})
			}
		}
	}
	return lineNumsProcessed
}

func part01(data *[][2]int) {
	total := 0
	for _, nums := range *data {
		if nums[0] != -1 && nums[0] != -2 {
			total += nums[0] * nums[1]
		}
	}
	fmt.Printf("Part 01: %d\n", total)
}

func part02(data *[][2]int) {
	total := 0
	do := true
	for _, nums := range *data {
		if nums[1] == DO {
			do = true
		} else if nums[1] == DONT {
			do = false
		} else if do {
			total += nums[0] * nums[1]
		}
	}
	fmt.Printf("Part 02: %d\n", total)
}

func Run(dir string) {
	utils.PartPrinter("DAY 03")
	data := processInput(dir + "/day03")
	part01(&data)
	part02(&data)
}
