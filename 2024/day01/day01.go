package day01

import (
	"aoc2024/utils"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func processInput() ([]int, []int) {
	leftArr := []int{}
	rightArr := []int{}
	lines := utils.ReadLines("./day01/input.txt")
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
	utils.PartPrinter("PART 01")
	distance := 0
	for index, l := range left {
		distance += utils.IntAbs(l - right[index])
	}
	fmt.Printf("total: %d\n", distance)
}

func part02(left []int, right []int) {
	utils.PartPrinter("PART 02")
	similarity := 0
	counts := utils.Counter(right)
	for _, l := range left {
		score, ok := counts[l]
		if ok {
			similarity += l * score
		}
	}
	fmt.Printf("similarity: %d\n", similarity)
}

func Run() {
	fmt.Println("XXXXXXXXXXXXXXX DAY 02 XXXXXXXXXXXXXXX")
	left, right := processInput()
	part01(left, right)
	part02(left, right)
}
