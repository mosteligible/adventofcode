package utils

import (
	"slices"
)

type Number struct {
	Num  string
	Next map[string]Number
}

func isPointedTo(num string, order *map[string]Number, allNums *[]string) bool {
	// from a list of numbers (allNums), check if a number (num) has path to it
	// based on the edge rules from (order)
	if len(*allNums) == 1 {
		return false
	}
	for _, aNum := range *allNums {
		if aNum == num {
			continue
		}
		nextNum := (*order)[aNum]
		if _, ok := nextNum.Next[num]; ok {
			return true
		}
	}
	return false
}

func TopologicalSort(items []string, order *map[string]Number) []string {
	sortedEntries := []string{}
	var numToAdd string
	// from the list of strings, node that is not pointed to will be added first
	for len(items) > 0 {
		// from the list of items, select first number
		idxToDelete := -1
		for idx, num := range items {
			// check if `num` has no nodes pointing to it in graph
			if slices.Contains(sortedEntries, num) {
				continue
			}
			if !isPointedTo(num, order, &items) {
				numToAdd = num
				idxToDelete = idx
				break
			}
		}
		if idxToDelete != -1 {
			sortedEntries = append(sortedEntries, numToAdd)
			items = append(items[:idxToDelete], items[idxToDelete+1:]...)
		}
	}
	return sortedEntries
}
