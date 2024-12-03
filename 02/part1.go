package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

const (
	ASC = iota
	DESC

	INPUT_FILE = "input.txt"
)

func safe(line []byte) bool {
	direction := ASC
	nums := strings.Split(string(line), " ")

	for i := 0; i < len(nums)-1; i++ {
		a, _ := strconv.Atoi(nums[i])
		b, _ := strconv.Atoi(nums[i+1])

		if i == 0 {
			if a > b {
				direction = DESC
			}
		}

		if i > 0 {
			if direction == ASC && b < a {
				return false
			}
			if direction == DESC && b > a {
				return false
			}
		}

		diff := math.Abs(float64(a) - float64(b))
		if diff > 3 || diff == 0 {
			return false
		}
	}

	return true
}

func main() {
	file, err := os.Open(INPUT_FILE)
	if err != nil {
		log.Println(err)
	}
	defer file.Close()

	sum := 0
	reader := bufio.NewReader(file)
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		safe := safe(line)
		if safe {
			sum++
		}
	}

	fmt.Println(sum)
}
