package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const INPUT_FILE = "input.txt"

func in(a int, s []int) bool {
	for _, v := range s {
		if a == v {
			return true
		}
	}

	return false
}

func main() {
	file, err := os.Open(INPUT_FILE)
	if err != nil {
		log.Println(err)
	}
	defer file.Close()

	updateLines := [][]int{}
	pageMap := make(map[int][]int)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		split := strings.Split(line, "|")
		if len(split) > 1 {
			before, _ := strconv.Atoi(split[0])
			after, _ := strconv.Atoi(split[1])
			pageMap[before] = append(pageMap[before], after)

		} else {
			nums := []int{}
			split := strings.Split(line, ",")
			for _, v := range split {
				num, _ := strconv.Atoi(v)
				nums = append(nums, num)
			}
			updateLines = append(updateLines, nums)
		}
	}

	ignoreLines := make(map[int]bool)
	for x, line := range updateLines {
		for i := len(line) - 1; i > 0; i-- {
			page := line[i]
			pagesBefore := line[:i]
			for _, pageBefore := range pagesBefore {
				if in(pageBefore, pageMap[page]) {
					ignoreLines[x] = true
				}
			}
		}
	}

	swapped := false
	for k := 0; k < len(updateLines); k++ {
		if _, ok := ignoreLines[k]; !ok {
			continue
		}
		updateLine := updateLines[k]
		for i := len(updateLine) - 1; i >= 0; i-- {
			page := updateLine[i]
			pagesThatShouldComeAfter := pageMap[page]
			for x := i; x >= 0; x-- {
				xPage := updateLine[x]
				if in(xPage, pagesThatShouldComeAfter) {
					updateLine[i], updateLine[x] = updateLine[x], updateLine[i]
					swapped = true
				}
			}
		}
		if swapped {
			k = k - 1
		}
		swapped = false
	}

	sum := 0
	for k := range ignoreLines {
		sum += updateLines[k][len(updateLines[k])/2]
	}
	fmt.Println(sum)
}
