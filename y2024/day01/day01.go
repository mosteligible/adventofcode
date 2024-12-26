package day01

import (
	"adventofcode/y2024/utils"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func processInput(dayDir string) ([]int, []int) {
	leftArr := []int{}
	rightArr := []int{}
	lines := utils.ReadLines(dayDir + "/input.txt")
	for _, line := range lines {
		splitRes := strings.Split(line, "   ")
		left, _ := strconv.Atoi(splitRes[0])
		right, _ := strconv.Atoi(splitRes[1])
		leftArr = append(leftArr, left)
		rightArr = append(rightArr, right)
	}
	slices.Sort(leftArr)
	slices.Sort(rightArr)
	return leftArr, rightArr
}

func part01(left []int, right []int) {
	distance := 0
	for index, l := range left {
		distance += utils.IntAbs(l - right[index])
	}
	fmt.Printf("Part 01: %d\n", distance)
}

func part02(left []int, right []int) {
	similarity := 0
	counts := utils.Counter(right)
	for _, l := range left {
		score, ok := counts[l]
		if ok {
			similarity += l * score
		}
	}
	fmt.Printf("Part 02: %d\n", similarity)
}

func Run(dirPath string) {
	fmt.Println("XXXXXXXXXXXXXXX DAY 01 XXXXXXXXXXXXXXX")
	left, right := processInput(dirPath + "/day01")
	part01(left, right)
	part02(left, right)
}
