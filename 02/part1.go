package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strconv"
)

const INPUT_FILE = "input.txt"

func safe(line []byte) bool {
	last, _ := strconv.Atoi(string(line[:2]))
	for i := 0; i < len(line); i++ {
		if i > 2 && line[i] != byte(' ') {
			current := line[i : i+2]
			currentInt, _ := strconv.Atoi(string(current))
			diff := math.Abs(float64(currentInt) - float64(last))
			fmt.Println(diff)
			if diff > 3 || diff == 0 {
				return false
			}

			i = i + 2
			last = currentInt
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
		if safe(line) {
			sum++
		}
	}

	fmt.Println(sum)
}
