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

func ReverseString(s string) string {
	strRune := []rune(s)
	for i, j := 0, len(strRune)-1; i < j; i, j = i+1, j-1 {
		intermediate := strRune[i]
		strRune[i] = strRune[j]
		strRune[j] = intermediate
	}
	return string(strRune)
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

func PartPrinter(partstr string) {
	fmt.Printf("XXXXXXXXXXXXXXXX %s XXXXXXXXXXXXXXXX\n", partstr)
}
