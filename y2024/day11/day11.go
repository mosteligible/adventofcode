package day11

import (
	"adventofcode/y2024/utils"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func processInput(daydir string) map[int]int {
	content := utils.Read(daydir + "/input.txt")
	content = strings.TrimPrefix(content, "\n")
	splitContent := strings.Split(content, " ")
	retval := map[int]int{}
	for _, numbah := range splitContent {
		n, _ := strconv.Atoi(numbah)
		if _, ok := retval[n]; ok {
			retval[n] += 1
		} else {
			retval[n] = 1
		}
	}
	if _, ok := retval[0]; !ok {
		retval[0] = 0
	}
	if _, ok := retval[1]; !ok {
		retval[1] = 0
	}
	return retval
}

func getNumDigits(num int) int {
	return int(math.Floor(math.Log10(float64(num)) + 1))
}

func CountRocks(rocks *map[int]int) int {
	total := 0

	for _, v := range *rocks {
		total += v
	}

	return total
}

func part01(rocks map[int]int, numBlinks int, part string) {
	for range numBlinks {
		newRocks := map[int]int{}
		// fmt.Printf(" len(newRocks): %d\n", len(newRocks))
		// fmt.Printf("%v, ", rocks)
		for r, count := range rocks {
			numDigits := getNumDigits(r)
			if r == 0 {
				if _, ok := newRocks[1]; !ok {
					newRocks[1] = 0
				}
				newRocks[1] += count
			} else if numDigits%2 == 0 {
				tenPow := numDigits / 2
				left := r / int(math.Pow10(tenPow))
				right := r % int(math.Pow10(tenPow))
				if _, ok := newRocks[left]; !ok {
					newRocks[left] = 0
				}
				if _, ok := newRocks[right]; !ok {
					newRocks[right] = 0
				}
				newRocks[left] += count
				newRocks[right] += count
			} else {
				if _, ok := newRocks[r]; !ok {
					newRocks[r] = 0
				}
				newRocks[r*2024] += count
			}
		}
		rocks = newRocks
	}
	fmt.Printf("%s: %d\n", part, CountRocks(&rocks))
}

func Run(dir string) {
	utils.PartPrinter("DAY 11")
	data := processInput(dir + "/day11")
	part01(data, 25, "PART 01")
	part01(data, 75, "PART 02")
}
