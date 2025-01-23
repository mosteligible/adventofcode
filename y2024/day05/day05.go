package day05

import (
	"adventofcode/y2024/utils"
	"fmt"
	"strconv"
	"strings"
)

func processInput(dayDir string) (map[string]utils.Number, [][]string) {
	lines := utils.ReadLines(dayDir + "/input_rules.txt")
	process := map[string]utils.Number{}
	for _, line := range lines {
		nums := strings.Split(line, "|")
		previous, prevok := process[nums[0]]
		next, nextok := process[nums[1]]
		if !nextok {
			next = utils.Number{Num: nums[1], Next: map[string]utils.Number{}}
			process[nums[1]] = next
		}
		if !prevok {
			previous = utils.Number{Num: nums[0], Next: map[string]utils.Number{nums[1]: next}}
		}
		if _, ok := previous.Next[next.Num]; !ok {
			previous.Next[next.Num] = next
		}
	}
	ordering := utils.ReadLines(dayDir + "/input_ordering.txt")
	orders := [][]string{}
	for _, order := range ordering {
		nums := strings.Split(order, ",")
		orders = append(orders, nums)
	}
	return process, orders
}

func isPrevPageSeen(seenPages []string, afterPages map[string]utils.Number) bool {
	for _, seenPageNo := range seenPages {
		if _, ok := afterPages[seenPageNo]; ok {
			return true
		}
	}
	return false
}

func part01(rules map[string]utils.Number, ordering [][]string) ([][]string, int) {
	validOrders := [][]string{}
	invalidOrders := [][]string{}

	for _, order := range ordering {
		seenPages := []string{}
		invalidOrder := false
		for _, pg_no := range order {
			afterPgNo := rules[pg_no]
			if isPrevPageSeen(seenPages, afterPgNo.Next) {
				invalidOrder = true
				invalidOrders = append(invalidOrders, order)
				break
			}
			seenPages = append(seenPages, pg_no)
		}
		if !invalidOrder {
			validOrders = append(validOrders, order)
		}
	}

	midSum := 0
	for _, order := range validOrders {
		mid := order[len(order)/2]
		midInt, _ := strconv.Atoi(mid)
		midSum += midInt
	}
	fmt.Println("Part 01:", midSum)

	return invalidOrders, midSum
}

func part02(rules map[string]utils.Number, ordering [][]string) {
	invalidOrders, _ := part01(rules, ordering)
	invalidSum := 0
	for _, order := range invalidOrders {
		sortedOrder := utils.TopologicalSort(order, &rules)
		midInt, _ := strconv.Atoi(sortedOrder[len(sortedOrder)/2])
		invalidSum += midInt
	}
	fmt.Println("Part 02:", invalidSum)
}

func Run(dir string) {
	utils.PartPrinter("DAY 05")
	rules, ordering := processInput(dir + "/day05")
	part02(rules, ordering)
}
