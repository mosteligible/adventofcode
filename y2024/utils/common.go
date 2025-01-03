package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
)

func ReadLines(filepath string) []string {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatalf("file not found: %s", filepath)
	}
	fileScanner := bufio.NewScanner(file)
	retval := []string{}

	for fileScanner.Scan() {
		retval = append(retval, fileScanner.Text())
	}

	return retval
}

func Read(filepath string) string {
	file, err := os.ReadFile(filepath)
	if err != nil {
		log.Fatalf("file not found: %s", filepath)
	}
	s := string(file)
	return s
}

func IntAbs(num int) int {
	if num < 0 {
		return num * -1
	}
	return num
}

func Counter[T comparable](list []T) map[T]int {
	counts := map[T]int{}
	for _, item := range list {
		_, ok := counts[item]
		if ok {
			counts[item] += 1
		} else {
			counts[item] = 1
		}
	}
	return counts
}

func PartPrinter(partstr string) {
	fmt.Printf("XXXXXXXXXXXXXXXX %s XXXXXXXXXXXXXXXX\n", partstr)
}

func PopIndex[T any](arr []T, idx int) (T, []T) {
	if idx >= len(arr) {
		log.Fatalf("index out of bound!")
	}
	retT := arr[idx]
	if idx == 0 {
		return retT, arr[idx+1:]
	}
	newArr := append(arr[:idx], arr[idx+1:]...)
	return retT, newArr
}

func ReverseString(s string) string {
	strRune := []rune(s)
	for i, j := 0, len(strRune)-1; i < j; i, j = i+1, j-1 {
		intermediate := strRune[i]
		strRune[i] = strRune[j]
		strRune[j] = intermediate
	}
	return string(strRune)
}

func GetIndex[T comparable](arr *[]T, element T) int {
	for idx, val := range *arr {
		if val == element {
			return idx
		}
	}
	return -1
}

func Count[T comparable](arr []T, element T) int {
	occurence := 0

	for _, item := range arr {
		if item == element {
			occurence++
		}
	}

	return occurence
}

func Replace[T any](arr *[]T, start int, end int, val T) {
	for ; start <= end; start++ {
		(*arr)[start] = val
	}
}

func ClearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func GetZero[T any]() T {
	var retval T
	return retval
}

func Sum(arr []int) int {
	total := 0
	for _, val := range arr {
		total += val
	}
	return total
}

func MakeMap[T comparable](arr []T) map[T]bool {
	retval := map[T]bool{}
	for _, val := range arr {
		retval[val] = true
	}
	return retval
}
