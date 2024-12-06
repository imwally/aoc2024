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

	updateLines := []string{}
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
			updateLines = append(updateLines, line)
		}
	}

	ignoreLines := make(map[int]bool)
	for x, line := range updateLines {
		pages := strings.Split(line, ",")
		for i := len(pages) - 1; i > 0; i-- {
			page, _ := strconv.Atoi(pages[i])
			pagesBefore := pages[:i]

			for _, pageBefore := range pagesBefore {
				pbInt, _ := strconv.Atoi(pageBefore)
				if in(pbInt, pageMap[page]) {
					ignoreLines[x] = true
				}
			}
		}
	}

	sum := 0
	for y, updateLine := range updateLines {
		if y < 1 {
			continue
		}
		if _, ok := ignoreLines[y]; !ok {
			split := strings.Split(updateLine, ",")
			v, _ := strconv.Atoi(split[len(split)/2])
			sum += v
		}
	}
	fmt.Println(sum)
}
