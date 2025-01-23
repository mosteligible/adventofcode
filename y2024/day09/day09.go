package day09

import (
	"adventofcode/y2024/utils"
	"fmt"
	"log"
	"strings"
)

const (
	SAMPLE          = "2333133121414131402"
	SPACE_ID        = -1
	SPACE_NOT_FOUND = -100
	FILE_NOT_FOUND  = -2
	MOVED_FILE_ID   = -5
)

func processInput(daydir string) []int {
	content := utils.Read(daydir + "/input.txt")
	// content := SAMPLE
	content = strings.TrimSuffix(content, "\n")
	retval := []int{}

	for _, r := range content {
		n := int(r - '0')
		// fmt.Printf("%c -> %d\n", r, n)
		retval = append(retval, n)
	}
	return retval
}

func interpretInput(input []int) *[]int {
	var currNum int
	retval := []int{}
	for index, num := range input {
		if index%2 == 0 {
			currNum = index / 2
		} else {
			currNum = SPACE_ID
		}
		for range num {
			retval = append(retval, currNum)
		}
	}
	return &retval
}

func checksum(fragmented *[]int) int {
	total := 0
	for idx, val := range *fragmented {
		if val >= 0 {
			total += idx * val
		}
	}
	return total
}

func showDisk(fragmented *[]int) {
	for _, num := range *fragmented {
		if num < -1 {
			// fmt.Printf("%d", num)
			fmt.Printf(".")
		} else if num == -1 {
			fmt.Printf(".")
		} else {
			fmt.Printf("%d", num)
		}
	}
	fmt.Printf("\n")
}

func getRightFileIdIndex(unfragmented *[]int, rightIndex int) int {
	// starting at `rightIndex` find next file id
	rightIndex -= 1
	for rightIndex >= 0 {
		if (*unfragmented)[rightIndex] != -1 {
			return rightIndex
		}
		rightIndex--
	}
	log.Fatalf("file id not found at right!")
	return FILE_NOT_FOUND
}

func part01(data []int) {
	unfragmented := *(interpretInput(data))
	left := 0
	right := getRightFileIdIndex(&unfragmented, len(unfragmented))
	for left <= right {
		if unfragmented[left] == -1 {
			unfragmented[left], unfragmented[right] = unfragmented[right], unfragmented[left]
			right = getRightFileIdIndex(&unfragmented, right)
		}
		left++
	}
	fmt.Printf("PART 01: %d\n", checksum(&unfragmented))
}

func getOccurences(files *[]int, idx int, isReverse bool) int {
	// get number of times value at index `idx` occurs in array at `files`
	// continuously
	currNum := (*files)[idx]
	size := len(*files)
	occurences := 0
	for idx < size && currNum == (*files)[idx] {
		occurences++
		if !isReverse {
			idx++
		} else {
			idx--
		}
	}
	return occurences
}

func firstSpaceIndexAndSize(arr []int) (int, int) {
	// get index of first space in array `arr` and also number of contiguous
	// spaces that follow
	index := utils.GetIndex(&arr, SPACE_ID)
	if index == -1 {
		return SPACE_NOT_FOUND, SPACE_NOT_FOUND
	}
	numSpaces := getOccurences(&arr, index, false)
	return index, numSpaces
}

func getFileIdOccurence(arr []int, rightIndex int) int {
	occurence := 0
	fileId := arr[rightIndex]
	for fileId == arr[rightIndex] {
		occurence++
		rightIndex--

	}
	return occurence
}

func getMatchingFileId(files *[]int, spaceIndex int, numSpaces int) (int, int, int) {
	// from last position, check if any file-set can fill current space size
	for start := len(*files) - 1; start > spaceIndex; start-- {
		fileId := (*files)[start]
		if fileId > 0 {
			occurences := getFileIdOccurence(*files, start)
			if occurences <= numSpaces {
				return fileId, occurences, start - occurences + 1
			}
			start -= occurences
			continue
		}
	}
	return FILE_NOT_FOUND, FILE_NOT_FOUND, FILE_NOT_FOUND
}

func part02(data *[]int) {
	unfragmented := interpretInput(*data)
	// []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	for {
		// find index of space .i.e. index of item that matches SPACE_ID
		spaceIndex, numSpaces := firstSpaceIndexAndSize(*unfragmented)
		if spaceIndex == SPACE_NOT_FOUND {
			break
		}
		// if space is found, check if files exist to fill that contiguous space
		fileId, numFiles, fileidStartIndex := getMatchingFileId(unfragmented, spaceIndex, numSpaces)
		if fileId == FILE_NOT_FOUND {
			utils.Replace(unfragmented, spaceIndex, numSpaces, FILE_NOT_FOUND)
		} else {
			utils.Replace(unfragmented, spaceIndex, numFiles, fileId)
			utils.Replace(
				unfragmented, fileidStartIndex, fileidStartIndex+numFiles-1, MOVED_FILE_ID,
			)
		}
	}
	fmt.Println("PART 02:", checksum(unfragmented), len(*unfragmented))
}

func Run(dir string) {
	utils.PartPrinter("DAY 09")
	data := processInput(dir + "/day09")
	part01(data)
	part02(&data)
}
