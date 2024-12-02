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
const DELIM = "   "

func main() {
	file, err := os.Open(INPUT_FILE)
	if err != nil {
		log.Println(err)
	}
	defer file.Close()

	leftSlice := []int{}
	rightSlice := []int{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		splitString := strings.Split(scanner.Text(), DELIM)
		leftNum, _ := strconv.Atoi(splitString[0])
		rightNum, _ := strconv.Atoi(splitString[1])

		leftSlice = append(leftSlice, leftNum)
		rightSlice = append(rightSlice, rightNum)
	}

	similarityMap := make(map[int]int, len(leftSlice))

	for i := 0; i < len(leftSlice); i++ {
		similarityMap[leftSlice[i]] = 0
	}

	for i := 0; i < len(rightSlice); i++ {
		current := rightSlice[i]
		if _, ok := similarityMap[current]; ok {
			similarityMap[current]++
		}
	}

	sum := 0
	for k, v := range similarityMap {
		sum += k * v
	}

	fmt.Println(sum)
}
