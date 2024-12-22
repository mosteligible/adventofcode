package day11

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

var input = "17639 47 3858 0 470624 9467423 5 188"

func processInput() map[int]int {
	content := input
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

func part01(rocks map[int]int, numBlinks int) {
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
	fmt.Printf("numrocks: %d\n", CountRocks(&rocks))
}

func Run() {
	data := processInput()
	fmt.Println(data)
	part01(data, 25)
	part01(data, 75)
}
