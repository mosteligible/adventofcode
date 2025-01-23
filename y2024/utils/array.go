package utils

import (
	"log"
)

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

func ArrayEqual[T comparable](left, right []T) bool {
	if len(left) != len(right) {
		return false
	}

	for index, item := range left {
		if item != right[index] {
			return false
		}
	}
	return true
}

func Replace[T any](arr *[]T, start int, numItems int, val T) {
	end := start + numItems
	if end >= len(*arr) {
		end = len(*arr)
	}
	for ; start < end; start++ {
		(*arr)[start] = val
	}
}

func ReverseSlice[T any](arr []T) []T {
	retval := make([]T, len(arr))
	left := 0
	for right := len(arr) - 1; right >= 0; right-- {
		retval[left] = arr[right]
		left++
	}
	return retval
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

func MakeMap[T comparable](arr []T) map[T]bool {
	retval := map[T]bool{}
	for _, val := range arr {
		retval[val] = true
	}
	return retval
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
