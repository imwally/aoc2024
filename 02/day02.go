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

	sum := 0
	for i := 0; i < len(leftSlice); i++ {
		timesFound := 0
		for j := 0; j < len(rightSlice); j++ {
			if leftSlice[i] == rightSlice[j] {
				timesFound++
			}
		}
		sum += leftSlice[i] * timesFound
	}

	fmt.Println(sum)
}
