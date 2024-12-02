package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"slices"
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

	slices.Sort(leftSlice)
	slices.Sort(rightSlice)

	sum := 0
	for i := 0; i < len(leftSlice); i++ {
		diff := math.Abs(float64(leftSlice[i]) - float64(rightSlice[i]))
		sum += int(diff)
	}

	fmt.Println(sum)
}
